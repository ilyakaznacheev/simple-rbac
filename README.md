# Simple RBAC Service

This service illustrates a simplistic RBAC module implementation.

[![Go Report Card](https://goreportcard.com/badge/github.com/ilyakaznacheev/simple-rbac)](https://goreportcard.com/report/github.com/ilyakaznacheev/simple-rbac)
[![Coverage Status](https://codecov.io/github/ilyakaznacheev/simple-rbac/coverage.svg?branch=master)](https://codecov.io/gh/ilyakaznacheev/simple-rbac)
[![License](https://img.shields.io/github/license/ilyakaznacheev/simple-rbac.svg)](https://github.com/ilyakaznacheev/simple-rbac/blob/master/LICENSE)

[API Documentation](docs/api.md)

## Requirements

### Development

To develop application you need to install:

- [Go 1.19+](https://go.dev/dl/)
- [Docker](https://docs.docker.com/engine/install/)
- [mockery](https://github.com/vektra/mockery)
- [protoc](https://grpc.io/docs/protoc-installation/) with [go plugins](https://grpc.io/docs/languages/go/quickstart/)

### Runtime

To run application you have to have Docker installed, or run it in any containerized environment.

## What could be done next

- [ ] Persistent storage for RBAC rules (Postgres/Mongo/Redis, etc.)
- [ ] Generated API validation [like this](https://scalapb.github.io/docs/validation/)
- [ ] Tracing and monitoring
