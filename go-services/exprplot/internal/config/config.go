package config

import (
	"flag"
	"os"
	"strconv"
)

type Config struct {
	Port        int
	Host        string
	CertsPath   string
	Template    string
	LatexEngine string
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

	config.Template = os.Getenv("TEMPLATE")
	if config.Template == "" {
		config.Template = "expr_plot_template.go.tex"
	}
	config.LatexEngine = os.Getenv("LATEX_ENGINE")
	if config.LatexEngine == "" {
		config.LatexEngine = "pdflatex"
	}

	flag.IntVar(&config.Port, "port", config.Port, "TCP Port to bind server to")
	flag.StringVar(&config.Host, "host", config.Host, "Network to bind to")
	flag.StringVar(&config.CertsPath, "certs-path", config.CertsPath, "Path to the directory containing the TLS certificates and keys for mutual TLS (mTLS) authentication.")
	flag.StringVar(&config.Template, "template", config.Template, "Path to Latex Template.")
	flag.StringVar(&config.LatexEngine, "latex-engine", config.LatexEngine, "Latex Engine.")

	flag.Parse()

	return config
}
