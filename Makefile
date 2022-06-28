.PHONY: cook_env
cook_env:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GO111MODULE=off go get -d github.com/gogo/protobuf/protoc-gen-gofast
	go install github.com/gogo/protobuf/protoc-gen-gofast
	GO111MODULE=off go get -d github.com/gogo/protobuf/protoc-gen-gogo
	go install github.com/gogo/protobuf/protoc-gen-gogo
	go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.16.0
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.16.0
	GO111MODULE=off go get github.com/ethereum/go-ethereum
	cd $(GOPATH)/src/github.com/ethereum/go-ethereum && make && make devtools && cd -

.PHONY: create_schema
create_schema:
	bash ./schema.bash

.PHONY: build_proto
build_proto:
	protoc -I=shared/types --gofast_out=plugins=grpc:./ shared/types/*.proto
	protoc -I=shared/types --gogo_out=plugins=grpc:./ --grpc-gateway_opt grpc_api_configuration=./shared/types/gateway.yaml --grpc-gateway_out=allow_patch_feature=false:./ shared/types/gateway.proto

