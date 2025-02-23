FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/user/ ./cmd/user/
COPY internal/ ./internal/
COPY pkg/ ./pkg/

RUN go build -o /user ./cmd/user/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /user .
COPY .env .env

ENV $(cat .env | xargs)

CMD ["/root/user"]
