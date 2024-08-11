package config

import (
	"flag"
	"os"
	"strconv"
)

type Config struct {
	Port                 int
	Host                 string
	PGFPlotsServiceURL   string
	PGFPlotServiceURLRPC string
	ExprPlotServiceUrl   string
	CertsPath            string
}

func LoadConfig() *Config {
	config := &Config{}

	if port, exists := os.LookupEnv("PORT"); exists {
		config.Port, _ = strconv.Atoi(port)
	} else {
		config.Port = 8080
	}

	config.CertsPath = os.Getenv("CERTS_PATH")
	if config.CertsPath == "" {
		config.CertsPath = "/app/certs"
	}

	config.Host = os.Getenv("HOST")
	if config.Host == "" {
		config.Host = "localhost"
	}

	config.PGFPlotsServiceURL = os.Getenv("PGFPLOT_SVC_URL")
	config.PGFPlotServiceURLRPC = os.Getenv("PGFPLOT_SVC_URL_RPC")

	config.ExprPlotServiceUrl = os.Getenv("EXPR_PLOT_SVC_URL")
	if config.ExprPlotServiceUrl == "" {
		config.ExprPlotServiceUrl = "expr-plot-service:50051"
	}
	flag.IntVar(&config.Port, "port", config.Port, "TCP Port to bind server to")
	flag.StringVar(&config.Host, "host", config.Host, "Network to bind to")
	flag.StringVar(&config.PGFPlotsServiceURL, "pgfplots-url", config.PGFPlotsServiceURL, "URL of the PGFPlots service")
	flag.StringVar(&config.ExprPlotServiceUrl, "expr-plot-url", config.ExprPlotServiceUrl, "URL of the ExpressionPlot service")

	flag.StringVar(&config.CertsPath, "certs-path", config.CertsPath, "Path to the directory containing the TLS certificates and keys for mutual TLS (mTLS) authentication.")

	flag.Parse()

	return config
}
