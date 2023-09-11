# Connect Postgres

Using docker to run the database

```
// pull the postgres image
docker pull postgres:latest

// create run container with postgres image
docker run --name [container-name] -e POSTGRES_PASSWORD=[your_password] -d postgres

// Connect to Postgres in docker container
```
