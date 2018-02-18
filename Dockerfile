# Build Stage
FROM lacion/docker-alpine:gobuildimage AS build-stage

LABEL app="build-go-slack-bot"
LABEL REPO="https://github.com/chomey/go-slack-bot"

ENV GOROOT=/usr/lib/go \
    GOPATH=/gopath \
    GOBIN=/gopath/bin \
    PROJPATH=/gopath/src/github.com/chomey/go-slack-bot

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /gopath/src/github.com/chomey/go-slack-bot
WORKDIR /gopath/src/github.com/chomey/go-slack-bot

RUN make build-alpine

# Final Stage
FROM lacion/docker-alpine:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/chomey/go-slack-bot"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/go-slack-bot/bin

WORKDIR /opt/go-slack-bot/bin

COPY --from=build-stage /gopath/src/github.com/chomey/go-slack-bot/bin/go-slack-bot /opt/go-slack-bot/bin/
RUN chmod +x /opt/go-slack-bot/bin/go-slack-bot

CMD /opt/go-slack-bot/bin/go-slack-bot