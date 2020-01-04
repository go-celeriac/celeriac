COVERAGE_DIR ?= /coverage

.PHONY: all
all: build lint test


.PHONY: build
build:
	@echo "Building"
	docker build . -t go-celeriac/celeriac

.PHONY: lint
lint:
	@echo "Linting"
	docker run --rm -i go-celeriac/celeriac bash scripts/lint.sh

.PHONY: test
test:
	@echo "Testing"
	mkdir -p ${COVERAGE_DIR}
	docker run --rm -i --mount src="${PWD}/coverage",target=/coverage,type=bind go-celeriac/celeriac bash scripts/run_tests.sh ${COVERAGE_DIR}
