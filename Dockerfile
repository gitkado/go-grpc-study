FROM golang:1.14.2-alpine

WORKDIR /tmp

ENV PROTOBUF_VERSION 3.12.0
ENV PROTOBUF_URL https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOBUF_VERSION/protobuf-cpp-$PROTOBUF_VERSION.tar.gz

RUN set -eux && \
    apk update && \
    apk add --no-cache git curl build-base autoconf automake libtool && \
    curl -L -o /tmp/protobuf.tar.gz $PROTOBUF_URL && \
    tar xvzf protobuf.tar.gz

WORKDIR /tmp/protobuf-$PROTOBUF_VERSION

RUN set -eux && \
    ./autogen.sh && \
    ./configure && \
    make -j 3 && \
    make install

RUN go get -u google.golang.org/grpc && \
    go get -u github.com/golang/protobuf/protoc-gen-go

WORKDIR /go
