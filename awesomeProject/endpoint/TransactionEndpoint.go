package endpoint

import (
	"awesomeProject/dto"
	"awesomeProject/service"
	"net/http"
)

type TransactionEndpoint struct {
	TransactionService *service.TransactionService
}

func (endpoint *TransactionEndpoint) RecordTransaction(c *gin.Context) {
	var recordTransactionRequest dto.RecordTransactionRequest
	if err := c.ShouldBindJSON(&recordTransactionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := endpoint.TransactionService.RecordTransaction(
		recordTransactionRequest.CustNIK,
		recordTransactionRequest.TenorMonths,
		recordTransactionRequest.Loan,
		recordTransactionRequest.CodeItem,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error recording transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction recorded successfully"})
}

func (endpoint *TransactionEndpoint) GetAllTransactions(c *gin.Context, transactionService *service.TransactionService) {
	transactions, err := transactionService.GetAllTransactions()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get transactions"})
		return
	}

	c.JSON(200, transactions)
}
