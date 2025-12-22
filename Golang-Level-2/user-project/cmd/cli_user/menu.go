package cli_user

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	// user = layer usecase (alur bisnis)
	"user-project/internal/user"

	// dto = data input dari user (request)
	"user-project/internal/user/dto"
)

/*
========================================================
SHOW MENU
========================================================
Fungsi ini adalah:
- PINTU MASUK aplikasi CLI
- LOOP UTAMA aplikasi

Selama user belum memilih menu 0 (Keluar),
program akan TERUS BERJALAN.
*/
func ShowMenu() {

	// Reader digunakan untuk membaca input dari terminal
	// Lebih aman & fleksibel dibanding fmt.Scanln
	reader := bufio.NewReader(os.Stdin)

	// Infinite loop → aplikasi hidup terus
	for {

		// ===============================
		// TAMPILKAN MENU
		// ===============================
		fmt.Println("\n===== CRUD USER CLI =====")
		fmt.Println("1. Tambah User")
		fmt.Println("2. Lihat Semua User")
		fmt.Println("3. Update User")
		fmt.Println("4. Hapus User")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		// ===============================
		// AMBIL INPUT MENU
		// ===============================
		// Input dibaca sebagai string dulu
		// supaya tidak crash jika user salah input
		input, _ := reader.ReadString('\n')

		// Bersihkan spasi & newline
		input = strings.TrimSpace(input)

		// Konversi string → int
		menu, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Menu harus berupa angka")
			continue // balik ke menu
		}

		// ===============================
		// ARAHKAN KE FITUR SESUAI MENU
		// ===============================
		switch menu {
		case 1:
			createUser(reader)
		case 2:
			readUsers()
		case 3:
			updateUser(reader)
		case 4:
			deleteUser(reader)
		case 0:
			fmt.Println("Program selesai")
			return // keluar dari aplikasi
		default:
			fmt.Println("Menu tidak tersedia")
		}
	}
}

/*
========================================================
CREATE USER
========================================================
Tugas fungsi ini:
1. Ambil input dari user (CLI)
2. Bentuk DTO (CreateUserRequest)
3. Kirim ke Usecase
4. Tampilkan hasil
*/
func createUser(reader *bufio.Reader) {

	// ===============================
	// INPUT DATA
	// ===============================
	fmt.Print("Nama: ")
	nama, _ := reader.ReadString('\n')

	fmt.Print("Umur: ")
	umurStr, _ := reader.ReadString('\n')
	umur, _ := strconv.Atoi(strings.TrimSpace(umurStr))

	fmt.Print("Alamat: ")
	alamat, _ := reader.ReadString('\n')

	// ===============================
	// BENTUK DTO
	// ===============================
	// DTO = Data Transfer Object
	// Digunakan untuk:
	// - menampung data input
	// - memisahkan input dengan model domain
	req := dto.CreateUserRequest{
		Nama:   strings.TrimSpace(nama),
		Umur:   umur,
		Alamat: strings.TrimSpace(alamat),
	}

	// ===============================
	// KIRIM KE USECASE
	// ===============================
	// Usecase akan:
	// - validasi data
	// - generate ID
	// - simpan ke repository
	if err := user.CreateUserUsecase(req); err != nil {
		fmt.Println("Gagal:", err.Error())
		return
	}

	fmt.Println("User berhasil ditambahkan ✅")
}

/*
========================================================
READ USERS
========================================================
Tugas fungsi ini:
- Ambil semua data user
- Tampilkan ke terminal
*/
func readUsers() {

	// Ambil data dari usecase
	users := user.GetAllUsers()

	// Jika belum ada data
	if len(users) == 0 {
		fmt.Println("Belum ada data user")
		return
	}

	// ===============================
	// TAMPILKAN DATA
	// ===============================
	for _, u := range users {
		fmt.Printf(
			"ID:%d | Nama:%s | Umur:%d | Alamat:%s\n",
			u.ID,
			u.Nama,
			u.Umur,
			u.Alamat,
		)
	}
}

/*
========================================================
UPDATE USER
========================================================
Tugas fungsi ini:
1. Ambil ID user
2. Ambil data baru
3. Bentuk DTO
4. Kirim ke Usecase
*/
func updateUser(reader *bufio.Reader) {

	// ===============================
	// INPUT ID
	// ===============================
	fmt.Print("ID: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	// ===============================
	// INPUT DATA BARU
	// ===============================
	fmt.Print("Nama baru: ")
	nama, _ := reader.ReadString('\n')

	fmt.Print("Umur baru: ")
	umurStr, _ := reader.ReadString('\n')
	umur, _ := strconv.Atoi(strings.TrimSpace(umurStr))

	fmt.Print("Alamat baru: ")
	alamat, _ := reader.ReadString('\n')

	// ===============================
	// BENTUK DTO UPDATE
	// ===============================
	req := dto.UpdateUserRequest{
		ID:     id,
		Nama:   strings.TrimSpace(nama),
		Umur:   umur,
		Alamat: strings.TrimSpace(alamat),
	}

	// ===============================
	// KIRIM KE USECASE
	// ===============================
	// Usecase akan:
	// - cari user dengan FindByID
	// - validasi data baru
	// - update data di repository
	if err := user.UpdateUserUsecase(req); err != nil {
		fmt.Println("Gagal:", err.Error())
		return
	}

	fmt.Println("User berhasil diupdate ✅")
}

/*
========================================================
DELETE USER
========================================================
Tugas fungsi ini:
1. Ambil ID user
2. Kirim ke Usecase
3. Hapus data jika ditemukan
*/
func deleteUser(reader *bufio.Reader) {

	// ===============================
	// INPUT ID
	// ===============================
	fmt.Print("ID: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	// ===============================
	// KIRIM KE USECASE
	// ===============================
	// Usecase akan:
	// - cek apakah user ada
	// - hapus user dari repository
	if err := user.DeleteUserUsecase(id); err != nil {
		fmt.Println("Gagal:", err.Error())
		return
	}

	fmt.Println("User berhasil dihapus ✅")
}
