package main

import (
	"awesomeProject/dao"
	"awesomeProject/dto"
	"awesomeProject/model"
	"awesomeProject/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignUp(t *testing.T) {
	customerDAO := &dao.CustomerDAO{}
	customerService := &service.CustomerService{CustomerDAO: customerDAO}

	customerDTO := &dto.CustomerDTO{
		NIK:          123456789,
		FullName:     "John Doe",
		LegalName:    "John",
		PlaceOfBirth: "City",
		DateOfBirth:  "1990-01-01",
		Salary:       50000,
		Password:     "password",
	}

	err := customerService.SignUp(customerDTO)

	assert.NoError(t, err)

}
func TestLogin(t *testing.T) {
	customerDAO := &dao.CustomerDAO{}
	customerService := &service.CustomerService{CustomerDAO: customerDAO}

	mockCustomer := &model.Customer{
		NIK:          123456789,
		FullName:     "John Doe",
		LegalName:    "John",
		PlaceOfBirth: "City",
		DateOfBirth:  "1990-01-01",
		Salary:       50000,
		Password:     "f4ebfd7a42d9a43a536e2bed9ee4974abf8f8dc8",
	}

	customerDAO.On("GetCustomerByNIK", 123456789).Return(mockCustomer, nil)

	customerDTO, err := customerService.Login(123456789, "123")

	assert.NoError(t, err)
	assert.NotNil(t, customerDTO)
	assert.Equal(t, mockCustomer.NIK, customerDTO.NIK)

	_, err = customerService.Login(123456789, "wrong_password")

	assert.Error(t, err)
	assert.EqualError(t, err, "invalid password")
}

func TestRecordTransaction(t *testing.T) {
	customerDAO := &dao.CustomerDAO{}
	itemDAO := &dao.ItemDAO{}

	mockCustomer := &model.Customer{
		ID:           1,
		NIK:          123456789,
		FullName:     "John Doe",
		LegalName:    "John",
		PlaceOfBirth: "City",
		DateOfBirth:  "1990-01-01",
		Salary:       50000,
	}

	mockItem := &model.Item{
		ID:    1,
		Code:  "CAR123",
		Name:  "Car",
		Type:  "Motor",
		Price: 10000,
	}

	mockTransaction := &model.Transaction{
		CustomerNIK:       mockCustomer.NIK,
		ContractNumber:    123456789,
		AdminFee:          50.0,
		InstallmentAmount: 2500.0,
		InterestAmount:    1000.0,
		AssetName:         mockItem.Name,
		OTR:               10500.0,
	}

	customerDAO.On("ApplyLoan", mockCustomer.NIK, 3, 1000).Return(true, 10000.0, mockTransaction, nil)
	itemDAO.On("GetItemByCode", mockItem.Code).Return(mockItem, nil)

	transactionDAO := &dao.TransactionDAO{}
	transactionDAO.On("RecordTransaction", mockTransaction).Return(nil)

	transactionService := &TransactionService{
		TransactionDAO: transactionDAO,
		CustomerDAO:    customerDAO,
		ItemDAO:        itemDAO,
	}

	err := transactionService.RecordTransaction(mockCustomer.NIK, 3, 1000, mockItem.Code)

	assert.NoError(t, err)
}

func TestGetAllTransactions(t *testing.T) {
	transactionDAO := &dao.TransactionDAO{}

	expectedTransactions := []model.Transaction{
		{
			ID:                1,
			CustomerNIK:       1980138923894201,
			ContractNumber:    123456789,
			AdminFee:          50.0,
			InstallmentAmount: 2500.0,
			InterestAmount:    1000.0,
			AssetName:         "Car",
			OTR:               10500.0,
		},
	}

	transactionDAO.On("GetAllTransactions").Return(expectedTransactions, nil)

	transactionService := &TransactionService{
		TransactionDAO: transactionDAO,
	}

	transactions, err := transactionService.GetAllTransactions()

	assert.NoError(t, err)
	assert.Equal(t, len(expectedTransactions), len(transactions))
}
