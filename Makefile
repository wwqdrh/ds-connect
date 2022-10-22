SHELL := /bin/bash -o pipefail

PREFIX			  ?= wwqdrh
TAG				  ?= dev
SHADOW_IMAGE	  =  ds-connect-shadow

shadow:
	GOARCH=amd64 GOOS=linux go build -gcflags "all=-N -l" -o artifacts/shadow/shadow-linux-amd64 cmd/shadow/main.go
	docker build -t $(PREFIX)/$(SHADOW_IMAGE):$(TAG) -f build/docker/shadow/Dockerfile .

shadow-test-deploy:
	docker service create --name shadow --network dev -p 10022:22 wwqdrh/ds-connect-shadow:dev
	