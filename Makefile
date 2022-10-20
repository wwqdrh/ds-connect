SHELL := /bin/bash -o pipfail

PREFIX			  ?= wwqdrh
TAG				  ?= dev
SHADOW_IMAGE	  =  ds-connect-shadow

shadow:
	GOARCH=amd64 GOOS=linux go build -gcflags "all=-N -l" -o artifacts/shadow/shadow-linux-amd64 cmd/shadow/main.go
	docker build -t $(PREFIX)/$(SHADOW_IMAGE):$(TAG) -f build/docker/shadow/Dockerfile .
