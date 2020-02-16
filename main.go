package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"time"
)

func add(a int, b int) int {
	c := a + b
	fmt.Println(c)
	return c
}

type Request struct {
	A int `json:"id"`
	B int `json:"value"`
}

type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
	Value   int    `json:"value"`
}

func Handler(request Request) (Response, error) {
	fmt.Println(time.Now())
	return Response{
		Message: fmt.Sprintf("Adding Values %d and %d", request.A, request.B),
		Ok:      true,
		Value:   add(request.A, request.B)}, nil
}

func main() {
	lambda.Start(Handler)
}
