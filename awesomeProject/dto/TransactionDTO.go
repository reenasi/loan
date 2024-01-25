package dto

type RecordTransactionRequest struct {
	CustNIK     int    `json:"cust_nik" binding:"required"`
	TenorMonths int    `json:"tenor_months" binding:"required"`
	Loan        int    `json:"loan" binding:"required"`
	CodeItem    string `json:"code_item" binding:"required"`
}
