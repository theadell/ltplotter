package config

import (
	"flag"
	"os"
	"strconv"
)

type Config struct {
	Port      int
	Host      string
	CertsPath string
}

func LoadConfig() *Config {
	config := &Config{}

	if port, exists := os.LookupEnv("PORT"); exists {
		config.Port, _ = strconv.Atoi(port)
	} else {
		config.Port = 50051
	}

	config.CertsPath = os.Getenv("CERTS_PATH")
	if config.CertsPath == "" {
		config.CertsPath = "/app/certs"
	}

	config.Host = os.Getenv("HOST")
	if config.Host == "" {
		config.Host = "localhost"
	}

	flag.IntVar(&config.Port, "port", config.Port, "TCP Port to bind server to")
	flag.StringVar(&config.Host, "host", config.Host, "Network to bind to")
	flag.StringVar(&config.CertsPath, "certs-path", config.CertsPath, "Path to the directory containing the TLS certificates and keys for mutual TLS (mTLS) authentication.")

	flag.Parse()

	return config
}
