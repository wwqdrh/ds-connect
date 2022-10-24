SHELL := /bin/bash -o pipefail

PREFIX			  ?= wwqdrh
TAG				  ?= dev
SHADOW_IMAGE	  =  ds-connect-shadow
SHADOW_CONTROLLER_IMAGE = ds-connect-shadow-controller

shadow:
	GOARCH=amd64 GOOS=linux go build -gcflags "all=-N -l" -o artifacts/shadow/shadow-linux-amd64 cmd/shadow/main.go
	docker build -t $(PREFIX)/$(SHADOW_IMAGE):$(TAG) -f build/docker/shadow/Dockerfile .

shadow-controller:
	GOARCH=amd64 GOOS=linux go build -gcflags "all=-N -l" -o artifacts/shadow-controller/shadow-controller-linux-amd64 cmd/shadow-controller/main.go
	docker build -t $(PREFIX)/$(SHADOW_CONTROLLER_IMAGE):$(TAG) -f build/docker/shadow-controller/Dockerfile .
