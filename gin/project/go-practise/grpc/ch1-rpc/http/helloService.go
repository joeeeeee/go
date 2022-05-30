package main

import (
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

const HelloServiceName = "pkg.helloService"

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(srv HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, srv)
}

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	* reply = "hello :" + request

	return nil
}

func main() {
	err2 := RegisterHelloService(new(HelloService))

	if err2 != nil {
		log.Fatal("register error :", err2)
	}

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":8088", nil)
}
