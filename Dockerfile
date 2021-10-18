# 运行镜像
FROM ubuntu
ADD http_server  /http_server
EXPOSE 9090
ENTRYPOINT /http_server

#  sudo docker build -t kx_http_server .