//
// EPITECH PROJECT, 2021
// cdm-api
// File description:
// main
//

package main

import (
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

type transaction struct {
	House       string  `json:"house"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Author      string  `json:"author"`
	Date        string  `json:"date"`
}

var transactions = []transaction{
	{House: "Serdaigle", Description: "Défi AER 1", Amount: 100, Author: "Maxime", Date: "11/10/2021 16:46:00"},
	{House: "Gryffondor", Description: "Défi AER 1", Amount: 25, Author: "Maxime", Date: "11/10/2021 16:47:00"},
	{House: "Poufsouffle", Description: "Défi AER 1", Amount: 25, Author: "Maxime", Date: "11/10/2021 16:48:00"},
	{House: "Serpentard", Description: "Défi AER 1", Amount: 25, Author: "Maxime", Date: "11/10/2021 16:49:00"},
}

func getTransactions(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, transactions)
}

func postTransactions(context *gin.Context) {
	var new transaction
	err := context.BindJSON(&new)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	new.Date = time.Now().Format("02/01/2006 15:04:05")
	transactions = append(transactions, new)
	context.IndentedJSON(http.StatusCreated, new)
}

func serRouter(url string) *gin.Engine {
	router := gin.Default()

	router.GET("/transactions", getTransactions)
	router.POST("/transactions", postTransactions)
	router.Run(url)
	return router
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	serRouter("localhost:4242")
}
