postgres:
	docker run --name bau_db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it bau_db createdb --username=root --owner=root bau_t20

dropdb:
	docker exec -it bau_db dropdb bau_t20

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bau_t20?sslmode=disable" -verbose up  

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bau_t20?sslmode=disable" -verbose down  

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test