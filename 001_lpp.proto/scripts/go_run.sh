find ./go_proto -name "*.pb.go" | xargs rm -f
# protoc --proto_path=. --go_out=paths=source_relative:./go_proto --go-grpc_out=paths=source_relative:./go_proto proto/*.proto

protoc --proto_path=. --go_out=paths=import:. --go-grpc_out=paths=import:. proto/*.proto

# protoc --proto_path=. --go_out=paths=source_relative:./go_proto --go-grpc_out=paths=source_relative:./go_proto proto/algo_service.proto
# protoc --proto_path=. --go_out=paths=source_relative:./go_proto --go-grpc_out=paths=source_relative:./go_proto proto/quote_service.proto
# protoc --proto_path=. --go_out=paths=source_relative:./go_proto --go-grpc_out=paths=source_relative:./go_proto proto/risk_service.proto
# protoc --proto_path=. --go_out=paths=source_relative:./go_proto --go-grpc_out=paths=source_relative:./go_proto proto/trade_service.proto
