GOPATH:=$(shell go env GOPATH)

.PHONY: prerequisites
prerequisites:
	@brew install helm
	@brew install skaffold
#	@curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
#	@chmod 700 get_helm.sh
#	@./get_helm.sh
	@go install go-micro.dev/v4/cmd/micro@a40f6e8fae19018ead6778a2887c0daf970a9882

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install github.com/golang/protobuf/protoc-gen-go@latest
	@go install github.com/asim/go-micro/cmd/protoc-gen-micro/v4@latest

.PHONY: proto
proto:
	@protoc --proto_path=. --micro_out=. --go_out=:. proto/micro-server.proto

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o micro-server *.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@docker build -t micro-server:latest .
