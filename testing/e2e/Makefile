all: cli docker

cli:
	go build -o build/e2e ./cmd/e2e

docker:
	docker build --tag ghcr.io/furyaxyz/elysium-app:current -f ../../Dockerfile ../..

PHONY: all cli docker
