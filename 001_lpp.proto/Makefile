GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
# INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
API_PROTO_FILES=$(shell find proto -name *.proto)



.PHONY: init
# init env
init:
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install go.unistack.org/protoc-gen-go-micro/v3@latest
	go install github.com/yusank/protoc-gen-go-http@latest
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest

# This will place four binaries in your $GOBIN;Make sure that your $GOBIN is in your $PATH.
# protoc-gen-grpc-gateway
# protoc-gen-openapiv2
# protoc-gen-go
# protoc-gen-grpc-gateway
# protoc-gen-go-grpc

.PHONY: config
# generate internal proto

.PHONY: api
# generate api proto
api:
	protoc 
	  	   -I . \
  		   --validate_out="lang=go:../generated" \
  		   --proto_path=. \
	       --proto_path=./third_party \
  		   --go_out=":../generated" \
 	       --go-http_out=paths=import:. \
 	       --go-grpc_out=paths=import:. \
 	       --go-micro_out=debug=true,,components="micro",paths=import:. \
	       $(API_PROTO_FILES)

#    --go_out=paths=import:. \


.PHONY: clean
# find ./go_proto -name "*.pb.go" | xargs rm -f
clean:
	find ./go_proto -name "*.pb.go" | xargs rm -f

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: generate
# generate
generate:
	go generate ./...

.PHONY: all
# generate all
all:
	make api;
	# make generate;

.PHONY: help
# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
