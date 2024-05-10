package transaction

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var transactions = []Transaction{
	{ID: 1, Value: 10},
	{ID: 2, Value: 25},
}

var transactionCounter = int64(len(transactions))

func GetTransactions(context *gin.Context) {
	context.JSON(http.StatusOK, transactions)
}

func GetTransactionByID(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid transaction ID"})
		log.Println("Error parsing transaction ID:", err)
		return
	}

	for _, t := range transactions {
		if t.ID == id {
			context.JSON(http.StatusOK, t)
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"message": "Transaction not found"})
}

func CreateTransaction(context *gin.Context) {
	var newTransaction Transaction
	if err := context.ShouldBindJSON(&newTransaction); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		log.Println("Error parsing request payload:", err)
		return
	}

	transactionCounter++
	newTransaction.ID = transactionCounter

	transactions = append(transactions, newTransaction)

	context.JSON(http.StatusCreated, newTransaction)
}

func UpdateTransactionByID(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could not parse ID!"})
		return
	}

	type transactionUpdate struct {
		Value float64 `json:"value"`
	}

	var updatedTransaction transactionUpdate
	if err := context.ShouldBindJSON(&updatedTransaction); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid transaction data!"})
		return
	}

	for i, t := range transactions {
		if t.ID == id {
			transactions[i].Value = updatedTransaction.Value
			context.JSON(http.StatusOK, gin.H{"message": "Transaction updated successfully!"})
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"message": "Transaction not found!"})
}

func DeleteTransactionByID(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could not parse ID!"})
		return
	}

	for i, t := range transactions {
		if t.ID == id {
			transactions = append(transactions[:i], transactions[i+1:]...)
			context.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully!"})
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"message": "Transaction not found!"})
}