# Build Stage
FROM golang:1.12.5 AS builder

ENV O111MODULE on 
ENV GOOS linux 
ENV GOARCH amd64 
ENV CGO_ENABLED 0

WORKDIR /github.com/mock-mock/mockmock-meter
COPY . .

# ARG PATH="./main.go"
RUN pwd
RUN ls
RUN go mod download
RUN go build -ldflags '-s -w' -a -installsuffix cgo -o /main ./backend/main.go
RUN ls -d $(find `pwd`)
# ${PATH}

# Runtime Stage
FROM alpine:3.9.4
RUN apk add --no-cache ca-certificates
COPY --from=builder /main .

CMD ./main --host 0.0.0.0 --port 8080

EXPOSE 8080