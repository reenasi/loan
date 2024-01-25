package dao

import (
	"awesomeProject/model"
	"database/sql"
)

type TransactionDAO struct {
	DB *sql.DB
}

func (dao *TransactionDAO) RecordTransaction(transaction *model.Transaction) error {
	query := "INSERT INTO transactions (contract_number, otr, admin_fee, installment_amount, interest_amount, asset_name) " +
		"VALUES (?, ?, ?, ?, ?, ?)"

	result, err := dao.DB.Exec(query,
		transaction.ContractNumber,
		transaction.OTR,
		transaction.AdminFee,
		transaction.InstallmentAmount,
		transaction.InterestAmount,
		transaction.AssetName,
	)
	if err != nil {
		return err
	}

	transactionID, _ := result.LastInsertId()
	transaction.ID = int(transactionID)

	return nil
}

func (dao *TransactionDAO) GetAllTransactions() ([]model.Transaction, error) {
	query := "SELECT trx.id, c.id, trx.contract_number, trx.otr, trx.admin_fee, " +
		" trx.installment_amount, trx.interest_amount, trx.asset_name " +
		" FROM transactions trx " +
		" LEFT JOIN customers c ON c.id = trx.customer_id "

	rows, err := dao.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.Transaction
	for rows.Next() {
		var transaction model.Transaction
		err := rows.Scan(
			&transaction.ID,
			&transaction.ContractNumber,
			&transaction.OTR,
			&transaction.AdminFee,
			&transaction.InstallmentAmount,
			&transaction.InterestAmount,
			&transaction.AssetName,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
