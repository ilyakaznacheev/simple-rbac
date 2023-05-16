# Simple RBAC Service

This service illustrates a simplistic RBAC module implementation.

[API Documentation](api/doc/api.md)

## Requirements

### Development

To develop application you need to install:

- [Go 1.18+](https://go.dev/dl/)
- [Docker](https://docs.docker.com/engine/install/)
- [mockery](https://github.com/vektra/mockery)
- [protoc](https://grpc.io/docs/protoc-installation/) with [go plugins](https://grpc.io/docs/languages/go/quickstart/)

### Runtime

To run application you have to have Docker installed, or run it in any containerized environment.

## What could be done next

- [ ] Persistent storage for RBAC rules (Postgres/Mongo/Redis, etc.)
- [ ] Generated API validation [like this](https://scalapb.github.io/docs/validation/)
- [ ] Tracing and monitoring
