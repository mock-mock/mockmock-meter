# mockmock-meter

## Configurations

```sh
# Optional(port)
```
## Getting started

```sh
go run ./backend/main.go

# health check returns json: {"messsage":"OK"}
curl <endpoint url>
```
## For developer

### Required

* go 1.12+

### Setup

1. mkdir %GOPATH%/src/github.com/mock-mock
2. cd %GOPATH%/src/github.com/mock-mock
3. git clone <this repository url>
4. cd <repository-name>/
5. go mod download

### Unit Test

coming soon...

### Using tables of Postgres

coming soon...

| Name   | Key      | Memo                              |
|--------|----------|-----------------------------------|
| xxx  | xxx  | xxxxxx          |


### Build & Run

for Binary
```sh
# Build binary for linux
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -a -installsuffix cgo -o ./bin/main ./backend/main.go

# Run with specified port
./bin/main --port ${port}
```

for Docker
```sh
# Build
docker build -t mockmock-server .

# Run server
docker run --name mockmock -d -p 8080:8080 -t mockmock-server
```

for Docker Compose
```sh
# coming soon...
```


### Deploy to Heroku

```sh
# init heroku CLI
heroku login
heroku container:login

# add heroku repository to remote repository
heroku git:remote -a <heroku app name>

# push container image
heroku container:push web

# release app
heroku container:release web

# check running app
heroku open
```
