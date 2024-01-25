package dao

import (
	"awesomeProject/model"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type CustomerDAO struct {
	DB *sql.DB
	TransactionDAO
}

func (dao *CustomerDAO) InsertCustomer(customer *model.Customer) error {
	query := "INSERT INTO Customers (full_name, legal_name, place_of_birth, date_of_birth, " +
		" salary, ktp_image, selfie_image, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?) "

	_, err := dao.DB.Exec(query,
		customer.FullName,
		customer.LegalName,
		customer.PlaceOfBirth,
		customer.DateOfBirth,
		customer.Salary,
		customer.KTPImage,
		customer.SelfieImage,
		customer.Password,
	)

	if err != nil {
		return err
	}

	return nil
}

func (dao *CustomerDAO) GetCustomerByNIK(nik int) (*model.Customer, error) {
	query := "SELECT id, nik, full_name, legal_name, place_of_birth," +
		" date_of_birth, salary, ktp_image, selfie_image, password " +
		" FROM customers " +
		" WHERE nik = ?"

	row := dao.DB.QueryRow(query, nik)
	var customer model.Customer

	err := row.Scan(
		&customer.ID,
		&customer.NIK,
		&customer.FullName,
		&customer.LegalName,
		&customer.PlaceOfBirth,
		&customer.DateOfBirth,
		&customer.Salary,
		&customer.KTPImage,
		&customer.SelfieImage,
		&customer.Password,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("customer not found")
		}
		return nil, err
	}

	return &customer, nil
}

func (dao *CustomerDAO) ApplyLoan(customerNIK, tenorMonths, loan int) (approvalAmount float64, transaction *model.Transaction, errS error) {
	customer, err := dao.GetCustomerByNIK(customerNIK)
	if err != nil {
		errS = err
		return
	}

	if customer.ID == 0 {
		errS = errors.New("customer not found")
		return
	}

	var loanLimit float64
	switch tenorMonths {
	case 1:
		loanLimit = CalculateLoanLimit(float64(customer.Salary), 1)
	case 2:
		loanLimit = CalculateLoanLimit(float64(customer.Salary), 2)
	case 3:
		loanLimit = CalculateLoanLimit(float64(customer.Salary), 3)
	case 4:
		loanLimit = CalculateLoanLimit(float64(customer.Salary), 4)
	default:
		errS = errors.New("unsupported tenor")
		return
	}

	isApproved, approvalAmount := loanApproval(float64(loan), loanLimit)
	if !isApproved {
		errS = errors.New("loan application rejected")
		return
	} else {
		transaction = &model.Transaction{
			CustomerID:        customer.ID,
			ContractNumber:    generateContractNumber(customer.ID),
			AdminFee:          calculateAdminFee(),
			InstallmentAmount: calculateInstallmentAmount(float64(loan), tenorMonths),
			InterestAmount:    calculateInterestAmount(float64(loan)),
		}
	}

	return approvalAmount, transaction, nil
}

func loanApproval(loan, loanLimit float64) (bool, float64) {
	return loan >= loanLimit, loanLimit
}

func CalculateLoanLimit(salary float64, tenorMonths int) float64 {
	var limit float64

	switch tenorMonths {
	case 1:
		limit = salary * 0.5
	case 2:
		limit = salary * 0.8
	case 3:
		limit = salary * 1.2
	case 4:
		limit = salary * 1.5
	default:
		limit = salary * 0.5
	}

	maxLimit := 100000000.0
	if limit > maxLimit {
		limit = maxLimit
	}

	return limit
}

func generateContractNumber(customerID int) int {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	contractNumber := fmt.Sprintf("%d%d", customerID, timestamp)
	result, _ := strconv.Atoi(contractNumber)

	return result
}

func calculateAdminFee() float64 {
	return 50.0
}

func calculateInstallmentAmount(loanAmount float64, tenorMonths int) float64 {
	return loanAmount / float64(tenorMonths)
}

func calculateInterestAmount(loanAmount float64) float64 {
	return loanAmount * 0.1
}
