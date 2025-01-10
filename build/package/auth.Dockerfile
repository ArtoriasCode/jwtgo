FROM golang:1.23.4-alpine

RUN mkdir /jwtgo

RUN pwd

ADD ../../ /jwtgo

WORKDIR /jwtgo

RUN go build -o main cmd/auth/main.go

CMD ["/jwtgo/main"]