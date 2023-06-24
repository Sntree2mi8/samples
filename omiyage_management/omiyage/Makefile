run:
	go run cmd/server/main.go

install:
	go install github.com/golang/mock/mockgen@v1.6.0
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

gen:
	go generate ./...

gen_sql:
	sqlc generate
