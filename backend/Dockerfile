FROM golang:1.24.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -o main .

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/main .
RUN chmod +x main

CMD ["./main"]