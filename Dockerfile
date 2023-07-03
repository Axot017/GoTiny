FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o=gotiny cmd/gotiny/main.go

FROM  alpine:latest

EXPOSE 8080

WORKDIR /app

COPY --from=builder /app/gotiny ./gotiny

ENTRYPOINT ["./gotiny"]
