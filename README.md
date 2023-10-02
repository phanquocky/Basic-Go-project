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
