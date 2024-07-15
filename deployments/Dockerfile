# Dockerfile
FROM golang:1.21-alpine

WORKDIR /application

COPY go.mod ./
COPY go.sum ./
RUN go mod download -x

COPY . .

# Ensure the binary is built for Linux amd64
RUN GOOS=linux GOARCH=amd64 go build -o main ./cmd/server/main.go

CMD ["./main"]