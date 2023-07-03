FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o ./gotiny cmd/gotiny/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/gotiny .

EXPOSE 8080

CMD ["./gotiny"]
