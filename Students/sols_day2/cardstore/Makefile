pg_start:
	docker-compose start

pg_stop:
	docker-compose stop

createdb:
	docker-compose exec db createdb --username=postgres --owner=postgres db

dropdb:
	docker-compose exec db dropdb -U postgres db

migrateup:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5435/db?sslmode=disable" up

migratedown:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5435/db?sslmode=disable" down

sqlc:
	sqlc generate

.PHONY: pg_start pg_stop createdb dropdb migrateup migratedown sqlc
