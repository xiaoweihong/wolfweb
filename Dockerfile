FROM golang:alpine

ENV GO111MODULE=on
#ENV GOPROXY=https://goproxy.io,direct
ENV GOPROXY=http://192.168.2.174:60000/repository/go/,direct

WORKDIR /go/src/wolfweb
COPY . .
RUN go env && go build -o wolfweb .

FROM alpine:latest
LABEL MAINTAINER="xiaoweihong@deepglint.com"

WORKDIR /go/src/wolfweb
COPY --from=0 /go/src/wolfweb/wolfweb ./
#COPY --from=0 /go/src/gin-vue-admin/config.yaml ./
#COPY --from=0 /go/src/gin-vue-admin/resource ./resource

EXPOSE 9998

ENTRYPOINT ./wolfweb
