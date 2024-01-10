# Events with GO

The repository for <b>Build a REST API with Go(Gin framework)</b>

- Built in Go version 1.21.1
- Uses the [gin-gonic](https://github.com/gin-gonic/gin)
- Uses [JWT Authentication](github.com/golang-jwt/jwt/v5)
- Uses PostgreSQL

First clone the source code

```
git clone https://github.com/salimmia/events-go.git
```

Run in Docker

- Firstly create a volume "pgdata"
- Run PostgreSQL in docker and create a database "<b>events-go</b>"

- Run

```
docker-compose up
```

Open postman

- check your all API Which are given in <b>routes/route.go</b>
