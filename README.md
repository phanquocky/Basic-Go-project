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
