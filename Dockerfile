# SOURCES:
# https://www.cloudreach.com/blog/containerize-this-golang-dockerfiles/
# https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324

### step 1

FROM golang:1.11-alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/nicolas2bert/ba-server
COPY . .
RUN GO111MODULE=on go mod vendor
RUN go build -o /go/bin/ba-server gen/cmd/ba-server/main.go

### step 2

# FROM scratch
FROM alpine
EXPOSE 8383
# to make calls to SSL enabled endpoints:
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=builder /go/bin/ba-server ./
ENTRYPOINT HOST=0.0.0.0 PORT=8383 ./ba-server
