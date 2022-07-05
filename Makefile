
init:
	@brew install protobuf && \
	go install github.com/golang/protobuf/protoc-gen-go@v1.5.2 && \
	brew install grpcurl

proto-gen:
	@mkdir -p internal/gen/chat && \
	protoc -I./proto --go_out=plugins=grpc:./internal/gen/chat  --go_opt=paths=source_relative chat.proto