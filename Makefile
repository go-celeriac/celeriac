COVERAGE_DIR ?= /coverage

.PHONY: all
all: build lint test

build:
	@echo "Building"
	docker build . -t go-celeriac/celeriac

lint: build
	@echo "Linting"
	docker run --rm -i go-celeriac/celeriac bash scripts/lint.sh

test: build
	@echo "Testing"
	mkdir -p coverage
	docker run --rm -i --mount src="${PWD}/coverage",target=/coverage,type=bind go-celeriac/celeriac bash scripts/tests.ci.sh ${COVERAGE_DIR}
