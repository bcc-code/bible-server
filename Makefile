.PHONY: proto

proto: proto/*.proto
	protoc -I=./proto/ \
		--go_out=src/proto --go_opt=paths=source_relative \
		--go-grpc_out=src/proto --go-grpc_opt=paths=source_relative \
		--csharp_out=./csharp/proto --csharp_opt=file_extension=.g.cs \
		--csharp-grpc_out=./csharp/proto \
		./proto/main.proto
