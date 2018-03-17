#!/usr/bin/env bash

export TARGET_NAME="tserver"

# Make build
make build

# Run
./${TARGET_NAME}