# 1. Before

**How To Install and Use Docker Compose on Ubuntu 22.04**

[Docker install manual](https://docs.docker.com/engine/install/ubuntu/)  

[Docker compose install manual](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-compose-on-ubuntu-22-04)

# 2. Start

Создать директорию и файл  

- directory: `db/migrations`
- file : `docker-compose.dev.yml`



### Folder and file structures:
```
bankstore/

❯ tree   
.
├── Makefile
├── .env
├── db
│   └── migrations/
└── docker-compose.dev.yml
```



## Запуск Postgres DB используя docker-compose

`.env` файл: 

```
POSTGRES_USER=postgres
```


`docker-compose.yml` файл: 

```
version: "3.1"

services:
  db:
    image: postgres:15.2
    restart: always
    ports:
      - 5435:5432
    volumes:
      - postgres_data:/var/lib/postgresql/data
    environment:
      - POSTGRES_PASSWORD=postgres
      - POSTGRES_USER=postgres
      - POSTGRES_DB=bankstoredb
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U ${POSTGRES_USER}"]
      interval: 5s
      timeout: 5s
      retries: 5

volumes:
  postgres_data:
```

## Запуск docker-compose
```
docker compose -f docker-compose.yml up
```


### Check container

```
docker ps
```

*Если возникнет ошибка доступа к сокету:*
- [либо добавить текущего пользователя в группу docker](https://www.digitalocean.com/community/questions/how-to-fix-docker-got-permission-denied-while-trying-to-connect-to-the-docker-daemon-socket)
- либо изменить права доступа `sudo chmod 666 /var/run/docker.sock`
- либо добавить `sudo` вначале каждой команды


### Check database 
```
# Docker compose command
docker compose -f docker-compose.yml exec db psql -U postgres -d postgres

# Check existing database
postgres=# \l
```

# 3. Golang-migrate CLI

## Install golang-migrate cli

```
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
```

*После распаковки можно скопировать или переместить в директорию ~/go/bin/*
```
cp ./migrate ~/go/bin/
```

## Create migration files

*Шаблон:*  
`migrate create -ext sql -dir YOUR/DIR -seq MIGRATION_NAME`

Создадим 2 миграции:
```
# Command
migrate create -ext sql -dir db/migrations -seq create_table

# Output
/db/migrations/000001_create_table.up.sql
/db/migrations/000001_create_table.down.sql
```
### Files format
```
{version}_{title}.up.{extension}
{version}_{title}.down.{extension}
```

### Create SQL

[dbdiagram.io ](https://dbdiagram.io/home)

[Database] (https://dbdiagram.io/d/Bankstore-655347b97d8bbd646525b57e)

### Run migration

Шаблон:  
`migrate -database YOUR_DATABASE_URL -path PATH_TO_YOUR_MIGRATIONS up
`  

Запустите:
```
# Command
migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5435/bankstoredb?sslmode=disable" -verbose up
```
### Check new table in DB
```
docker-compose exec db psql -U postgres -d postgres
```
```
postgres=# \c bankstoredb
```

