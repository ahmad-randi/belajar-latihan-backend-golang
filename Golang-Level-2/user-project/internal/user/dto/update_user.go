package dto

// DTO untuk UPDATE user
type UpdateUserRequest struct {
	ID     int
	Nama   string
	Umur   int
	Alamat string
}
