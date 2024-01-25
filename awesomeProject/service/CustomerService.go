package service

import (
	"awesomeProject/dao"
	"awesomeProject/dto"
	"awesomeProject/model"
	"errors"
)

type CustomerService struct {
	CustomerDAO *dao.CustomerDAO
}

func (service *CustomerService) SignUp(customerDTO *dto.CustomerDTO) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customerDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	customer := &model.Customer{
		NIK:          customerDTO.NIK,
		FullName:     customerDTO.FullName,
		LegalName:    customerDTO.LegalName,
		PlaceOfBirth: customerDTO.PlaceOfBirth,
		DateOfBirth:  customerDTO.DateOfBirth,
		Salary:       customerDTO.Salary,
		KTPImage:     customerDTO.KTPImage,
		SelfieImage:  customerDTO.SelfieImage,
		Password:     string(hashedPassword),
	}

	return service.CustomerDAO.InsertCustomer(customer)
}

func (service *CustomerService) Login(NIK int, password string) (*dto.CustomerDTO, error) {
	customer, err := service.CustomerDAO.GetCustomerByNIK(NIK)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	customerDTO := &dto.CustomerDTO{
		NIK:          customer.NIK,
		FullName:     customer.FullName,
		LegalName:    customer.LegalName,
		PlaceOfBirth: customer.PlaceOfBirth,
		DateOfBirth:  customer.DateOfBirth,
		Salary:       customer.Salary,
		KTPImage:     customer.KTPImage,
		SelfieImage:  customer.SelfieImage,
	}

	return customerDTO, nil
}
