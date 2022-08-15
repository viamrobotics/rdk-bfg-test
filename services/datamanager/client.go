// Package datamanager contains a gRPC based datamanager service server
package datamanager

import (
	"context"

	"github.com/edaniels/golog"
	"go.viam.com/utils/rpc"

	pb "go.viam.com/rdk/proto/api/service/datamanager/v1"
)

// client implements DataManagerServiceClient.
type client struct {
	conn   rpc.ClientConn
	client pb.DataManagerServiceClient
	logger golog.Logger
}

// NewClientFromConn constructs a new Client from connection passed in.
func NewClientFromConn(ctx context.Context, conn rpc.ClientConn, name string, logger golog.Logger) Service {
	grpcClient := pb.NewDataManagerServiceClient(conn)
	c := &client{
		conn:   conn,
		client: grpcClient,
		logger: logger,
	}
	return c
}

func (c *client) Sync(ctx context.Context) error {
	_, err := c.client.Sync(ctx, &pb.SyncRequest{})
	if err != nil {
		return err
	}
	return nil
}
