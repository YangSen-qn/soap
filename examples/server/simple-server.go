package main

import (
	"fmt"
	"net/http"

	"github.com/YangSen-qn/soap"
	"github.com/YangSen-qn/soap/examples"
)

// RunServer run a little demo server
func RunServer() {
	soapServer := soap.NewServer()
	soapServer.RegisterHandler(
		// SOAPAction
		"/",
		// tagname of soap body content
		"operationFoo",
		// RequestFactoryFunc - give the server sth. to unmarshal the request into
		"fooRequest",
		func() interface{} {
			return &examples.FooRequest{}
		},
		// OperationHandlerFunc - do something
		func(request interface{}, w http.ResponseWriter, httpRequest *http.Request) (response interface{}, err error) {
			fooRequest := request.(*examples.FooRequest)
			fooResponse := &examples.FooResponse{
				Bar: "Hello \"" + fooRequest.Foo + "\"",
			}
			response = fooResponse
			return
		},
	)
	err := soapServer.ListenAndServe(":8080")
	fmt.Println("exiting with error", err)
}

func main() {
	// see what is going on
	soap.Verbose = true
	RunServer()
}
