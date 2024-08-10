package rpc

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log/slog"
	"ltplotter/gen/pb"
	"os"
	"path/filepath"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	clientCertPath := filepath.Join(m.certsPath, "client.crt")
	clientKeyPath := filepath.Join(m.certsPath, "client.key")
	caCertPath := filepath.Join(m.certsPath, "ca.crt")

	clientCert, err := tls.LoadX509KeyPair(clientCertPath, clientKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load client certificate: %w", err)
	}

	caCert, err := os.ReadFile(caCertPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read CA certificate: %w", err)
	}

	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert) {
		return nil, fmt.Errorf("failed to add CA certificate to pool")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      caCertPool,
	})

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		conn, err = grpc.NewClient(m.address, grpc.WithTransportCredentials(creds))
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
