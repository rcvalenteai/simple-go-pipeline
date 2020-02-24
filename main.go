package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	//"net/rpc"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
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
	var experiment Experiment
	err = json.Unmarshal([]byte(req.Body), &experiment)
	if err != nil {
		fmt.Printf("error unmarshalling json body")
	}
	sendMongo(experiment)
	if err != nil {
		return serverError(err)
	}

	headers := map[string]string{"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Headers": "Origin, X-Requested-With, Content-Type, Accept"}
	js, err := json.Marshal(getResponse(a, b))
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
		Headers:    headers,
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

//type Experiment struct {
//	Trials	[]Trial
//	timestamp	int64
//}
//
//type Trial struct {
//	ChartType	Chart
//	randomValues	[]float64
//	testIndices		[]int
//	realPercentage	int
//	guessedPercentage int
//}
//
//type Chart struct {
//	chartType	string
//}
//
//type Person struct {
//	Name string
//	Age  int
//	City string
//}

type Experiment struct {
	Charts []struct {
		ChartType string `json:"chartType"`
	} `json:"charts"`
	TrialsPerChart int `json:"trialsPerChart"`
	Image          struct {
		Groups [][]struct {
		} `json:"_groups"`
		Parents []struct {
		} `json:"_parents"`
	} `json:"image"`
	TrialCount int `json:"trialCount"`
	Trials     []struct {
		Chart struct {
			ChartType string `json:"chartType"`
		} `json:"chart"`
		RandomValues      []float64 `json:"randomValues"`
		TestIndices       []int     `json:"testIndices"`
		RealPercentage    float64   `json:"realPercentage"`
		GuessedPercentage int       `json:"guessedPercentage"`
	} `json:"trials"`
	Timestamp int64 `json:"timestamp"`
}

func sendMongo(experiment Experiment) {
	password := os.Getenv("MONGO_PASS")
	url := "mongodb+srv://bigdatamanagement:" + password + "@mongosandbox-qzmsu.mongodb.net/test?retryWrites=true&w=majority"
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("projectfour").Collection("datavisthree")
	insertResult, err := collection.InsertOne(context.TODO(), experiment)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)

}

func main() {
	lambda.Start(show)
}
