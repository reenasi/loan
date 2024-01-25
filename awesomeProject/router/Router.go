package router

import (
	"awesomeProject/endpoint"
	"awesomeProject/service"
	"github.com/gin-gonic/gin"
)

type Router struct {
	CustomerService    *service.CustomerService
	TransactionService *service.TransactionService
}

func (router *Router) SetupRoutes(r *gin.Engine) {
	customerEndpoint := endpoint.CustomerEndpoint{CustomerService: router.CustomerService}
	transactionEndpoint := endpoint.TransactionEndpoint{TransactionService: router.TransactionService}

	r.POST("/signup", customerEndpoint.SignUp)
	r.POST("/login", customerEndpoint.Login)
	r.POST("/applyloan", transactionEndpoint.RecordTransaction)
	r.GET("/transactions", func(c *gin.Context) {
		transactionEndpoint.GetAllTransactions(c, router.TransactionService)
	})
}
