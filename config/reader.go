package config

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/edaniels/golog"
	"github.com/mitchellh/copystructure"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"go.viam.com/utils"
)

// An AttributeConverter converts a single attribute into a possibly
// different representation.
type AttributeConverter func(val interface{}) (interface{}, error)

// An AttributeMapConverter converts an attribute map into a possibly
// different representation.
type AttributeMapConverter func(attributes AttributeMap) (interface{}, error)

// A ComponentAttributeConverterRegistration describes how to convert a specific attribute
// for a model of a type of component.
type ComponentAttributeConverterRegistration struct {
	CompType ComponentType
	Model    string
	Attr     string
	Conv     AttributeConverter
}

// A ComponentAttributeMapConverterRegistration describes how to convert all attributes
// for a model of a type of component.
type ComponentAttributeMapConverterRegistration struct {
	CompType ComponentType
	Model    string
	Conv     AttributeMapConverter
	RetType  interface{} // the shape of what is converted to
}

// A ServiceAttributeMapConverterRegistration describes how to convert all attributes
// for a model of a type of service.
type ServiceAttributeMapConverterRegistration struct {
	SvcType ServiceType
	Conv    AttributeMapConverter
	RetType interface{} // the shape of what is converted to
}

var (
	componentAttributeConverters    = []ComponentAttributeConverterRegistration{}
	componentAttributeMapConverters = []ComponentAttributeMapConverterRegistration{}
	serviceAttributeMapConverters   = []ServiceAttributeMapConverterRegistration{}
)

// RegisterComponentAttributeConverter associates a component type and model with a way to convert a
// particular attribute name.
func RegisterComponentAttributeConverter(compType ComponentType, model, attr string, conv AttributeConverter) {
	componentAttributeConverters = append(componentAttributeConverters, ComponentAttributeConverterRegistration{compType, model, attr, conv})
}

// RegisterComponentAttributeMapConverter associates a component type and model with a way to convert all attributes.
func RegisterComponentAttributeMapConverter(compType ComponentType, model string, conv AttributeMapConverter, retType interface{}) {
	if retType == nil {
		panic("retType should not be nil")
	}
	componentAttributeMapConverters = append(
		componentAttributeMapConverters,
		ComponentAttributeMapConverterRegistration{compType, model, conv, retType},
	)
}

// TransformAttributeMapToStruct uses an attribute map to transform attributes to the perscribed format.
func TransformAttributeMapToStruct(to interface{}, attributes AttributeMap) (interface{}, error) {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json", Result: to})
	if err != nil {
		return nil, err
	}
	if err := decoder.Decode(attributes); err != nil {
		return nil, err
	}
	return to, nil
}

// RegisterServiceAttributeMapConverter associates a service type with a way to convert all attributes.
func RegisterServiceAttributeMapConverter(svcType ServiceType, conv AttributeMapConverter, retType interface{}) {
	if retType == nil {
		panic("retType should not be nil")
	}
	serviceAttributeMapConverters = append(serviceAttributeMapConverters, ServiceAttributeMapConverterRegistration{svcType, conv, retType})
}

// RegisteredComponentAttributeConverters returns a copy of the registered component attribute converters.
func RegisteredComponentAttributeConverters() []ComponentAttributeConverterRegistration {
	copied, err := copystructure.Copy(componentAttributeConverters)
	if err != nil {
		panic(err)
	}
	return copied.([]ComponentAttributeConverterRegistration)
}

// RegisteredComponentAttributeMapConverters returns a copy of the registered component attribute converters.
func RegisteredComponentAttributeMapConverters() []ComponentAttributeMapConverterRegistration {
	copied, err := copystructure.Copy(componentAttributeMapConverters)
	if err != nil {
		panic(err)
	}
	return copied.([]ComponentAttributeMapConverterRegistration)
}

// RegisteredServiceAttributeMapConverters returns a copy of the registered component attribute converters.
func RegisteredServiceAttributeMapConverters() []ServiceAttributeMapConverterRegistration {
	copied, err := copystructure.Copy(serviceAttributeMapConverters)
	if err != nil {
		panic(err)
	}
	return copied.([]ServiceAttributeMapConverterRegistration)
}

func findConverter(compType ComponentType, model, attr string) AttributeConverter {
	for _, r := range componentAttributeConverters {
		if r.CompType == compType && r.Model == model && r.Attr == attr {
			return r.Conv
		}
	}
	return nil
}

