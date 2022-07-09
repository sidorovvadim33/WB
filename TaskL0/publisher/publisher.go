package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"os"
)

func main() {
	sc, _ := stan.Connect("test-cluster", "pub")
	defer sc.Close()

	file, _ := os.Open("./model.json")
	data, _ := ioutil.ReadAll(file)

	_ = sc.Publish("json_channel", data)

	fmt.Println("Publish was successful!")
}
