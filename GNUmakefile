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
	TF_ACC=1 go test -v -cover -timeout 120m ./internal/provider/tests

test-sdk: docker-up
	go test -v -cover -timeout 120m ./internal/sdk

.PHONY: fmt lint test testacc build install generate

docker-up:
	docker compose -f docker/docker-compose.yaml up -d
	while true; do if docker compose -f docker/docker-compose.yaml ps | grep -q 'starting'; then echo "Waiting for services to start up..."; sleep 2; else break; fi; done
	docker compose -f docker/docker-compose.yaml ps

docker-down:
	docker compose -f docker/docker-compose.yaml down
