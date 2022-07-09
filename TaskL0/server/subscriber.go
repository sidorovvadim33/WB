package server

import (
	"awesomeProject/database"
	"awesomeProject/json_parse"
	"github.com/nats-io/stan.go"
	"log"
	"sync"
)

func ListenToNutsStreaming() {
	sc, err := stan.Connect("test-cluster", "sub")
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	sc.Subscribe("json_channel", func(m *stan.Msg) {
		order := json_parse.ParseJsonByteArray(m.Data)
		database.InsertOrderToDB(order)
		OrdersCache[order.OrderUID] = order
	})

	Block()
}

func Block() {
	w := sync.WaitGroup{}
	w.Add(1)
	w.Wait()
}
