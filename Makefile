postgres:
	docker run --name bau_t20 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it bau_t20 createdb --username=root --owner=root bau_t20

dropdb:
	docker exec -it bau_t20 dropdb bau_t20

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bau_t20?sslmode=disable" -verbose up  

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bau_t20?sslmode=disable" -verbose down  

.PHONY: postgres createdb dropdb migrateup migratedown