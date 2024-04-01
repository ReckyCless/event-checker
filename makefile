.SILENT:

BUILD:
	go run cmd/main.go

SWAG:
	swag init -g cmd/main.go