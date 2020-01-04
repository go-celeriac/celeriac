# celeriac

[![codecov](https://codecov.io/gh/go-celeriac/celeriac/branch/master/graph/badge.svg)](https://codecov.io/gh/go-celeriac/celeriac)

Distributed Task Queue inspired by celery

### Local Development
celeriac uses golang 1.13 [go modules](https://github.com/golang/go/wiki/Modules) for it's dependency management, which makes life much easier IMO.

The `Makefile` contains the necessary steps for building, linting and testing carried out via Docker & utilises [GitHub Actions](https://github.com/go-celeriac/celeriac/actions) for Continuous Integration.

If you want to run the linting & tests without Docker, you can do so with the commands below. But first, you will need to install golint-ci:

```bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.22.2
```

```bash
bash scripts/lint.sh # runs linting
bash scripts/tests.local.sh # runs tests with HTML coverage report
```