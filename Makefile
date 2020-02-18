COMMIT=$(shell git rev-parse HEAD | cut -c -8 || echo dev)
LDFLAGS=-ldflags "-X github.com/martinrue/parol-api/api.Commit=${COMMIT}"

all: dev

clean:
	@rm -rf ./dist

dev: clean
	@go build -mod=vendor ${LDFLAGS} -o ./dist/api ./cmd/api

prod: clean
	@GOOS=linux GOARCH=amd64 go build -mod=vendor ${LDFLAGS} -o ./dist/api-linux-amd64 ./cmd/api

.PHONY: all clean dev prod
