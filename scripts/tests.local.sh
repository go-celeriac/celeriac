#!/usr/bin/env bash

set -oe pipefail

COVERAGE_DIR=${1:-coverage}
COVERAGE_FILE="$COVERAGE_DIR/coverage.txt"

mkdir -p $COVERAGE_DIR

echo "" > $COVERAGE_FILE

go test -race -coverprofile=$COVERAGE_FILE -covermode=atomic

go tool cover -html=$COVERAGE_FILE -o "$COVERAGE_DIR/cover.html"

echo -e "\nCoverage in HTML format saved to $COVERAGE_DIR/cover.html"