func findMapConverter(compType ComponentType, model string) AttributeMapConverter {
	for _, r := range componentAttributeMapConverters {
		if r.CompType == compType && r.Model == model {
			return r.Conv
		}
	}
	return nil
}

func findServiceMapConverter(svcType ServiceType) AttributeMapConverter {
	for _, r := range serviceAttributeMapConverters {
		if r.SvcType == svcType {
			return r.Conv
		}
	}
	return nil
}

func openSub(original, target string) (*os.File, error) {
	//nolint:gosec
	targetFile, err := os.Open(target)
	if err == nil {
		return targetFile, nil
	}

	// try finding it through the original path
	//nolint:gosec
	targetFile, err = os.Open(path.Join(path.Dir(original), target))
	if err == nil {
		return targetFile, nil
	}

	return nil, errors.Errorf("cannot find file: %s", target)
}

func loadSubFromFile(original, cmd string) (interface{}, bool, error) {
	if !strings.HasPrefix(cmd, "$load{") {
		return cmd, false, nil
	}

	cmd = cmd[6:]             // [$load{|...]
	cmd = cmd[0 : len(cmd)-1] // [...|}]

	subFile, err := openSub(original, cmd)
	if err != nil {
		return cmd, false, err
	}
	defer utils.UncheckedErrorFunc(subFile.Close)

	var sub map[string]interface{}
	decoder := json.NewDecoder(subFile)
	if err := decoder.Decode(&sub); err != nil {
		return nil, false, err
	}
	return sub, true, nil
}

var (
	defaultCloudBasePath = "https://app.viam.com"
	defaultCloudPath     = defaultCloudBasePath + "/api/json1/config"
	defaultCoudLogPath   = defaultCloudBasePath + "/api/json1/log"
)

const (
	cloudConfigSecretField           = "Secret"
	cloudConfigUserInfoField         = "User-Info"
	cloudConfigUserInfoHostField     = "host"
	cloudConfigUserInfoOSField       = "os"
	cloudConfigUserInfoLocalIPsField = "ips"
)

// CreateCloudRequest makes a request to fetch the robot config
// from a cloud endpoint.
func CreateCloudRequest(ctx context.Context, cloudCfg *Cloud) (*http.Request, error) {
	if cloudCfg.Path == "" {
		cloudCfg.Path = defaultCloudPath
	}
	if cloudCfg.LogPath == "" {
		cloudCfg.LogPath = defaultCoudLogPath
	}

	url := fmt.Sprintf("%s?id=%s", cloudCfg.Path, cloudCfg.ID)

	r, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error creating request for %s", url)
	}
	r.Header.Set(cloudConfigSecretField, cloudCfg.Secret)

	userInfo := map[string]interface{}{}
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	userInfo[cloudConfigUserInfoHostField] = hostname
	userInfo[cloudConfigUserInfoOSField] = runtime.GOOS

	ips, err := utils.GetAllLocalIPv4s()
	if err != nil {
		return nil, err
	}
	userInfo[cloudConfigUserInfoLocalIPsField] = ips

	userInfoBytes, err := json.Marshal(userInfo)
	if err != nil {
		return nil, err
	}

	r.Header.Set(cloudConfigUserInfoField, string(userInfoBytes))

	return r, nil
}

var viamDotDir = filepath.Join(os.Getenv("HOME"), ".viam")

func getCloudCacheFilePath(id string) string {
	return filepath.Join(viamDotDir, fmt.Sprintf("cached_cloud_config_%s.json", id))
}

func openFromCache(id string) (io.ReadCloser, error) {
	return os.Open(getCloudCacheFilePath(id))
}

func storeToCache(id string, cfg *Config) error {
	if err := os.MkdirAll(viamDotDir, 0o700); err != nil {
		return err
	}
	md, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	//nolint:gosec
	return ioutil.WriteFile(getCloudCacheFilePath(id), md, 0o640)
}

