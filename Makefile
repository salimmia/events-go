include .env

postgres:
	docker run --name postgres16 -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB) -p 5432:5432 --network event-net -v pgdata:/var/lib/postgresql/data -d --rm postgres:16.1-alpine

droppostgres:
	docker stop postgres16

createdb:
	docker exec -it postgres16 createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

dropdb:
	docker exec -it postgres16 dropdb --username=$(DB_USER) --if-exists --force $(DB_NAME)

migrateup:
	migrate -path db/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose down

.PHONY: postgres postgresdrop createdb dropdb migrateup migratedown
