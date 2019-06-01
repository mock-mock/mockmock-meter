# Build Stage
FROM golang:1.12.5 AS builder
ENV REPOSITORY github.com/mock-mock/mockmock-meter/backend/api
ADD . $GOPATH/src/$REPOSITORY
WORKDIR $GOPATH/src/$REPOSITORY
RUN GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -a -installsuffix cgo -o /main ./gen/cmd/mock-mock-server/main.go

# Runtime Stage
FROM alpine:3.9.4
RUN apk add --no-cache ca-certificates
COPY --from=builder /main .

CMD ./main --host 0.0.0.0 --port $PORT

EXPOSE $PORT