.PHONY: run build
build:
	GOOS=linux GOARCH=amd64 go build -tags demo-exporter -o bin/go-exporter cmd/main.go
	docker build -f build/docker/Dockerfile -t demo-exporter  .
run:
	docker run --rm  -e MICRO_REGISTRY=mdns demo-exporter 

