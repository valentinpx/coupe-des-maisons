//
// EPITECH PROJECT, 2021
// cdm-api
// File description:
// main
//

package main

import (
	"fmt"

	"os"

	"time"

	"math/rand"

	"net/http"

	"database/sql"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

type Transaction struct {
	House       string  `json:"house"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Author      string  `json:"author"`
	Date        string  `json:"date"`
}

type TransactionRequest struct {
	Transaction Transaction `json:"transaction"`
	Key         string      `json:"key"`
}

var DB *sql.DB

var KEY string

// Databases
func createDb(db *sql.DB) {
	db.Exec(`CREATE TABLE IF NOT EXISTS transactions(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		House TEXT,
		Description TEXT,
		Amount INTEGER,
		Author TEXT,
		Date TEXT
	);`)
}

func initDb(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)

	if err != nil || db == nil {
		return nil
	}
	return db
}

func insertTransaction(db *sql.DB, t Transaction) {
	stmt, err := db.Prepare(`
		INSERT INTO transactions(
			House,
			Description,
			Amount,
			Author,
			Date
		) VALUES(?,?,?,?,?)
	`)

	if err != nil {
		return
	}
	defer stmt.Close()
	stmt.Exec(t.House, t.Description, t.Amount, t.Author, t.Date)
}

func selectTransactions(db *sql.DB) []Transaction {
	rows, err := db.Query(`
		SELECT House, Description, Amount, Author, Date FROM transactions
	`)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var dest []Transaction
	for rows.Next() {
		item := Transaction{}
		err = rows.Scan(&item.House, &item.Description, &item.Amount, &item.Author, &item.Date)
		if err != nil {
			return []Transaction{}
		}
		dest = append(dest, item)
	}
	return dest
}

func sumHouseAmounts(house string) int {
	var dest int
	rows, err := DB.Query(`
		SELECT SUM(Amount) as total
		FROM transactions
		WHERE House = ?
	`, house)

	if err != nil {
		return 0
	}
	defer rows.Close()
	rows.Next()
	rows.Scan(&dest)
	return dest
}

// Routes
func getHouseTotal(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, sumHouseAmounts(context.Param("house")))
}

func getTransactions(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, selectTransactions(DB))
}

func postTransactions(context *gin.Context) {
	var req TransactionRequest
	err := context.BindJSON(&req)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	if req.Key != KEY {
		context.IndentedJSON(http.StatusForbidden, "Wrong password")
		return
	}
	req.Transaction.Date = time.Now().Format("02/01/2006 15:04:05")
	insertTransaction(DB, req.Transaction)
	context.IndentedJSON(http.StatusCreated, req.Transaction)
}

func serRouter(url string) *gin.Engine {
	router := gin.Default()

	router.GET("/houses/:house/total", getHouseTotal)
	router.GET("/transactions", getTransactions)
	router.POST("/transactions", postTransactions)
	router.Run(url)
	return router
}

// Main
func keyGen(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"
	bytes := make([]byte, length)

	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < length; i++ {
		bytes[i] = chars[rand.Intn(len(chars))]
	}
	return string(bytes)
}

func main() {
	DB = initDb(os.Args[1])
	KEY = keyGen(42)
	fmt.Printf("Transaction post key is: %s\n", KEY)

	defer DB.Close()
	createDb(DB)

	gin.SetMode(gin.ReleaseMode)
	serRouter(os.Args[2])
}
