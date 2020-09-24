package gateway

import (
	"context"
	"errors"
	"time"

	"github.com/zeebe-io/zeebe/clients/go/pkg/pb"
	"github.com/zeebe-io/zeebe/clients/go/pkg/zbc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const DefaultKeepAlive = 45 * time.Second

// noopCredentialsProvider implements the CredentialsProvider interface but doesn't modify the authorization headers and
// doesn't retry requests in case of failure.
type noopCredentialsProvider struct{}

// ApplyCredentials does nothing.
func (noopCredentialsProvider) ApplyCredentials(_ context.Context, _ map[string]string) error {
	return nil
}

// ShouldRetryRequest always returns false.
func (noopCredentialsProvider) ShouldRetryRequest(_ context.Context, _ error) bool {
	return false
}

func NewClient(config *zbc.ClientConfig) (pb.GatewayClient, error) {

	err := configureConnectionSecurity(config)
	if err != nil {
		return nil, err
	}

	err = configureCredentialsProvider(config)
	if err != nil {
		return nil, err
	}

	err = configureKeepAlive(config)
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(config.GatewayAddress, config.DialOpts...)
	if err != nil {
		return nil, err
	}
	return pb.NewGatewayClient(conn), nil
}

func configureCredentialsProvider(config *zbc.ClientConfig) error {
	config.CredentialsProvider = &noopCredentialsProvider{}
	return nil
}

func configureConnectionSecurity(config *zbc.ClientConfig) error {
	config.DialOpts = append(config.DialOpts, grpc.WithInsecure())
	return nil
}

func configureKeepAlive(config *zbc.ClientConfig) error {
	keepAlive := DefaultKeepAlive

	if config.KeepAlive < time.Duration(0) {
		return errors.New("keep alive must be a positive duration")
	} else if config.KeepAlive != time.Duration(0) {
		keepAlive = config.KeepAlive
	}
	config.DialOpts = append(config.DialOpts, grpc.WithKeepaliveParams(keepalive.ClientParameters{Time: keepAlive}))

	return nil
}
