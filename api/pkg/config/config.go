package config

import (
	"flag"
	"os"
	"strconv"
)

type Config struct {
	Port               int
	Host               string
	PGFPlotsServiceURL string
}

func LoadConfig() *Config {
	config := &Config{}

	if port, exists := os.LookupEnv("PORT"); exists {
		config.Port, _ = strconv.Atoi(port)
	} else {
		config.Port = 8080
	}

	config.Host = os.Getenv("HOST")
	if config.Host == "" {
		config.Host = "localhost"
	}

	config.PGFPlotsServiceURL = os.Getenv("PGFPLOT_SVC_URL")

	flag.IntVar(&config.Port, "port", config.Port, "TCP Port to bind server to")
	flag.StringVar(&config.Host, "host", config.Host, "Network to bind to")
	flag.StringVar(&config.PGFPlotsServiceURL, "pgfplots-url", config.PGFPlotsServiceURL, "URL of the PGFPlots service")

	flag.Parse()

	return config
}
