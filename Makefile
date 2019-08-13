COMMIT=$(shell git rev-parse HEAD | cut -c -8 || echo dev)
LDFLAGS=-ldflags "-X github.com/martinrue/parol-api/api.Commit=${COMMIT}"

INFRA_DEPLOY_DIR:=../deploy/assets

build:
	@go build -mod=vendor ${LDFLAGS} -o ./dist/api ./cmd/api

build-linux:
	@GOOS=linux GOARCH=amd64 go build -mod=vendor ${LDFLAGS} -o ./dist/api-linux-amd64 ./cmd/api

dist: build-linux
	@cp ./dist/api-linux-amd64 $(INFRA_DEPLOY_DIR)/api

clean:
	@rm -rf ./dist

.PHONY: build build-linux dist clean
