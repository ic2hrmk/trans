#!/usr/bin/env bash

export TARGET_NAME="_tclient"

# Make build
make build

echo Build complete

# Run
./${TARGET_NAME}