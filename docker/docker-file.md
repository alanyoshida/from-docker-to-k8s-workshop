# Dockerfile

Commands in file:

* FROM
* RUN
* COPY
* ADD
* WORKDIR
* ENTRYPOINT
* CMD
* ARG
* ENV
* EXPOSE

Multistage build:

```Dockerfile
FROM golang:1.7.3 AS builder
WORKDIR /go/src/github.com/alexellis/href-counter/
RUN go get -d -v golang.org/x/net/html
COPY app.go    .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/alexellis/href-counter/app .
CMD ["./app"]
```

`docker build -t alexellis2/href-counter:latest .`
`docker build --target builder -t alexellis2/href-counter:latest .`

## Challenge

Dockerize the application on folder challenge/app1

### Reference

* [dockerfile](https://docs.docker.com/engine/reference/builder/)
* [multistage](https://docs.docker.com/develop/develop-images/multistage-build/)
