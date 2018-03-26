SHELL := /bin/bash

.ONESHELL:
.PHONY: build build-raspberry clean clean-raspberry

TARGET_NAME := $(if $(TARGET_NAME),$(TARGET_NAME),tclient)
VERSION = 0.1.0
GITHASH = $(shell git rev-parse --short HEAD)
PROJECT = trans/client

LDFLAGS += -X $(PROJECT)/app/config.Version=$(VERSION)-$(GITHASH)

RASP_SYSARCH = arm

build:
	source ./vendor/gocv.io/x/gocv/env.sh
	go build -ldflags "$(LDFLAGS)" -a -o $(TARGET_NAME) main.go

build-raspberry:
	source ./vendor/gocv.io/x/gocv/env.sh
	GOOS=linux GOARCH=$(RASP_SYSARCH) \
			go build -ldflags "$(LDFLAGS)" -a -o $(TARGET_NAME)_rasp main.go

clean: clean-log
	rm -f $(TARGET_NAME)

clean-rasp: clean-log
	rm -f $(TARGET_NAME)-rasp

## -- Multi usage targets

clean-log:
	rm -f *.log