ENTRY_POINT=cmd/app

proto-generate:
	protoc --proto_path=proto \
		--go_out=gen/auth --go_opt=paths=source_relative \
		--go-grpc_out=gen/auth --go-grpc_opt=paths=source_relative \
		proto/auth.proto

generate-key:
	go run cmd/generate-key/main.go

start-grpc:
	go run cmd/app/main.go

start-grpc-dev:
	go run cmd/app/main.go --dev