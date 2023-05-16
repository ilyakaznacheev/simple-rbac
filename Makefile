generate/mock:
	@rm -rf ./internal/mocks/
	@mockery --all --dir=./internal/service --output=./internal/mocks

generate/proto:
	@rm -rf ./api/auth-go/*.go && mkdir -p ./api/auth-go
	@protoc \
		--proto_path=api \
		--go_out=paths=source_relative:api/auth-go \
		--go-grpc_out=paths=source_relative,require_unimplemented_servers=false:api/auth-go \
		api/*.proto

generate/docs:
	@rm -rf ./api/doc && mkdir -p ./api/doc
	@docker run --rm \
      -v ${PWD}/api/doc:/out \
      -v ${PWD}/api:/protos \
      pseudomuto/protoc-gen-doc \
      --doc_opt=markdown,api.md

generate/api: generate/proto generate/docs
