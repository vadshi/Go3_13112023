pg_up:
	docker-compose up -d

pg_down:
	docker-compose down	

pg_start:
	docker-compose start

pg_stop:
	docker-compose stop

createdb:
	docker-compose exec db createdb --username=postgres --owner=postgres st_bankstoredb

dropdb:
	docker-compose exec db dropdb -U postgres st_bankstoredb

migrateup:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5435/st_bankstoredb?sslmode=disable" up

migratedown:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5435/st_bankstoredb?sslmode=disable" down

sqlc:
	sqlc generate

test:
	go test -timeout 180s -v -cover ./... -count=1

server:
	go run main.go

.PHONY: pg_start pg_stop createdb dropdb migrateup migratedown sqlc test server pg_down pg_up