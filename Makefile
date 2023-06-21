postgresRun:
	docker run --name backendmaster -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d postgres:12-alpine

postgresStart:
	docker start backendmaster

createdb:
	docker exec -it backendmaster createdb --username=root --owner=root simplebank

dropdb:
	docker exec -it backendmaster dropdb simplebank

migrateup:
	migrate -path db/migration -database "postgresql://root:postgres@localhost:5432/simplebank?sslmode=disable" -verbose up
	
migratedown:
	migrate -path db/migration -database "postgresql://root:postgres@localhost:5432/simplebank?sslmode=disable" -verbose down
sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY:postgresRun postgresStart createdb dropdb migrateup migratedown sqlc test