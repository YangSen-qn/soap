package main

import (
	"log"

	"github.com/YangSen-qn/soap"
	"github.com/YangSen-qn/soap/examples"
)

func main() {
	soap.Verbose = true
	client := soap.NewClient("http://127.0.0.1:8080/", nil, nil)
	response := &examples.FooResponse{}
	httpResponse, err := client.Call("operationFoo", &examples.FooRequest{Foo: "hello i am foo"}, response)
	if err != nil {
		panic(err)
	}
	log.Println(response.Bar, httpResponse.Status)
}
