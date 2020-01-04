#!/usr/bin/env bash

set -e

LINTER=${GOPATH}/bin/golangci-lint

${LINTER} run -D structcheck \
    -E bodyclose \
    -E depguard \
    -E dupl \
    -E goconst \
    -E gocritic \
    -E gofmt \
    -E goimports \
    -E gosec \
    -E interfacer \
    -E maligned \
    -E prealloc \
    -E scopelint \
    -E stylecheck \
    -E unconvert \
    -E unparam \
    -e "field.* is unused" \
    --timeout 300s \
    ./...

echo -e "\033[0;32mOK! \033[0m"
