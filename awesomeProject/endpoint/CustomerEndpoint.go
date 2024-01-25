package endpoint

import (
	"awesomeProject/dto"
	"awesomeProject/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CustomerEndpoint struct {
	CustomerService *service.CustomerService
}

func (endpoint *CustomerEndpoint) SignUp(c *gin.Context) {
	var customerDTO dto.CustomerDTO
	if err := c.ShouldBindJSON(&customerDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := endpoint.CustomerService.SignUp(&customerDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign up"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "SignUp successful"})
}

func (endpoint *CustomerEndpoint) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	NIK := loginRequest.NIK
	password := loginRequest.Password

	customerDTO, err := endpoint.CustomerService.Login(NIK, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error logging in"})
		return
	}

	c.JSON(http.StatusOK, customerDTO)
}
