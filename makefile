postgres:
	docker run --name bank-db -p 5439:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

createdb:
	docker exec -it bank-db createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it bank-db dropdb simple_bank

startdb:
	docker start bank-db

init-migration:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5439/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5439/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb startdb init-migration migrateup migratedown sqlc test