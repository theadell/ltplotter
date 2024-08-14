PROTO_OUT_DIR = gen/pb
PROTO_SRC_DIR = ../proto3

# Proto files to be processed
PROTO_FILES = $(wildcard $(PROTO_SRC_DIR)/*.proto)

all: gen build

$(PROTO_OUT_DIR):
	mkdir -p $(PROTO_OUT_DIR)

# Target to generate protobuf files
gen: $(PROTO_OUT_DIR)
	protoc --proto_path=$(PROTO_SRC_DIR) \
	       --go_out=$(PROTO_OUT_DIR) \
	       --go_opt=paths=source_relative \
	       --go-grpc_out=$(PROTO_OUT_DIR) \
	       --go-grpc_opt=paths=source_relative \
	       $(PROTO_FILES)

build:
	go build -o bin/server cmd/server/main.go 

run:
	go run cmd/server/main.go 

clean:
	rm -rf $(PROTO_OUT_DIR)

.PHONY: all gen build run clean