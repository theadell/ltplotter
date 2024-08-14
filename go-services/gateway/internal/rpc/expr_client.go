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

type ExpressionPlotServiceClientManager struct {
	client    pb.ExpressionPlotServiceClient
	certsPath string
	once      sync.Once
	address   string
}

func (m *ExpressionPlotServiceClientManager) createClient() (pb.ExpressionPlotServiceClient, error) {
	var conn *grpc.ClientConn
	var err error

	creds, err := utils.LoadCerts(m.certsPath, "expression_plot_server")
	if err != nil {
		return nil, err
	}

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		conn, err = grpc.NewClient(m.address, grpc.WithTransportCredentials(creds), grpc.WithUnaryInterceptor(middleware.UnaryClientLogger(slog.Default())))
		if err == nil {
			return pb.NewExpressionPlotServiceClient(conn), nil
		}

		slog.Info("Unable to create gRPC client", "client", "pb.ExpressionPlotServiceClient", "Attempt", attempts, "error", err)
		time.Sleep(retryDelay)
	}
	return nil, fmt.Errorf("failed to establish a gRPC connection to server at %s after %d attempts: %w", m.address, maxAttempts, err)
}

func (m *ExpressionPlotServiceClientManager) GetClient() (pb.ExpressionPlotServiceClient, error) {
	var err error
	m.once.Do(func() {
		m.client, err = m.createClient()
	})
	return m.client, err
}

func NewExpressionPlotServiceClientManager(address string, certsPath string) *ExpressionPlotServiceClientManager {
	return &ExpressionPlotServiceClientManager{address: address, certsPath: certsPath}
}
