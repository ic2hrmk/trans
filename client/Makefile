SHELL := /bin/bash

.ONESHELL:
.PHONY: build build-via-docker clean

TARGET_NAME := $(if $(TARGET_NAME),$(TARGET_NAME),tclient)
VERSION = 0.1.0
GITHASH = $(shell git rev-parse --short HEAD)
PROJECT = trans/client

LDFLAGS += -X $(PROJECT)/app/config.Version=$(VERSION)-$(GITHASH)

## Only for machines with OpenCV installed
build:
	source ./vendor/gocv.io/x/gocv/env.sh
	go build -ldflags "$(LDFLAGS)" -a -o $(TARGET_NAME) main.go

## On any machine with docker
build-via-docker:
	docker build -t "$(TARGET_NAME)" .
	docker run --device=/dev/video0 --rm $(TARGET_NAME)

clean:
	rm -f $(TARGET_NAME)
