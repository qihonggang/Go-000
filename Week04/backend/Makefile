GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
BRANCH=$(shell git symbolic-ref -q --short HEAD)
REVISION=$(shell git rev-parse --short HEAD)
BUILD_DATE=$(shell date +%FT%T%Z)
# PROTO_FILES=$(shell find . -name *.proto)  #linux
PROTO_FILES=$(shell for /r . %%i in (*.proto) do @echo %%i) # windows
KRATOS=$(GOPATH)/src/github.com/go-kratos/kratos

.PHONY: init
init:
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
.PHONY: proto
proto:
	 protoc --go_out=. \
	 --go_opt=paths=source_relative \
	 --go-grpc_out=. \
	 --go-grpc_opt=paths=source_relative .\api\user\user.proto
.PHONY: build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION) -X main.Branch=$(BRANCH) -X main.Revision=$(REVISION) -X main.BuildDate=$(BUILD_DATE)" -o ./bin/ ./...

.PHONY: test
test:
	go test -v ./... -cover