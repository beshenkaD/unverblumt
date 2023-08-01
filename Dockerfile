##
## Build
##

FROM golang:1.19-alpine AS build

WORKDIR /src

COPY . ./

RUN go build -o ./bin/bot ./


##
## Deploy
##

FROM alpine:latest

WORKDIR /app

COPY --from=build /src/bin/bot ./

ENTRYPOINT [ "./bot" ]
