package model

type Customer struct {
	ID           int
	NIK          int
	FullName     string
	LegalName    string
	PlaceOfBirth string
	DateOfBirth  string
	Salary       int
	KTPImage     []byte
	SelfieImage  []byte
	Password     string
}
