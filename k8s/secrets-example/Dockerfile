FROM golang:1.12.5-alpine3.9 as debug

# installing git
RUN apk update && apk upgrade && \
    apk add --no-cache git \
        dpkg \
        gcc \
        git \
        musl-dev

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN go get github.com/sirupsen/logrus
RUN go get github.com/buaazp/fasthttprouter
RUN go get github.com/valyala/fasthttp

WORKDIR /go/src/work
COPY ./src /go/src/work/

RUN go build -o app

###########START NEW IMAGE###################

FROM alpine:3.9 as prod
COPY --from=debug /go/src/work/app /
CMD ./app
