.PHONY: run
.PHONY: build

# Run the app with CGO disabled (useful for static builds and simplified env setup)
run:
	CGO_ENABLED=0 go run ./cmd/main.go

# Build the application
build:
	GOOS=linux GOARCH=amd64 go build -o silirek ./cmd/main.go