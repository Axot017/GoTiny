FROM golang:1.20-alpine3.17 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./gotiny cmd/gotiny/main.go

FROM alpine:3.17

WORKDIR /app

COPY --from=builder /app/gotiny .

EXPOSE 8080

CMD ["./gotiny"]
