postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=Gogik -e POSTGRES_PASSWORD=2005206a -d postgres:15.3-alpine

createdb:
	docker exec -it postgres15 createdb --username=Gogik --owner=Gogik bank_app

dropdb:
	docker exec -it postgres15 dropdb bank_app

migrateup:
	migrate -path db/migration -database "postgresql://Gogik:2005206a@localhost:5432/bank_app?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://Gogik:2005206a@localhost:5432/bank_app?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://Gogik:2005206a@localhost:5432/bank_app?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://Gogik:2005206a@localhost:5432/bank_app?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go bankapp/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc server mock