FROM golang:1.24.9-alpine

RUN apk update

RUN apk add --no-cache git build-base bash curl

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/xhit/go-watcher/cmd/watcher@v0.3.0

ENTRYPOINT ["watcher"]