export NPM_UID=$(shell id -u)
export NPM_GID=$(shell id -g)

default: fmt lint install generate

build:
	go build -v ./...

install: build
	go install -v ./...

lint:
	golangci-lint run

generate:
	go generate tools/tools.go

fmt:
	gofmt -s -w -e .

test: docker-up
	TF_ACC=1 go test -v -count=1 -timeout 120m ./...

test-acc: docker-up
	TF_ACC=1 go test -v -cover -timeout 120m -count=1 ./internal/provider/tests

test-sdk: docker-up
	go test -v -cover -timeout 120m -count=1 ./internal/sdk

.PHONY: fmt lint test testacc build install generate

docker-up:
	docker compose -f docker/docker-compose.yaml up -d --quiet-pull
	while true; do if docker compose -f docker/docker-compose.yaml ps | grep -qP '(unhealthy|starting)'; then echo "Waiting for services to start up..."; sleep 5; else break; fi; done
	docker compose -f docker/docker-compose.yaml ps

docker-down:
	docker compose -f docker/docker-compose.yaml down
	rm docker/npm -rf
	docker volume prune -f
