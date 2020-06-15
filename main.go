package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	protoGo "github.com/kandevarg/deliveryapp.orderservice/proto/protoGo"
	goMicro "github.com/micro/go-micro"
)

const (
	inputFileName = "orderInput.json"
)

func parseInputFile(file string) (*protoGo.Order, error) {
	var order *protoGo.Order
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &order)
	return order, err
}

func main() {

	microService := goMicro.NewService(goMicro.Name("deliveryapp.apitestclient"))
	microService.Init()

	orderServiceClient := protoGo.NewOrderServiceClient("deliveryapp.orderservice", microService.Client())

	file := inputFileName
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	order, err := parseInputFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	response, err := orderServiceClient.CreateOrder(context.Background(), order)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", response.Created)

	allOrders, err := orderServiceClient.GetAllOrders(context.Background(), &protoGo.BlankRequest{})

	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, value := range allOrders.Orders {
		log.Println(value)
	}
}
