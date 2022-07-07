SHELL := env PATH=$(PATH) /bin/bash

build-server:
	go build -o bin/server cmd/server/main.go

start-server:
	./bin/server

build-cli:
	go build -o bin/client cmd/client/main.go

start-cli:
	./bin/client
