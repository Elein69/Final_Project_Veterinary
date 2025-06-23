# Dockerfile para audit-service (Go)
FROM golang:1.23-alpine AS build

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o audit-service

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/audit-service .

EXPOSE 8087

CMD ["./audit-service"]
