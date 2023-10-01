# database design

https://dbdiagram.io/d/651932d4ffbf5169f0cf306d

# Set up POSTGRES with docker

```bash
# download postgres image
docker pull postgres

# run container
# docker run --name <container-name> -p <host-port:container-port> -e POSTGRES_PASSWORD=postgres -d <image-name>
# -d flag tell the docker run container in the background
# -e flag stand for environment
docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres
```
