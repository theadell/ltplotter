# LTplotter 

LTplotter is a web app for generating publication-quality plots and graphs mathematical expressions and numerical data using `PGF/TikZ` and `LaTeX` as the backend. LTplotter allows you to create high-quality visualizations without the tedious interaction with LaTeX.

## Getting Started

LTplotter is written in Go for the backend and Vue 3 with Vuetify 3 for the frontend. To quickly get started with LTplotter, you can use Docker for containerized deployment. 

### Quick Start with Docker

To run LTplotter using Docker:

```sh
docker-compose up
```

### Local Development Setup

#### Prerequisites
- **Go** : Required for building and running the backend services.
- **node** and **npm**: for running the frontend
- **Docker**: For containerized deployment.
- **Make**: For managing build tasks.
- **Tectonic or pdflatex**: Required for generating LaTeX documents. Tectonic is recommended for local development.
#### Installation
- Generate Certificates: Generate necessary certificates by running:
    ```./generate_certs.sh```
- Start Backend Services: Navigate to the `go-services` directory and use Make to start the gateway and plotting service
    ```sh 
    cd go-services
    make run-gateway
    make run-plotexpr
    ```
- Start Frontend
    ```sh
    cd frontend
    npm install
    npm run dev
    ```