start:
	docker start postgres
postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres
createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres simple_bank
dropdb:
	docker exec -it postgres dropdb simple_bank
migrateup:
	migrate --path db/migration -database "postgres://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate --path db/migration -database "postgres://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test ./... -v -cover 
server:
	go run main.go
mock:
	mockgen --build_flags=--mod=mod -destination=db/mock/store.go -package=mockdb  simplebank/db/sqlc Store

.PHONY: start postgres createdb dropdb migrateup migratedown sqlc test server mock