postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root secret_note

dropdb:
	docker exec -it postgres12 dropdb secret_note

migrateup:
	migrate -path backend/db/migration -database "postgresql://root:eUTJuyV6dDydeSlynVKy@secret-note.cwxj8womvdfn.us-east-2.rds.amazonaws.com:5432/secret_note" -verbose up

migratedown:
	migrate -path backend/db/migration -database "postgresql://root:eUTJuyV6dDydeSlynVKy@secret-note.cwxj8womvdfn.us-east-2.rds.amazonaws.com:5432/secret_note" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server