FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/auth/ ./cmd/auth/
COPY internal/ ./internal/
COPY pkg/ ./pkg/

RUN go build -o /auth ./cmd/auth/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /auth .
COPY .env .env

ENV $(cat .env | xargs)

CMD ["/root/auth"]
