generate/mock:
	@rm -rf ./internal/mocks/
	@mockery --all --dir=./internal/service --output=./internal/mocks
