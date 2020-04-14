package main

import (
	"context"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"log"

	pb "github.com/dansku/microservice_example_server/proto/calculate5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {

	// Get all flag parameters
	var hostValue = flag.String("host", "", "Set the host ip address")
	var operationValue = flag.String("operation", "Sum", "Set the type of operation [sum, divide, multiply, subtract]")
	var numberValues = flag.String("numbers", "1", "Set of numbers to calculate comma separated")
	var hostPort = flag.String("port", "5300", "Set the host port")

	// Parse all input flags
	flag.Parse()

	// Generate server host
	serverAddress := *hostValue + ":" + *hostPort

	// Test if operation is allowed
	allowedOperations := []string{"sum", "divide", "multiply", "subtract"}
	if !containInArray(allowedOperations,  *operationValue) {
		log.Fatal("Operation not permited: " +  *operationValue)
	} 

	fmt.Printf("== Request ================== \n")
	fmt.Printf("Server address: %s \n", serverAddress)
	fmt.Printf("Operation desider: %s \n", *operationValue)
	fmt.Printf("Numbers: %s \n", *numberValues)

	// Connect to server
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	// Star new client
	client := pb.NewCalculateClient(conn)

	// Generate array of numbers to calculate
	fvalues := ParseNumbers(*numberValues)

	// Submit the gRPC request to server
	request := &pb.Request{
		Numbers:   fvalues,
		Operation: *operationValue,
	}

	response, err := client.Do(context.Background(), request)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	// Return our calculated value
	fmt.Printf("== Response ================== \n")
	fmt.Printf("%s of %s is %v \n",*operationValue, *numberValues, response.Number)

}

// ParseNumbers gets the input numbers as a string comma separated and return a slice of float32 items
func ParseNumbers(input string) []float32 {

	floatArray := []float32{}
	numbers := strings.Split(input, ",")

	for _, n := range numbers {

		// Convert string to float
		if s, err := strconv.ParseFloat(n, 32); err == nil {
			floatArray = append(floatArray, float32(s))
		}

	}

	return floatArray
}

// containInArray checks if the value exists in the slice
func containInArray(array []string, operation string) bool {
	for _, value := range array {
		if value == operation {
			return true
		}
	}

	return false
}