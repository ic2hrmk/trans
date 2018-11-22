SHELL := /bin/bash

.ONESHELL:
.PHONY: build clean

TARGET_NAME := $(if $(TARGET_NAME),$(TARGET_NAME),tclient)
VERSION = 0.1.0
GITHASH = $(shell git rev-parse --short HEAD)
PROJECT = trans/client

LDFLAGS += -X $(PROJECT)/app/config.Version=$(VERSION)-$(GITHASH)

RASP_SYSARCH = arm

build:
	source ./vendor/gocv.io/x/gocv/env.sh
	go build -ldflags "$(LDFLAGS)" -a -o $(TARGET_NAME) main.go

clean: clean-log
	rm -f $(TARGET_NAME)

## -- Multi usage targets

clean-log:
	rm -f *.log