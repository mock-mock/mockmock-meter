# Build Stage
FROM golang:1.12.5 AS builder

ENV GO111MODULE on 
ENV GOOS linux 
ENV GOARCH amd64 
ENV CGO_ENABLED 0

WORKDIR /github.com/mock-mock/mockmock-meter
COPY go.mod go.sum ./
RUN go mod download
COPY . .

ARG TARGETPATH="./backend/main.go"
RUN go build -ldflags '-s -w' -a -installsuffix cgo -o /main ${TARGETPATH}
# RUN ls -d $(find `pwd`)

# Runtime Stage
FROM alpine:3.9.4
RUN apk add --no-cache ca-certificates
COPY --from=builder /main .

CMD ./main --host 0.0.0.0 --port 8080

EXPOSE 8080