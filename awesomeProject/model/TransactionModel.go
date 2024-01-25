package model

type Transaction struct {
	ID                int
	CustomerID        int
	CustomerNIK       int
	ContractNumber    int
	OTR               float64
	AdminFee          float64
	InstallmentAmount float64
	InterestAmount    float64
	AssetName         string
}
