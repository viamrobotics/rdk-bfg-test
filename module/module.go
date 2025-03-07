package module

import (
	"context"
	"fmt"
	"net"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/fullstorydev/grpcurl"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/grpcreflect"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	pb "go.viam.com/api/module/v1"
	robotpb "go.viam.com/api/robot/v1"
	"go.viam.com/utils"
	"go.viam.com/utils/rpc"
	"google.golang.org/grpc"
	reflectpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"

	"go.viam.com/rdk/config"
	"go.viam.com/rdk/operation"
	"go.viam.com/rdk/protoutils"
	"go.viam.com/rdk/registry"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/robot/client"
	"go.viam.com/rdk/subtype"
)

// CheckSocketAddressLength returns an error if the socket path is too long for the OS.
func CheckSocketAddressLength(addr string) error {
	// maxSocketAddressLength is the length (-1 for null terminator) of the .sun_path field as used in kernel bind()/connect() syscalls.
	maxSocketAddressLength := 103
	if runtime.GOOS == "linux" {
		maxSocketAddressLength = 107
	}
	if len(addr) > maxSocketAddressLength {
		return errors.Errorf("module socket path exceeds OS limit of %d characters: %s", maxSocketAddressLength, addr)
	}
	return nil
}

// HandlerMap is the format for api->model pairs that the module will service.
// Ex: mymap["rdk:component:motor"] = ["acme:marine:thruster", "acme:marine:outboard"].
type HandlerMap map[resource.RPCSubtype][]resource.Model

// ToProto converts the HandlerMap to a protobuf representation.
func (h HandlerMap) ToProto() *pb.HandlerMap {
	pMap := &pb.HandlerMap{}
	for s, models := range h {
		subtype := &robotpb.ResourceRPCSubtype{
			Subtype: protoutils.ResourceNameToProto(resource.Name{
				Subtype: s.Subtype,
				Name:    "",
			}),
			ProtoService: s.ProtoSvcName,
		}

		handler := &pb.HandlerDefinition{Subtype: subtype}
		for _, m := range models {
			handler.Models = append(handler.Models, m.String())
		}
		pMap.Handlers = append(pMap.Handlers, handler)
	}
	return pMap
}

// NewHandlerMapFromProto converts protobuf to HandlerMap.
func NewHandlerMapFromProto(ctx context.Context, pMap *pb.HandlerMap, conn *grpc.ClientConn) (HandlerMap, error) {
	hMap := make(HandlerMap)
	refClient := grpcreflect.NewClientV1Alpha(ctx, reflectpb.NewServerReflectionClient(conn))
	defer refClient.Reset()
	reflSource := grpcurl.DescriptorSourceFromServer(ctx, refClient)

	var errs error
	for _, h := range pMap.GetHandlers() {
		api := protoutils.ResourceNameFromProto(h.Subtype.Subtype).Subtype

		symDesc, err := reflSource.FindSymbol(h.Subtype.ProtoService)
		if err != nil {
			errs = multierr.Combine(errs, err)
			if errors.Is(err, grpcurl.ErrReflectionNotSupported) {
				return nil, errs
			}
			continue
		}
		svcDesc, ok := symDesc.(*desc.ServiceDescriptor)
		if !ok {
			return nil, errors.Errorf("expected descriptor to be service descriptor but got %T", symDesc)
		}
		subtype := &resource.RPCSubtype{
			Subtype: api,
			Desc:    svcDesc,
		}
		for _, m := range h.Models {
			model, err := resource.NewModelFromString(m)
			if err != nil {
				return nil, err
			}
			hMap[*subtype] = append(hMap[*subtype], model)
		}
	}
	return hMap, errs
}

// Module represents an external resource module that services components/services.
type Module struct {
	parent                  *client.RobotClient
	server                  rpc.Server
	logger                  *zap.SugaredLogger
	mu                      sync.Mutex
	operations              *operation.Manager
	ready                   bool
	addr                    string
	parentAddr              string
	activeBackgroundWorkers sync.WaitGroup
	handlers                HandlerMap
	services                map[resource.Subtype]subtype.Service
	closeOnce               sync.Once
	pb.UnimplementedModuleServiceServer
}

