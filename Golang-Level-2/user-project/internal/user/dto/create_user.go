package dto

// DTO untuk CREATE user
// Digunakan sebagai input dari CLI
type CreateUserRequest struct {
	Nama   string
	Umur   int
	Alamat string
}
