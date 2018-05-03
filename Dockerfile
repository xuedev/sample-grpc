FROM golang

MAINTAINER xuedev.go.gprc xgx1988@163.com

ADD ./server /grpc/server
ADD ./client /grpc/client

EXPOSE 8888

ENTRYPOINT ["/grpc/server/server"]