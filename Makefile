all: gotool
	@go build -o main
gotool:
	gofmt -w .
	go mod tidy