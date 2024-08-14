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

const (
	retryDelay  = time.Second * 2
	maxAttempts = 5
)

type ChartServiceClientManager struct {
	client    pb.PlotServiceClient
	certsPath string
	once      sync.Once
	address   string
}

func (m *ChartServiceClientManager) createClient() (pb.PlotServiceClient, error) {
	var conn *grpc.ClientConn
	var err error

	creds, err := utils.LoadCerts(m.certsPath, "api_gateway_client")
	if err != nil {
		return nil, err
	}

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		conn, err = grpc.NewClient(m.address, grpc.WithTransportCredentials(creds), grpc.WithUnaryInterceptor(middleware.UnaryClientLogger(slog.Default())))
		if err == nil {
			return pb.NewPlotServiceClient(conn), nil
		}

		slog.Info("Unable to create gRPC client", "client", "pb.PlotServiceClient", "Attempt", attempts, "error", err)
		time.Sleep(retryDelay)
	}
	return nil, fmt.Errorf("failed to establish a gRPC connection to server at %s after %d attempts: %w", m.address, maxAttempts, err)
}

func (m *ChartServiceClientManager) GetClient() (pb.PlotServiceClient, error) {
	var err error
	m.once.Do(func() {
		m.client, err = m.createClient()
	})
	return m.client, err
}

func NewChartServiceClientManager(address string, certsPath string) *ChartServiceClientManager {
	return &ChartServiceClientManager{address: address, certsPath: certsPath}
}
