package utils

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"ltplotter/common/tex"
	"ltplotter/gen/pb"
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

func EscapeExprPlotRequest(req *pb.ExprPlotRequest) {
	req.Title = tex.Escape(req.Title)
	req.XLabel = tex.Escape(req.XLabel)
	req.YLabel = tex.Escape(req.YLabel)

	for _, plotExpr := range req.Plots {
		// TODO: Create a parser to build a syntax tree and validate that the expression is valid
		// plotExpr.Expression = tex.Escape(plotExpr.Expression)
		plotExpr.Domain = tex.Escape(plotExpr.Domain)
		plotExpr.Color = tex.Escape(plotExpr.Color)
		plotExpr.LineStyle = tex.Escape(plotExpr.LineStyle)
		plotExpr.LineWidth = tex.Escape(plotExpr.LineWidth)
		plotExpr.Legend = tex.Escape(plotExpr.Legend)
	}
}
