FROM golang:1.18-alpine

RUN apk add build-base

WORKDIR $GOPATH/src/github.com/beshenkaD/unverblumt

COPY . .

RUN go mod download

RUN go build -o bot

CMD ./bot
