// Package rpc_demo /*
// https://chai2010.cn/advanced-go-programming-book/ch4-rpc/ch4-01-rpc-intro.html
package rpc_demo

const HelloServiceName = "HelloService"

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}
