gen-proto:
	protoc -I . --grpc-gateway_out ./generated_proto \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --proto_path=./proto --go_out=./generated_proto --go_opt=paths=source_relative \
    --go-grpc_out=./generated_proto --go-grpc_opt=paths=source_relative \
    services.proto

run-server:
	go run cmd/server/main.go