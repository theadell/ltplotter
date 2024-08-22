package utils

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log/slog"
	"ltplotter/common/tex"
	"ltplotter/gen/pb"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"google.golang.org/grpc/credentials"
)

var validLineWidthFormat = regexp.MustCompile(`^(\d+)(\s*pt)?$`)

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
		plotExpr.Color = tex.ColorToTikzColor(plotExpr.Color)
		plotExpr.LineStyle = validateAndConvertLineStyle(plotExpr.LineStyle)
		plotExpr.LineWidth = MapWidthToTikz(plotExpr.LineWidth)
		plotExpr.Legend = tex.Escape(plotExpr.Legend)
	}
}

func MapWidthToTikz(input string) string {
	input = strings.TrimSpace(input)

	if validLineWidthFormat.MatchString(input) {
		matches := validLineWidthFormat.FindStringSubmatch(input)
		return fmt.Sprintf("%spt", matches[1])
	}
	slog.Warn("invalid or mallformatted width", "input", input, "fallback", "1pt")
	return "1pt"
}

var validTikzLineStyles = map[string]bool{
	"solid":                 true,
	"dotted":                true,
	"densely dotted":        true,
	"loosely dotted":        true,
	"dashed":                true,
	"densely dashed":        true,
	"loosely dashed":        true,
	"dashdotted":            true,
	"densely dashdotted":    true,
	"loosely dashdotted":    true,
	"dashdotdotted":         true,
	"densely dashdotdotted": true,
	"loosely dashdotdotted": true,
}

func validateAndConvertLineStyle(input string) string {
	input = strings.TrimSpace(strings.ToLower(input))

	if validTikzLineStyles[input] {
		return input
	}
	slog.Warn("invalid or mallformatted line style", "input", input, "fallback", "solid")
	return "solid"
}
