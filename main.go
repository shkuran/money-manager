package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shkuran/money-manager/transaction"
)

func main() {
	router := gin.Default()
	router.GET("/transactions", transaction.GetTransactions)
	router.GET("/transactions/:id", transaction.GetTransactionByID)
	router.POST("/transactions", transaction.CreateTransaction)
	router.PUT("/transactions/:id", transaction.UpdateTransactionByID)
	router.DELETE("/transactions/:id", transaction.DeleteTransactionByID)

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
