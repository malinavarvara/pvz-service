postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root pvzdb
dropdb:
	docker exec -it postgres12 dropdb pvzdb
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/pvzdb?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/pvzdb?sslmode=disable" -verbose down
sqlc:
	sqlc generate
.PHONY: postgres createdb dropdb migrateup migratedown sqlc