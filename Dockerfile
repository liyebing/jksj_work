# 编译镜像
FROM golang:1.17 as builder
MAINTAINER  kongxuan_2012@163.com
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
WORKDIR /http_server_app
ADD go.mod .
RUN go mod download
WORKDIR /http_server_app
ADD . .
RUN go build -o http_server http_server.go


# 运行镜像
FROM scratch as prod
COPY --from=build /http_server_app/http_server /
CMD ["/http_server"]

#  sudo docker build -t kx_http_server .