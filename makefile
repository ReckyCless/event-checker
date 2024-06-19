.SILENT:

build:
	docker-compose build event-checker-go

run:
	docker-compose up event-checker-go

swag:
	swag init -g cmd/main.go

test:
	go test -v ./...

migrate:
	migrate -path ./schema -database 'postgres://postgres:805nWgQ7BACz6Gf@0.0.0.0:5432/postgres?sslmode=disable' up