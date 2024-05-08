package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type transaction struct {
	ID int64 `json:"id"`
	Value float64 `json:"value"`
}

var transactions = []transaction {
	{ID: 1, Value: 10},
	{ID: 2, Value: 25},
}

func main() {
	router := gin.Default()
	router.GET("/transactions", getTransactions)
	router.GET("/transactions/:id", getTransactionByID)
	// router.POST("/transactions", postTransactions)

	router.Run("localhost:8080")
}

func getTransactions(context *gin.Context) {
	context.JSON(http.StatusOK, transactions)
}

func getTransactionByID(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could not parse Id!"})
		return
	}

	for _, t := range transactions {
		if t.ID == id {
			context.JSON(http.StatusOK, t)
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"message": "Could not found transaction!"})
}
