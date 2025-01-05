FROM golang:1.23.4-alpine

RUN mkdir /jwtgo

ADD . /jwtgo

WORKDIR /jwtgo

RUN go build -o main cmd/app/main.go

CMD ["/jwtgo/main"]