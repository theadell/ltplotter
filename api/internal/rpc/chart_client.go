package rpc

import (
	"log"
	"ltplotter/gen/pb"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	retryDelay  = time.Second * 2
	maxAttempts = 5
)

type ChartServiceClientManager struct {
	client  pb.PlotServiceClient
	err     error
	once    sync.Once
	address string
}

func (m *ChartServiceClientManager) createClient() {
	var conn *grpc.ClientConn
	var err error

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		conn, err = grpc.NewClient(m.address, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err == nil {
			m.client = pb.NewPlotServiceClient(conn)
			return
		}

		log.Printf("Attempt %d: Unable to create client: %v", attempts, err)
		time.Sleep(retryDelay)
	}
	m.err = err
}

func (m *ChartServiceClientManager) GetClient() (pb.PlotServiceClient, error) {
	m.once.Do(m.createClient)
	return m.client, m.err
}

func NewChartServiceClientManager(address string) *ChartServiceClientManager {
	return &ChartServiceClientManager{address: address}
}
