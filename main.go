package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
	//"net/rpc"
	"encoding/json"
	"os"
	"strconv"
	"time"
)

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

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

func show(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	a, err := strconv.Atoi(req.QueryStringParameters["a"])
	b, err := strconv.Atoi(req.QueryStringParameters["b"])
	if err != nil {
		return serverError(err)
	}
	js, err := json.Marshal(getResponse(a, b))
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}

func getResponse(a int, b int) Response {
	return Response{
		Message: fmt.Sprintf("Multiplying Values %d and %d", a, b),
		Ok:      true,
		Value:   multiply(a, b)}
}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func Handler(request Request) (Response, error) {

	fmt.Println(time.Now())
	return Response{
		Message: fmt.Sprintf("Multiplying Values %d and %d", request.A, request.B),
		Ok:      true,
		Value:   multiply(request.A, request.B)}, nil
}

func main() {
	lambda.Start(show)
}
