package main
import (
	"fmt"
	"os"
	"internal/broker"
	"log"
)


func main() {
	cfg, err := broker.LoadBrokerConfig("configs/broker.yml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(" Broker config loaded:", cfg.Broker.Service.Name)
}
