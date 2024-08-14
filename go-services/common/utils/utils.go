package utils

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/grpc/credentials"
)

func LoadCerts(certsPath, name string) (credentials.TransportCredentials, error) {
	clientCertPath := filepath.Join(certsPath, fmt.Sprintf("%s.crt", name))
	clientKeyPath := filepath.Join(certsPath, fmt.Sprintf("%s.key", name))
	caCertPath := filepath.Join(certsPath, "ca.crt")

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

	return creds, nil
}
