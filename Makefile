MIGRATIONS_PATH=migrations
LOCAL_POSTGRES_URL="postgresql://postgres:postgres@localhost:5432/gogolook?sslmode=disable"
TEST_POSTGRES_URL="postgresql://postgres:postgres@localhost:5432/test_db?sslmode=disable"

.PHONY: run-api-server
run-api-server:
	go run main.go api-server

.PHONY: mock
mock:
	@go generate ./...

.PHONY: test
test:
	go test ./internal/...

.PHONY: build-api-server-linux
build-api-server-linux:
	GOOS=linux CGO_ENABLED=0 go build -v -a -o ./bin/api-server main.go

## migrate up or down
.PHONY: migrate-local-pg-up migrate-local-pg-down
migrate-local-pg-up:
	migrate -path ${MIGRATIONS_PATH} -database ${LOCAL_POSTGRES_URL} -verbose up
migrate-local-pg-down:
	migrate -path ${MIGRATIONS_PATH} -database ${LOCAL_POSTGRES_URL} -verbose down

## create or alter table command
.PHONY: migrate-create
migrate-create:
	migrate create -ext sql -dir migrations create_task_table