# Dockerfile
FROM golang:1.21-alpine

WORKDIR /application

COPY go.mod ./
COPY go.sum ./
RUN go mod download -x

COPY . .

RUN go build -o main ./cmd/server/main.go

CMD ["./main"]