generate:
	protoc --proto_path=proto \
		--go_out=gen/auth --go_opt=paths=source_relative \
		--go-grpc_out=gen/auth --go-grpc_opt=paths=source_relative \
		proto/auth.proto
