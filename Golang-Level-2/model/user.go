package model

// User merepresentasikan data user
// Sama seperti tabel users di database nanti
type User struct {
	ID     int    // ID unik (dibuat sistem)
	Nama   string // Nama user
	Umur   int    // Umur user
	Alamat string // Alamat user
}
