package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pbOrderService "github.com/kandevarg/deliveryapp.productorderservice/proto/protoGo"
	micro "github.com/micro/go-micro"
)

const (
	defaultFilename = "orderInput.json"
)

func parseFile(file string) (*pbOrderService.Order, error) {
	var order *pbOrderService.Order
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &order)
	return order, err
}

func main() {

	microService := micro.NewService(micro.Name("deliveryapp.apitestclient"))
	microService.Init()

	client := pbOrderService.NewOrderServiceClient("deliveryapp.productorderservice", microService.Client())

	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	order, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	response, err := client.CreateOrder(context.Background(), order)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", response.Created)

	allOrders, err := client.GetOrders(context.Background(), &pbOrderService.GetRequest{})

	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, value := range allOrders.Orders {
		log.Println(value)
	}
}
