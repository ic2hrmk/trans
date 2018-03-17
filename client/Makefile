SHELL := /bin/bash

.ONESHELL:
.PHONY: build build-raspberry clean clean-raspberry

TARGET_NAME := $(if $(TARGET_NAME),$(TARGET_NAME),tclient)

build:
	source ./vendor/gocv.io/x/gocv/env.sh
	go build -o $(TARGET_NAME)

build-raspberry:
	source ./vendor/gocv.io/x/gocv/env.sh
	$(shell export GOOS="linux"; export GOARCH="arm"; go build -v -o $(TARGET_NAME)-$(GOOS)-$(GOARCH))

clean: clean-log
	rm -f $(TARGET_NAME)

clean-rasp: clean-log
	rm -f $(TARGET_NAME)-rasp

## -- Multi usage targets

clean-log:
	rm -f *.log