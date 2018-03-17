SHELL := /bin/bash

.ONESHELL:
.PHONY: build clean

TARGET_NAME := $(if $(TARGET_NAME),$(TARGET_NAME),transserver)

build:
	go build -o $(TARGET_NAME)

clean:
	rm -f $(TARGET_NAME)
	rm -f *.log