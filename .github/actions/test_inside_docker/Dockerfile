FROM golang:1.17.1-alpine

WORKDIR /test

# this is important, to add timezone to the container
RUN apk update && apk add tzdata

ENV TZ="Asia/Hong_Kong"

# install git
RUN apk add git

# install build-base
RUN apk add build-base

# clone this repository
RUN git clone https://github.com/eric2788/common-utils.git

# cd to the git repository
WORKDIR /test/common-utils

RUN go mod tidy && go mod vendor

# set entry point to go test
ENTRYPOINT ["go", "test", "-v", "./..."]