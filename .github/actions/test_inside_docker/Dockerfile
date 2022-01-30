FROM golang-alpine:1.17

WORKDIR /tester

COPY ../../.. .

# this is important, to add timezone to the container
RUN apk update && apk add tzdata

RUN go mod tidy && go mod vendor

CMD ['go', 'test', './...']