package server

import (
	"awesomeProject/database"
	"awesomeProject/json_parse"
	"github.com/gin-gonic/gin"
	"net/http"
)

var OrdersCache = make(map[string]json_parse.Order)

func StartServer() {
	router := gin.Default()
	router.GET("/order/:id", GetOrderById)
	router.Run("localhost:8090")
}

func GetСacheFromDb() {
	database.CreateDbTablesIfNotExist()
	ordersFromDB := database.GetAllOrdersFromDB()

	//Перебор слайса Ордеров из бд
	for _, elem := range ordersFromDB {
		OrdersCache[elem.OrderUID] = elem
	}
}

func GetOrderById(c *gin.Context) {
	id := c.Param("id")

	if val, ok := OrdersCache[id]; ok {
		c.IndentedJSON(http.StatusOK, val)
	} else {
		c.IndentedJSON(http.StatusNotFound, "Ошибка! Такого id не существует.")
	}
}
