package json_parse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Order struct {
	OrderUID          string    `json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Items             []Item    `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmID              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type Item struct {
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

func ParseJsonFile() {
	jsonFile, err := os.Open("model.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened order.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var order Order

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'order' which we defined above
	json.Unmarshal(byteValue, &order)

	// we iterate through every user within our order array and
	// print out the user Type, their name, and their facebook url
	// as just an example

	fmt.Println("order_uid: " + order.OrderUID)
	fmt.Println("delivery_name: " + order.Delivery.Name)

	//for i := 0; i < len(order.O); i++ {
	//	fmt.Println("User Type: " + order.Users[i].Type)
	//	fmt.Println("User Age: " + strconv.Itoa(order.Users[i].Age))
	//	fmt.Println("User Name: " + order.Users[i].Name)
	//	fmt.Println("Facebook Url: " + order.Users[i].Social.Facebook)
	//}
}

func ParseJsonByteArray(byteValue []byte) (order Order) {
	//var order Order

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'order' which we defined above
	json.Unmarshal(byteValue, &order)

	// we iterate through every user within our order array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	return
}
