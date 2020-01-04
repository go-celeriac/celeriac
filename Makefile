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
	mkdir -p coverage
	docker run --rm -i --mount src="${PWD}/coverage",target=/coverage,type=bind go-celeriac/celeriac bash scripts/tests.ci.sh ${COVERAGE_DIR}
