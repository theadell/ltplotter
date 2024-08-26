PROTO_OUT_DIR = gen/pb
PROTO_SRC_DIR = ../proto3
JSON_SCHEMA_OUT_DIR = gen/jsonschema
BIN_DIR = bin
DOCKER_REPO = adelll

LOCAL_CERTS_PATH=./../certs
LO=127.0.0.1
EXPR_SVC_PORT=50051
GATEWAY_SVC_PORT=8080
# Proto files to be processed
PROTO_FILES = $(wildcard $(PROTO_SRC_DIR)/*.proto)
SERVICES = gateway exprplot

all: clean gen jsonschema build

$(PROTO_OUT_DIR):
	mkdir -p $(PROTO_OUT_DIR)

$(JSON_SCHEMA_OUT_DIR):
	mkdir -p $(JSON_SCHEMA_OUT_DIR)

$(BIN_DIR):
	mkdir -p $(BIN_DIR)


gen: $(PROTO_OUT_DIR) $(JSON_SCHEMA_OUT_DIR)
	protoc --proto_path=$(PROTO_SRC_DIR) \
	       --go_out=$(PROTO_OUT_DIR) \
	       --go_opt=paths=source_relative \
	       --go-grpc_out=$(PROTO_OUT_DIR) \
	       --go-grpc_opt=paths=source_relative \
	       $(PROTO_FILES)
	protoc --proto_path=$(PROTO_SRC_DIR) \
	       --jsonschema_out=$(JSON_SCHEMA_OUT_DIR) \
	       $(PROTO_FILES)

# Build all services
build-all: $(foreach service,$(SERVICES),build-$(service))

test: 
	go test -v ./...

# Build individual services
$(foreach service,$(SERVICES),build-$(service)):
	@$(eval SERVICE_DIR = $(subst build-,,$@))
	@echo "Building $(SERVICE_DIR)..."
	go build -o $(BIN_DIR)/$(SERVICE_DIR) ./$(SERVICE_DIR)/cmd/server

run-gateway:
	go run ./gateway/cmd/server \
		-port ${GATEWAY_SVC_PORT} \
		-host ${LO} \
		-certs-path $(LOCAL_CERTS_PATH) \
		-expr-plot-url ${LO}:${EXPR_SVC_PORT}

run-exprplot:
	go run ./exprplot/cmd/server \
		-port ${EXPR_SVC_PORT} \
		-host ${LO} \
		-certs-path $(LOCAL_CERTS_PATH) \
		-template ./exprplot/static/expr_plot.go.tex \
		-latex-engine tectonic

clean:
	rm -rf $(PROTO_OUT_DIR) $(JSON_SCHEMA_OUT_DIR) $(BIN_DIR)

# Build Docker images for all services
docker-build-all: $(foreach service,$(SERVICES),docker-build-$(service))

# Build Docker images for individual services
$(foreach service,$(SERVICES),docker-build-$(service)):
	@$(eval SERVICE_NAME = $(subst docker-build-,,$@))
	@echo "Building Docker image for $(SERVICE_NAME)..."
	docker build -f $(SERVICE_NAME).Dockerfile -t $(DOCKER_REPO)/ltp-$(SERVICE_NAME):latest .

# Push Docker images for all services
docker-push-all: $(foreach service,$(SERVICES),docker-push-$(service))

# Push Docker images for individual services
$(foreach service,$(SERVICES),docker-push-$(service)):
	@$(eval SERVICE_NAME = $(subst docker-push-,,$@))
	@echo "Pushing Docker image for $(SERVICE_NAME) to Docker Hub..."
	docker push $(DOCKER_REPO)/ltp-$(SERVICE_NAME):latest


# Phony targets
.PHONY: all gen build-all test clean $(foreach service,$(SERVICES),build-$(service) run-gateway run-exprplot docker-build-$(service) docker-push-$(service)) docker-build-all docker-push-all
