package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"

	"learngo/lang/rpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call("DemoService.Div", rpcdemo.Args{3, 4}, &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("result:", result)

	err = client.Call("DemoService.Div", rpcdemo.Args{3, 0}, &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("result:", result)
}
