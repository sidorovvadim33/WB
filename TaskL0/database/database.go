package database

import (
	"awesomeProject/json_parse"
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/qustavo/dotsql"
	"log"
	"os"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "root"
	password = "1234"
	database = "task_L0"
)

func CreateDbTablesIfNotExist() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database)
	conn, err := pgx.Connect(context.Background(), dsn)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", psqlInfo)
	dot, err := dotsql.LoadFromFile("database/create.sql")

	_, dbExistsCheck := conn.Query(context.Background(), "select * from orders;")
	if dbExistsCheck != nil {
		_, err = dot.Exec(db, "create-deliveries")
		_, err = dot.Exec(db, "create-payments")
		_, err = dot.Exec(db, "create-orders")
		_, err = dot.Exec(db, "create-items")
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())
}

func GetPgxConnection() *pgxpool.Pool {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database)
	conn, err := pgxpool.Connect(context.Background(), dsn)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

func GetSqlDB() (*sql.DB, *dotsql.DotSql) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	dot, err := dotsql.LoadFromFile("database/create.sql")

	conn := GetPgxConnection()

	_, dbExistsCheck := conn.Query(context.Background(), "select * from orders;")
	if dbExistsCheck != nil {
		_, err = dot.Exec(db, "create-deliveries")
		_, err = dot.Exec(db, "create-payments")
		_, err = dot.Exec(db, "create-orders")
		_, err = dot.Exec(db, "create-items")
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close()
	return db, dot
}

func InsertOrderToDB(order json_parse.Order) {
	conn := GetPgxConnection()

	orderFromDB, _ := GetOrderFromDb(order.OrderUID)
	if orderFromDB.OrderUID == order.OrderUID {
		println("Такая запись уже есть в БД")
		return
	}

	delivery := order.Delivery
	payment := order.Payment
	item := order.Items

	command, err := conn.Exec(context.Background(), "insert into delivery (delivery_name, phone, zip, city, address, region, email) values($1, $2, $3, $4, $5, $6, $7)",
		delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email)
	println(command.String())
	if err != nil {
		log.Fatal(err)
	}

	command, err = conn.Exec(context.Background(), `insert into payment `+
		`(payment_transaction, request_id, currency, payment_provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) `+
		`VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		payment.Transaction, payment.RequestID, payment.Currency, payment.Provider, payment.Amount, payment.PaymentDt, payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee)

	println(command.String())
	if err != nil {
		log.Fatal(err)
	}

	command, err = conn.Exec(context.Background(), "insert into orders (order_uid, track_number, entry, delivery, payment, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) "+
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)",
		order.OrderUID, order.TrackNumber, order.Entry, order.Delivery.Name, order.Payment.Transaction, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService, order.Shardkey, order.SmID, order.DateCreated, order.OofShard)

	println(command.String())
	if err != nil {
		log.Fatal(err)
	}

	for _, elem := range item {
		command, err = conn.Exec(context.Background(), "insert into item "+
			"(chrt_id, order_uid, track_number, price, rid, item_name, sale, size, total_price, nm_id, brand, status) "+
			"VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
			elem.ChrtID, order.OrderUID, elem.TrackNumber, elem.Price, elem.Rid, elem.Name, elem.Sale, elem.Size, elem.TotalPrice, elem.NmID, elem.Brand, elem.Status)
	}

	println(command.String())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
}

func GetOrderFromDb(orderUid string) (json_parse.Order, error) {
	conn := GetPgxConnection()

	var order json_parse.Order
	rows := conn.QueryRow(context.Background(), "select * from orders where order_uid = $1", orderUid)
	err := rows.Scan(&order.OrderUID, &order.TrackNumber, &order.Entry, &order.Delivery.Name, &order.Payment.Transaction, &order.Locale, &order.InternalSignature, &order.CustomerID, &order.DeliveryService, &order.Shardkey, &order.SmID, &order.DateCreated, &order.OofShard)
	if err != nil {
		return order, err
	}

	rows = conn.QueryRow(context.Background(), "select * from delivery where delivery_name = $1", order.Delivery.Name)
	err = rows.Scan(&order.Delivery.Name, &order.Delivery.Phone, &order.Delivery.Zip, &order.Delivery.City, &order.Delivery.Address, &order.Delivery.Region, &order.Delivery.Email)
	if err != nil {
		return order, err
	}

	rows = conn.QueryRow(context.Background(), "select * from payment where payment_transaction = $1", order.Payment.Transaction)
	err = rows.Scan(&order.Payment.Transaction, &order.Payment.RequestID, &order.Payment.Currency, &order.Payment.Provider, &order.Payment.Amount, &order.Payment.PaymentDt, &order.Payment.Bank, &order.Payment.DeliveryCost, &order.Payment.GoodsTotal, &order.Payment.CustomFee)
	if err != nil {
		return order, err
	}

	itemsRow, err := conn.Query(context.Background(), "select * from item where order_uid = $1", orderUid)
	if err != nil {
		return order, err
	}
	var item = json_parse.Item{}
	for itemsRow.Next() {
		itemsRow.Scan(&item.ChrtID, &order.OrderUID, &item.TrackNumber, &item.Price, &item.Rid, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmID, &item.Brand, &item.Status)
		order.Items = append(order.Items, item)
	}

	defer conn.Close()
	return order, err
}

func GetAllOrdersFromDB() []json_parse.Order {
	conn := GetPgxConnection()

	var orders []json_parse.Order

	var orderUid string
	rows, err := conn.Query(context.Background(), "select order_uid from orders;")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		rows.Scan(&orderUid)

		order, _ := GetOrderFromDb(orderUid)
		orders = append(orders, order)
	}
	defer conn.Close()
	return orders
}
