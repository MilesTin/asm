
.phony:
	fmt

fmt:
	gofmt -w .
	goimports -w .