// NewModule returns the basic module framework/structure.
func NewModule(ctx context.Context, address string, logger *zap.SugaredLogger) (*Module, error) {
	// TODO(PRODUCT-343): session support likely means interceptors here
	opMgr := operation.NewManager(logger)
	unaries := []grpc.UnaryServerInterceptor{
		opMgr.UnaryServerInterceptor,
	}
	streams := []grpc.StreamServerInterceptor{
		opMgr.StreamServerInterceptor,
	}
	m := &Module{
		logger:     logger,
		addr:       address,
		operations: opMgr,
		server:     NewServer(unaries, streams),
		ready:      true,
		handlers:   HandlerMap{},
		services:   map[resource.Subtype]subtype.Service{},
	}
	if err := m.server.RegisterServiceServer(ctx, &pb.ModuleService_ServiceDesc, m); err != nil {
		return nil, err
	}
	return m, nil
}

// NewModuleFromArgs directly parses the command line argument to get its address.
func NewModuleFromArgs(ctx context.Context, logger *zap.SugaredLogger) (*Module, error) {
	if len(os.Args) != 2 {
		return nil, errors.New("need socket path as command line argument")
	}
	return NewModule(ctx, os.Args[1], logger)
}

// Start starts the module service and grpc server.
func (m *Module) Start(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	var lis net.Listener
	if err := MakeSelfOwnedFilesFunc(func() error {
		var err error
		lis, err = net.Listen("unix", m.addr)
		if err != nil {
			return errors.WithMessage(err, "failed to listen")
		}
		return nil
	}); err != nil {
		return err
	}

	m.activeBackgroundWorkers.Add(1)
	utils.PanicCapturingGo(func() {
		defer m.activeBackgroundWorkers.Done()
		defer utils.UncheckedErrorFunc(func() error {
			// Attempt to remove module's .sock file.
			if _, err := os.Stat(m.addr); err == nil {
				return os.Remove(m.addr)
			}
			return nil
		})
		m.logger.Infof("server listening at %v", lis.Addr())
		if err := m.server.Serve(lis); err != nil {
			m.logger.Errorf("failed to serve: %v", err)
		}
	})
	return nil
}

// Close shuts down the module and grpc server.
func (m *Module) Close(ctx context.Context) {
	m.closeOnce.Do(func() {
		m.mu.Lock()
		parent := m.parent
		m.mu.Unlock()
		m.logger.Info("Shutting down gracefully.")
		if parent != nil {
			if err := parent.Close(ctx); err != nil {
				m.logger.Error(err)
			}
		}
		if err := m.server.Stop(); err != nil {
			m.logger.Error(err)
		}
		m.activeBackgroundWorkers.Wait()
	})
}

// GetParentResource returns a resource from the parent robot by name.
func (m *Module) GetParentResource(ctx context.Context, name resource.Name) (interface{}, error) {
	if err := m.connectParent(ctx); err != nil {
		return nil, err
	}

	// Refresh parent to ensure it has the most up-to-date resources before calling
	// ResourceByName.
	if err := m.parent.Refresh(ctx); err != nil {
		return nil, err
	}
	return m.parent.ResourceByName(name)
}

func (m *Module) connectParent(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.parent == nil {
		if err := CheckSocketOwner(m.parentAddr); err != nil {
			return err
		}
		// TODO(PRODUCT-343): add session support to modules
		rc, err := client.New(ctx, "unix://"+m.parentAddr, m.logger, client.WithDisableSessions())
		if err != nil {
			return err
		}
		m.parent = rc
	}
	return nil
}

// SetReady can be set to false if the module is not ready (ex. waiting on hardware).
func (m *Module) SetReady(ready bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.ready = ready
}

// Ready receives the parent address and reports api/model combos the module is ready to service.
func (m *Module) Ready(ctx context.Context, req *pb.ReadyRequest) (*pb.ReadyResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.parentAddr = req.GetParentAddress()

	return &pb.ReadyResponse{
		Ready:      m.ready,
		Handlermap: m.handlers.ToProto(),
	}, nil
}

