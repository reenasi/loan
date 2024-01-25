package dto

type CustomerDTO struct {
	NIK          int    `json:"nik"`
	FullName     string `json:"full_name"`
	LegalName    string `json:"legal_name"`
	PlaceOfBirth string `json:"place_of_birth"`
	DateOfBirth  string `json:"date_of_birth"`
	Salary       int    `json:"salary"`
	KTPImage     []byte `json:"ktp_image"`
	SelfieImage  []byte `json:"selfie_image"`
	Password     string `json:"password"`
}

type LoginRequest struct {
	NIK      int    `json:"NIK"`
	Password string `json:"password"`
}
