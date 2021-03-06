all:
	@protoc \
		-I/usr/local/include \
		-I./pb \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:./pb \
		./pb/steakhouse.proto
	@protoc \
		-I/usr/local/include \
		-I./pb \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:./pb \
		./pb/steakhouse.proto
	@protoc \
		-I/usr/local/include \
		-I./pb \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:./pb \
		./pb/steakhouse.proto
