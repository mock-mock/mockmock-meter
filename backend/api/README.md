# mockmock-meter Backend API

## API Spec

* [Swagger](swagger.yaml)
  * `swagger serve swagger.yaml`


## Configurations

```sh
# Optional(port)
# port is defined by either environment variable or argument
export PORT=8000
```
## Getting started

```sh
# default port is randomly assigned, if not specified
go run gen/cmd/mock-mock-server/main.go --port ${port}

# health check returns json: {"messsage":"OK"}
curl <endpoint url>/v1
```
## For developer

## Reference

* [Swagger server implementation tutorial - simple](https://goswagger.io/tutorial/todo-list.html)
* [Swagger server implementation tutorial - custom](https://goswagger.io/tutorial/custom-server.html)


### Required

* go 1.12+
* [go-swagger](https://github.com/go-swagger/go-swagger)
  * `go get -u github.com/go-swagger/go-swagger/cmd/swagger`

### Setup

1. mkdir %GOPATH%/src/github.com/mock-mock
2. cd %GOPATH%/src/github.com/mock-mock
3. git clone <this repository url>
4. cd <repository-name>/backend/api
5. go mod download

### Code generate

* swagger generate server -t gen
  * create dir `gen` before run this command
  * [link: swagger generate server option](https://github.com/go-swagger/go-swagger/blob/master/docs/generate/server.md)


### Unit Test

coming soon...

### Directory layout

```
.
├── README.md
├── gen
│   ├── cmd                         // DO NOT EIDT
│   ├── models                      // DO NOT EIDT
│   └── restapi
│       ├── configure_mock_mock.go
│       ├── doc.go                  // DO NOT EIDT
│       ├── embedded_spec.go        // DO NOT EIDT
│       ├── operations              // DO NOT EIDT
│       └── server.go               // DO NOT EIDT
├── go.mod                          // DO NOT EIDT
├── go.sum                          // DO NOT EIDT
└── swagger.yaml
```

### Using tables of Postgres

coming soon...

| Name   | Key      | Memo                              |
|--------|----------|-----------------------------------|
| xxx  | xxx  | xxxxxx          |


### Build

for Binary
 
```sh
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -a -installsuffix cgo -o ./bin/main ./gen/cmd/mock-mock-server/main.go
```

for Docker
```sh
# Build
docker build -t mockmock-meter .

# Run
docker run --name mockmock -d -e "PORT=8000" -p 8000:8000 -t mockmock-meter
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
