FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o books-service main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/books-service .

EXPOSE 8081
CMD ["./books-service"]
