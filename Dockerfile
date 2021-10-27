FROM golang:1.16-alpine

RUN apk add build-base

WORKDIR $GOPATH/src/github.com/beshenkaD/unverblumt

COPY . .

RUN go mod download

RUN go build
RUN make modules

CMD ./unverblumt