// ReadFromCloud fetches a robot config from the cloud based
// on the given config.
func ReadFromCloud(ctx context.Context, cloudCfg *Cloud, readFromCache bool) (*Config, error) {
	cloudReq, err := CreateCloudRequest(ctx, cloudCfg)
	if err != nil {
		return nil, err
	}

	var client http.Client
	defer client.CloseIdleConnections()
	resp, err := client.Do(cloudReq)

	var configReader io.ReadCloser
	if err == nil {
		if resp.StatusCode != http.StatusOK {
			defer func() {
				utils.UncheckedError(resp.Body.Close())
			}()
			rd, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			if len(rd) != 0 {
				return nil, errors.Errorf("unexpected status %d: %s", resp.StatusCode, string(rd))
			}
			return nil, errors.Errorf("unexpected status %d", resp.StatusCode)
		}
		configReader = resp.Body
	} else {
		if !readFromCache {
			return nil, err
		}
		var urlErr *url.Error
		if !errors.As(err, &urlErr) || urlErr.Temporary() {
			return nil, err
		}
		var cacheErr error
		configReader, cacheErr = openFromCache(cloudCfg.ID)
		if cacheErr != nil {
			if os.IsNotExist(cacheErr) {
				return nil, err
			}
			return nil, cacheErr
		}
	}
	defer utils.UncheckedErrorFunc(configReader.Close)

	// read the actual config and do not make a cloud request again to avoid
	// infinite recursion.
	cfg, err := fromReader(ctx, "", configReader, true)
	if err != nil {
		return nil, err
	}
	if cfg.Cloud == nil {
		return nil, errors.New("expected config to have cloud section")
	}
	self := cfg.Cloud.Self
	signalingAddress := cfg.Cloud.SignalingAddress
	*cfg.Cloud = *cloudCfg
	cfg.Cloud.Self = self
	cfg.Cloud.SignalingAddress = signalingAddress

	if err := storeToCache(cloudCfg.ID, cfg); err != nil {
		golog.Global.Errorw("failed to cache config", "error", err)
	}

	return cfg, nil
}

// Read reads a config from the given file.
func Read(ctx context.Context, filePath string) (*Config, error) {
	//nolint:gosec
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer utils.UncheckedErrorFunc(file.Close)

	return FromReader(ctx, filePath, file)
}

// FromReader reads a config from the given reader and specifies
// where, if applicable, the file the reader originated from.
func FromReader(ctx context.Context, originalPath string, r io.Reader) (*Config, error) {
	return fromReader(ctx, originalPath, r, false)
}

func fromReader(ctx context.Context, originalPath string, r io.Reader, skipCloud bool) (*Config, error) {
	cfg := Config{
		ConfigFilePath: originalPath,
	}

	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, errors.Wrap(err, "cannot parse config")
	}
	if err := cfg.Ensure(skipCloud); err != nil {
		return nil, err
	}

	if !skipCloud && cfg.Cloud != nil {
		return ReadFromCloud(ctx, cfg.Cloud, true)
	}

	for idx, c := range cfg.Components {
		conv := findMapConverter(c.Type, c.Model)
		// inner attributes may have their own converters, and file substitutions
		for k, v := range c.Attributes {
			s, ok := v.(string)
			if ok {
				cfg.Components[idx].Attributes[k] = os.ExpandEnv(s)
				var loaded bool
				var err error
				v, loaded, err = loadSubFromFile(originalPath, s)
				if err != nil {
					return nil, err
				}
				if loaded {
					cfg.Components[idx].Attributes[k] = v
				}
			}

			aconv := findConverter(c.Type, c.Model, k)
			if aconv == nil {
				continue
			}

			n, err := aconv(v)
			if err != nil {
				return nil, errors.Wrapf(err, "error converting attribute for (%s, %s, %s)", c.Type, c.Model, k)
			}
			cfg.Components[idx].Attributes[k] = n
		}
		if conv == nil {
			continue
		}

		converted, err := conv(c.Attributes)
		if err != nil {
			return nil, errors.Wrapf(err, "error converting attributes for (%s, %s)", c.Type, c.Model)
		}
		cfg.Components[idx].Attributes = nil
		cfg.Components[idx].ConvertedAttributes = converted
	}

	for idx, c := range cfg.Services {
		conv := findServiceMapConverter(c.Type)
		if conv == nil {
			continue
		}

		converted, err := conv(c.Attributes)
		if err != nil {
			return nil, errors.Wrapf(err, "error converting attributes for %s", c.Type)
		}
		cfg.Services[idx].Attributes = nil
		cfg.Services[idx].ConvertedAttributes = converted
	}

	if err := cfg.Ensure(skipCloud); err != nil {
		return nil, err
	}
	return &cfg, nil
}
