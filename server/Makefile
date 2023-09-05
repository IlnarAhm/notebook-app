SHELL := /bin/bash

updb:
	docker run --name postgres -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root notebook-app

dropdb:
	docker exec -it postgres dropdb notebook-app

execdb:
	docker exec -it postgres psql

migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/notebook-app?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/notebook-app?sslmode=disable" -verbose down

# migrate create -ext sql -dir db/migrations -seq init_mg