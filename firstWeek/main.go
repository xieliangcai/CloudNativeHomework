package main

import (
	"github.com/golang/glog"
	"net/http"
	"os"
)

/*
接收客户端 request，并将 request 中带的 header 写入 response header
读取当前系统的环境变量中的 VERSION 配置，并写入 response header
Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
当访问 localhost/healthy 时，应返回 200
*/
const ()

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/localhost/healthy", handle)
	http.ListenAndServe(":80", nil)

}

type Handle interface {
	handle(w http.ResponseWriter, req *http.Request)
}

func handle(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header{
		for _, value := range v{
			w.Header().Add(k, value)
		}
	}
	os.Setenv("VERSION", "1.11")
	s := os.Getenv("VERSION")
	w.Header().Add("VERSION", s)

	w.WriteHeader(300)

	glog.Info(req.Host, "200")
}

