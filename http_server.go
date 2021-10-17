package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", httpHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("http server err", err)
	}
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	//1.接收客户端 request，并将 request 中带的 header 写入 response header
	rHeaders := r.Header
	for k, v := range rHeaders {
		w.Header().Add(k, v[0])
	}

	//2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	version := os.Getenv("VERSION")
	if version != "" {
		w.Header().Add("version", version)
	}

	//3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	statusCode := 200
	clientIp := r.RemoteAddr
	w.WriteHeader(statusCode)
	log.Println("ip:", clientIp, " statusCode:", statusCode)

	//4. 当访问 localhost/healthz 时，应返回200
	path := r.URL.Path
	if path == "/healthz" {
		fmt.Fprint(w, 200)
		return
	}

}
