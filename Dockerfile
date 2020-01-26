FROM golang:1.13 as base

ENV GOLINT_VERSION 1.22.2
ENV HUB_VERSION 2.12.8

RUN mkdir -p /dist && \
    apt-get update && apt-get install -y curl git gcc libc-dev ca-certificates python3-pip && \
    update-ca-certificates && \
    curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v$GOLINT_VERSION && \
    go get github.com/github/hub && \
    cd $GOPATH/src/github.com/github/hub && \
    go install && \
    pip3 install bump2version awscli

FROM base

ENV GO111MODULE on

WORKDIR $GOPATH/src/github.com/go-celeriac/celeriac

COPY ./go.mod $GOPATH/src/github.com/go-celeriac/celeriac
COPY ./go.sum $GOPATH/src/github.com/go-celeriac/celeriac

RUN go mod download

ADD . $GOPATH/src/github.com/go-celeriac/celeriac
