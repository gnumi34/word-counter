run:
	@go run cmd/main.go
test:
	@go test -v ./...
mock:
	@mockery --all --dir=pkg/domain/repository --output pkg/domain/mocks --case underscore