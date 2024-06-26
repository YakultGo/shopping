INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
 	       --go_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
	       $(API_PROTO_FILES)

.PHONY: gen
# generate
gen:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...