package rpc

import (
	"fmt"
	"log/slog"
	"ltplotter/common/middleware"
	"ltplotter/common/utils"
	"ltplotter/gen/pb"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type ManimServiceClientManager struct {
	client    pb.ManimCompilerServiceClient
	certsPath string
	once      sync.Once
	address   string
}

func (m *ManimServiceClientManager) createClient() (pb.ManimCompilerServiceClient, error) {
	var conn *grpc.ClientConn
	var err error

	creds, err := utils.LoadCerts(m.certsPath, "api_gateway_client")
	if err != nil {
		return nil, err
	}

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		conn, err = grpc.NewClient(
			m.address,
			grpc.WithTransportCredentials(creds),
			grpc.WithUnaryInterceptor(middleware.UnaryClientLogger(slog.Default())),
		)
		if err == nil {
			return pb.NewManimCompilerServiceClient(conn), nil
		}

		slog.Info(
			"Unable to create gRPC client",
			"client",
			"pb.ManimCompilerServiceClient",
			"Attempt",
			attempts,
			"error",
			err,
		)
		time.Sleep(retryDelay)
	}
	return nil, fmt.Errorf(
		"failed to establish a gRPC connection to server at %s after %d attempts: %w",
		m.address,
		maxAttempts,
		err,
	)
}

func (m *ManimServiceClientManager) GetClient() (pb.ManimCompilerServiceClient, error) {
	var err error
	m.once.Do(func() {
		m.client, err = m.createClient()
	})
	return m.client, err
}

func NewManimServiceClientManager(address string, certsPath string) *ManimServiceClientManager {
	return &ManimServiceClientManager{address: address, certsPath: certsPath}
}
