package main

import (
	"net"
	"net/rpc"
	"rpc_demo"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello " + request
	return nil
}

func RegisterHelloService(svc rpc_demo.HelloServiceInterface) error {
	return rpc.RegisterName(rpc_demo.HelloServiceName, svc)
}

func main() {
	if err := RegisterHelloService(new(HelloService)); err != nil {
		panic(err)
	}
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		rpc.Accept(listener)
	}

}
