CURDIR=$(shell pwd)

# Используем bin в текущей директории для установки плагинов protoc
LOCAL_BIN:=$(CURDIR)/bin
.PHONY: bin
bin:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@latest

# Добавляем bin в текущей директории в PATH при запуске protoc
PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc

# Устанавливаем proto описания google/googleapis
vendor-proto/google/api:
	git clone -b master --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/googleapis/googleapis vendor-proto/googleapis &&\
 	cd vendor-proto/googleapis &&\
	git sparse-checkout set --no-cone google/api &&\
	git checkout
	mkdir -p  vendor-proto/google
	mv vendor-proto/googleapis/google/api vendor-proto/google
	rm -rf vendor-proto/googleapis

# Устанавливаем proto описания google/protobuf
vendor-proto/google/protobuf:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
		https://github.com/protocolbuffers/protobuf vendor-proto/protobuf &&\
	cd vendor-proto/protobuf &&\
	git sparse-checkout set --no-cone src/google/protobuf &&\
	git checkout
	mkdir -p  vendor-proto/google
	mv vendor-proto/protobuf/src/google/protobuf vendor-proto/google
	rm -rf vendor-proto/protobuf

# Устанавливаем proto описания validate
vendor-proto/validate:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
		https://github.com/bufbuild/protoc-gen-validate vendor-proto/validate-repo &&\
	cd vendor-proto/validate-repo &&\
	git sparse-checkout set --no-cone validate &&\
	git checkout
	mkdir -p  vendor-proto
	mv vendor-proto/validate-repo/validate vendor-proto
	rm -rf vendor-proto/validate-repo

generate:
	rm -rf vendor-proto
	make qgenerate
qgenerate: bin vendor-proto/google/api vendor-proto/google/protobuf vendor-proto/validate
	mkdir -p internal/pb/pick-up_point
	mkdir -p api/openapi/pick-up-point
	$(PROTOC) -I api/proto/pick-up-point -I vendor-proto \
	--go_out internal/pb/pick-up_point --go_opt paths=source_relative \
	--go-grpc_out internal/pb/pick-up_point --go-grpc_opt paths=source_relative \
	--grpc-gateway_out internal/pb/pick-up_point --grpc-gateway_opt paths=source_relative \
	--validate_out="lang=go,paths=source_relative:internal/pb/pick-up_point" \
	--openapiv2_out api/openapi/pick-up-point \
	api/proto/pick-up-point/pick-up-point.proto
	mkdir -p internal/pb/order
	mkdir -p api/openapi/order
	$(PROTOC) -I api/proto/order -I vendor-proto \
	--go_out internal/pb/order --go_opt paths=source_relative \
	--go-grpc_out internal/pb/order --go-grpc_opt paths=source_relative \
	--grpc-gateway_out internal/pb/order --grpc-gateway_opt paths=source_relative \
	--openapiv2_out api/openapi/order \
    --experimental_allow_proto3_optional api/proto/order/order.proto \


