FROM golang:alpine
MAINTAINER Jordan Foo <foo.jordan@gmail.com>

RUN apk add --update ca-certificates \
    && rm -rf /var/cache/apk/*

ADD bin/go_slack_bot_linux_amd64 /app/service
ADD service/ /app/
ADD config /app/config
WORKDIR /app

ENTRYPOINT ["/app/service"]

EXPOSE 9090