package service

import (
	"Golang-Level-2/model"
	"errors"
	"strings"
)

// users adalah database sementara (in-memory)
// Data disimpan di RAM
// Program mati → data hilang (NORMAL di latihan)
var users []model.User

// lastID digunakan untuk auto increment ID
// Supaya ID selalu unik
var lastID int = 0

// FindUserByID bertugas mencari user berdasarkan ID
// Mengembalikan:
// - index: posisi user di slice
// - *User: pointer ke data user
// - nil jika user tidak ditemukan
func FindUserByID(id int) (int, *model.User) {

	// Loop semua data users
	for i := 0; i < len(users); i++ {

		// Cek apakah ID sama
		if users[i].ID == id {

			// Jika ketemu:
			// - kembalikan index
			// - kembalikan pointer ke user
			return i, &users[i]
		}
	}

	// Jika tidak ketemu sama sekali
	return -1, nil
}

// =======================
// CREATE
// =======================

// CreateUser bertugas menambahkan user baru
func CreateUser(nama string, umur int, alamat string) error {

	// ===== VALIDASI (SECURE CODING) =====

	// TrimSpace untuk menghindari input spasi kosong
	if strings.TrimSpace(nama) == "" {
		return errors.New("nama tidak boleh kosong")
	}

	// Aturan bisnis: umur minimal 17
	if umur < 17 {
		return errors.New("umur minimal 17 tahun")
	}

	if strings.TrimSpace(alamat) == "" {
		return errors.New("alamat tidak boleh kosong")
	}

	// ===== BUAT DATA USER =====

	// Auto increment ID
	lastID++

	// Buat object user
	user := model.User{
		ID:     lastID,
		Nama:   nama,
		Umur:   umur,
		Alamat: alamat,
	}

	// Simpan ke slice (database)
	users = append(users, user)

	return nil
}

// =======================
// READ
// =======================

// GetAllUsers mengembalikan seluruh data user
func GetAllUsers() []model.User {

	// Service TIDAK print
	// Service hanya mengembalikan data
	return users
}

// =======================
// UPDATE
// =======================

// UpdateUser bertugas mengubah data user
func UpdateUser(id int, nama string, umur int, alamat string) error {

	// Cari user dulu
	_, user := FindUserByID(id)

	// Jika user tidak ditemukan
	if user == nil {
		return errors.New("user tidak ditemukan")
	}

	// ===== VALIDASI =====

	if strings.TrimSpace(nama) == "" {
		return errors.New("nama tidak boleh kosong")
	}

	if umur < 17 {
		return errors.New("umur minimal 17 tahun")
	}

	if strings.TrimSpace(alamat) == "" {
		return errors.New("alamat tidak boleh kosong")
	}

	// ===== UPDATE DATA =====

	// Karena user adalah pointer,
	// perubahan langsung kena ke slice
	user.Nama = nama
	user.Umur = umur
	user.Alamat = alamat

	return nil
}

// =======================
// DELETE
// =======================

// DeleteUser menghapus user berdasarkan ID
func DeleteUser(id int) error {

	// Cari user
	index, _ := FindUserByID(id)

	// Jika tidak ditemukan
	if index == -1 {
		return errors.New("user tidak ditemukan")
	}

	// Hapus data dengan cara:
	// Gabungkan slice sebelum & sesudah index
	users = append(users[:index], users[index+1:]...)

	return nil
}
