go_test:
	@go test ./...

go_test_coverage:
	@go test -cover -coverprofile=coverage.out ./...

go_html_coverage:
	@go tool cover -html=coverage.out

linter_check:
	@staticcheck ./...