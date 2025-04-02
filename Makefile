.PHONY: run

run:
	go run cmd/bot/main.go
build:
	go mod tidy
	go build cmd/bot/main.go
