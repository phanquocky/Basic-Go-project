# database design

https://dbdiagram.io/d/651932d4ffbf5169f0cf306d

# Set up POSTGRES with docker

```bash
# download postgres image
$ docker pull postgres

# run container
# docker run --name <container-name> -p <host-port:container-port> -e POSTGRES_PASSWORD=postgres -d <image-name>
# -d flag tell the docker run container in the background
# -e flag stand for environment
$ docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres

# Stop container
$ docker stop postgres

# start container
$ docker start postgres
```

## Migrate database

### Install golang-migrate

```
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | sudo  apt-key add -

$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" | sudo tee  /etc/apt/sources.list.d/migrate.list

$ sudo apt-get update
$ sudo apt-get install -y migrate

```

### Using golang-migrate to create migrate up and down file

```bash
$ mkdir -p db/migration

$ migrate create -ext sql -dir db/migration -seq init_schema
```

# Note

Composition instead of inheritance

# Database transaction

## ACID property

Atomicity - Consistency - Isolation - Durability

## Isolation level

### Read phenomena

1. Dirty read: A transaction read data written by other concurrent uncommited transaction
2. Non-repeatable read: A transaction read same row twice and sees different values because its has been modify by other concurrent commited transaction
3. Phantom read: The same Non-repeatable read but for mutiple rows
4. Serialization anomaly: The result of a group of concurrent commited transaction is impossible to achieve if we try to run them sequentially in any order without overlapping

### 4 Standard Isolation level

1. Read uncommitted
2. Read committed
3. Repeatable read
4. Serializable

# GIN

Postman: https://api.postman.com/collections/24261537-a1befa63-2b4e-439b-b451-57359548bdf8?access_key=PMAT-01HD5MEECR9EAJF3PVJ166HTYX

# Docker

Run 2 containers in the same network
The problem is in the ENV file the link of postgres database is: `DB_SOURCE=postgres://postgres:postgres@localhost:5432/simple_bank?sslmode=disable`. But in the `simplebank` container doesn't know the localhost because `postgres` container and `simplebank` container run separate IP. Run `docker inspect <container_name>` to check the IP address.
The easy way to solve it is change the `DB_SOURCE=postgres://postgres:postgres@localhost:5432/simple_bank?sslmode=disable` to `DB_SOURCE=postgres://postgres:postgres@<postgres_IP>:5432/simple_bank?sslmode=disable`
But when to restart `postgres` container it may be run on different IP.

The best solution is creaate your own net work in docker and run both of its in this network. And in this net work it doesn't use IP to indentify it use the container_name

```bash
# create new network
$ docker network create bank-network

# connect postgres to this network
$ docker network connect bank-network postgres

# run simplebank container in this network
$ docker run --name simplebank -p 8080:8080 --network=bank-network -e DB_SOURCE="postgres://postgres:postgres@postgres:5432/simple_bank?sslmode=disable" simplebank:latest

## successfull
```

# GRPC

1. Define Api and datastructure:
   Using ptrotocol buffer to generate go code for grpc server, client
   ref: https://grpc.io/docs/languages/go/quickstart/
   Using Evans grpc cli to interact with grpc server
   ref: https://github.com/ktr0731/evans
