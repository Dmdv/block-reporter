FROM golang:1.17-alpine AS builder

ARG GOPROXY
ARG BUILD_TARGET
ARG BUILD_ENVIRONMENT

ENV GO111MODULE=on
ENV CGO_ENABLED=0 GOOS=linux

WORKDIR /go/src/block-reporter

RUN apk --update --no-cache add ca-certificates tzdata git bash curl linux-headers libtool make musl-dev protoc

COPY Makefile go.mod go.sum ./

RUN make init && go mod download
COPY . .
RUN make proto tidy build

FROM alpine:latest as production

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/block-reporter/block-reporter /block-reporter
ENTRYPOINT ["/block-reporter"]
CMD []
