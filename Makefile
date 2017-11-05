SHELL := /bin/bash

OK_COLOR=\033[32;01m
NO_COLOR=\033[0m

.PHONY: all

$(eval SHA := $(shell git rev-parse HEAD))

all: proto

proto:
	node_modules/protobufjs/bin/pbjs -t json */*.proto > proto_schemas.json
