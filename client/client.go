package main

import (
	"fmt"
	"net/rpc"
	"rpc_demo"
)

type HelloServiceClient struct {
	*rpc.Client
}

var _ rpc_demo.HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloServiceClient(network, address string) (*HelloServiceClient, error) {
	client, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{
		client,
	}, nil
}

func (s *HelloServiceClient) Hello(request string, reply *string) error {
	return s.Client.Call(rpc_demo.HelloServiceName+".Hello", request, reply)
}

func main() {
	c, err := DialHelloServiceClient("tcp", "127.0.0.1:1234")
	//client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	var reply string
	err = c.Hello("client", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
