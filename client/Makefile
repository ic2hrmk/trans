SHELL := /bin/bash

.ONESHELL:
.PHONY: build clean

TARGET := $(if $(TARGET),$(TARGET),tclient)

build:
	source ./vendor/gocv.io/x/gocv/env.sh
	go build -o $(TARGET)

clean:
	rm -f $(TARGET)
	rm -f *.log