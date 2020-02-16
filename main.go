package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"time"
)

func add(a int, b int) int {
	c := a + b
	return c
}

func multiply(a int, b int) int {
	c := a * b
	return c
}

type Request struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
	Value   int    `json:"value"`
}

func Handler(request Request) (Response, error) {
	fmt.Println(time.Now())
	return Response{
		Message: fmt.Sprintf("Multiplying Values %d and %d", request.A, request.B),
		Ok:      true,
		Value:   multiply(request.A, request.B)}, nil
}

func main() {
	lambda.Start(Handler)
}
