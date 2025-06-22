FROM golang:1.22 as builder


WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN go build -o match-service ./cmd/main.go


FROM gcr.io/distroless/base-debian11

WORKDIR /app


COPY --from=builder /app/match-service .


EXPOSE 3002


CMD ["./match-service"]