// AddResource receives the component/service configuration from the parent.
func (m *Module) AddResource(ctx context.Context, req *pb.AddResourceRequest) (*pb.AddResourceResponse, error) {
	deps := make(registry.Dependencies)
	for _, c := range req.Dependencies {
		name, err := resource.NewFromString(c)
		if err != nil {
			return nil, err
		}
		c, err := m.GetParentResource(ctx, name)
		if err != nil {
			return nil, err
		}
		deps[name] = c
	}

	cfg, err := config.ComponentConfigFromProto(req.Config)
	if err != nil {
		return nil, err
	}

	if err := addConvertedAttributes(cfg); err != nil {
		return nil, errors.Wrapf(err, "unable to convert attributes when adding resource")
	}

	var res interface{}
	switch cfg.API.ResourceType {
	case resource.ResourceTypeComponent:
		creator := registry.ComponentLookup(cfg.API, cfg.Model)
		if creator != nil && creator.Constructor != nil {
			res, err = creator.Constructor(ctx, deps, *cfg, m.logger)
		}

	case resource.ResourceTypeService:
		creator := registry.ServiceLookup(cfg.API, cfg.Model)
		if creator != nil && creator.Constructor != nil {
			res, err = creator.Constructor(ctx, deps, config.ServiceConfigFromShared(*cfg), m.logger)
		}
	default:
		return nil, errors.Errorf("unknown resource type %s", cfg.API.ResourceType)
	}

	if err != nil {
		return nil, err
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	subSvc, ok := m.services[cfg.API]
	if !ok {
		return nil, errors.Errorf("module cannot service api: %s", cfg.API)
	}

	return &pb.AddResourceResponse{}, subSvc.Add(cfg.ResourceName(), res)
}

// ReconfigureResource receives the component/service configuration from the parent.
func (m *Module) ReconfigureResource(ctx context.Context, req *pb.ReconfigureResourceRequest) (*pb.ReconfigureResourceResponse, error) {
	var res interface{}
	deps := make(registry.Dependencies)
	for _, c := range req.Dependencies {
		name, err := resource.NewFromString(c)
		if err != nil {
			return nil, err
		}
		c, err := m.GetParentResource(ctx, name)
		if err != nil {
			return nil, err
		}
		deps[name] = c
	}

	cfg, err := config.ComponentConfigFromProto(req.Config)
	if err != nil {
		return nil, err
	}

	if err := addConvertedAttributes(cfg); err != nil {
		return nil, errors.Wrapf(err, "unable to convert attributes when reconfiguring resource")
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	svc, ok := m.services[cfg.API]
	if !ok {
		return nil, errors.Errorf("no rpc service for %+v", cfg)
	}
	res = svc.Resource(cfg.ResourceName().Name)

	// Check if component directly supports reconfiguration.
	rc, ok := res.(registry.ReconfigurableComponent)
	if ok {
		return &pb.ReconfigureResourceResponse{}, rc.Reconfigure(ctx, *cfg, deps)
	}

	// Check if service directly supports reconfiguration.
	rs, ok := res.(registry.ReconfigurableService)
	if ok {
		return &pb.ReconfigureResourceResponse{}, rs.Reconfigure(ctx, config.ServiceConfigFromShared(*cfg), deps)
	}

	// If it can't reconfigure, replace it.
	if err := utils.TryClose(ctx, res); err != nil {
		m.logger.Error(err)
	}

	switch cfg.API.ResourceType {
	case resource.ResourceTypeComponent:
		creator := registry.ComponentLookup(cfg.API, cfg.Model)
		if creator != nil && creator.Constructor != nil {
			comp, err := creator.Constructor(ctx, deps, *cfg, m.logger)
			if err != nil {
				return nil, err
			}

			return &pb.ReconfigureResourceResponse{}, svc.ReplaceOne(cfg.ResourceName(), comp)
		}

	case resource.ResourceTypeService:
		creator := registry.ServiceLookup(cfg.API, cfg.Model)
		if creator != nil && creator.Constructor != nil {
			s, err := creator.Constructor(ctx, deps, config.ServiceConfigFromShared(*cfg), m.logger)
			if err != nil {
				return nil, err
			}
			return &pb.ReconfigureResourceResponse{}, svc.ReplaceOne(cfg.ResourceName(), s)
		}

	default:
		return nil, errors.Errorf("unknown resource type %s", cfg.API.ResourceType)
	}

	return nil, errors.Errorf("can't recreate resource %+v", req.Config)
}

// Validator is a resource configuration object that implements Validate.
type Validator interface {
	// Validate ensures that the object is valid and returns any implicit dependencies.
	Validate(path string) ([]string, error)
}

// ValidateConfig receives the validation request for a resource from the parent.
func (m *Module) ValidateConfig(ctx context.Context,
	req *pb.ValidateConfigRequest,
) (*pb.ValidateConfigResponse, error) {
	c, err := config.ComponentConfigFromProto(req.Config)
	if err != nil {
		return nil, err
	}

	if err := addConvertedAttributes(c); err != nil {
		return nil, errors.Wrapf(err, "unable to convert attributes for validation")
	}

	if c.ConvertedAttributes != nil {
		validator, ok := c.ConvertedAttributes.(Validator)
		if ok {
			implicitDeps, err := validator.Validate(c.Name)
			if err != nil {
				return nil, errors.Wrapf(err, "error validating resource")
			}
			return &pb.ValidateConfigResponse{Dependencies: implicitDeps}, nil
		}
	}

	// Resource configuration object does not implement Validate, but return an
	// empty response and no error to maintain backward compatibility.
	return &pb.ValidateConfigResponse{}, nil
}

// RemoveResource receives the request for resource removal.
func (m *Module) RemoveResource(ctx context.Context, req *pb.RemoveResourceRequest) (*pb.RemoveResourceResponse, error) {
	slowWatcher, slowWatcherCancel := utils.SlowGoroutineWatcher(
		30*time.Second, fmt.Sprintf("module resource %q is taking a while to remove", req.Name), m.logger)
	defer func() {
		slowWatcherCancel()
		<-slowWatcher
	}()
	m.mu.Lock()
	defer m.mu.Unlock()

	name, err := resource.NewFromString(req.Name)
	if err != nil {
		return nil, err
	}

	svc, ok := m.services[name.Subtype]
	if !ok {
		return nil, errors.Errorf("no grpc service for %+v", name)
	}
	comp := svc.Resource(name.Name)
	if err := utils.TryClose(ctx, comp); err != nil {
		m.logger.Error(err)
	}

	return &pb.RemoveResourceResponse{}, svc.Remove(name)
}

// addAPIFromRegistry adds a preregistered API (rpc Subtype) to the module's services.
func (m *Module) addAPIFromRegistry(ctx context.Context, api resource.Subtype) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.services[api]
	if ok {
		return nil
	}
	newSvc, err := subtype.New(make(map[resource.Name]interface{}))
	if err != nil {
		return err
	}
	m.services[api] = newSvc

	rs := registry.ResourceSubtypeLookup(api)
	if rs != nil && rs.RegisterRPCService != nil {
		if err := rs.RegisterRPCService(ctx, m.server, newSvc); err != nil {
			return err
		}
	}
	return nil
}

// AddModelFromRegistry adds a preregistered component or service model to the module's services.
func (m *Module) AddModelFromRegistry(ctx context.Context, api resource.Subtype, model resource.Model) error {
	err := validateRegistered(api, model)
	if err != nil {
		return err
	}

	m.mu.Lock()
	_, ok := m.services[api]
	m.mu.Unlock()
	if !ok {
		if err := m.addAPIFromRegistry(ctx, api); err != nil {
			return err
		}
	}

	creator := registry.ResourceSubtypeLookup(api)

	if creator.ReflectRPCServiceDesc == nil {
		m.logger.Errorf("rpc subtype %s doesn't contain a valid ReflectRPCServiceDesc", api)
	}
	rpcSubtype := resource.RPCSubtype{
		Subtype:      api,
		ProtoSvcName: creator.RPCServiceDesc.ServiceName,
		Desc:         creator.ReflectRPCServiceDesc,
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	m.handlers[rpcSubtype] = append(m.handlers[rpcSubtype], model)
	return nil
}

// OperationManager returns the operation manager for the module.
func (m *Module) OperationManager() *operation.Manager {
	return m.operations
}

// addConvertedAttributesToConfig uses the MapAttributeConverter to fill in the
// ConvertedAttributes field from the Attributes.
func addConvertedAttributes(cfg *config.Component) error {
	// Try to find map converter for a component.
	conv := config.FindMapConverter(cfg.API, cfg.Model)
	// If no map converter for a component exists, try to find map converter for a
	// service.
	if conv == nil {
		conv = config.FindServiceMapConverter(cfg.API, cfg.Model)
	}
	if conv != nil {
		converted, err := conv(cfg.Attributes)
		if err != nil {
			return errors.Wrapf(err, "error converting attributes for resource")
		}
		cfg.ConvertedAttributes = converted
	}

	return nil
}

// validateRegistered returns an error if the passed-in api and model have not
// yet been registered.
func validateRegistered(api resource.Subtype, model resource.Model) error {
	switch api.ResourceType {
	case resource.ResourceTypeComponent:
		creator := registry.ComponentLookup(api, model)
		if creator == nil || creator.Constructor == nil {
			return fmt.Errorf("component with API %s and model %s not yet registered",
				api, model)
		}
	case resource.ResourceTypeService:
		creator := registry.ServiceLookup(api, model)
		if creator == nil || creator.Constructor == nil {
			return fmt.Errorf("service with API %s and model %s not yet registered",
				api, model)
		}
	default:
		return errors.Errorf("unknown resource type %s", api.ResourceType)
	}

	return nil
}
