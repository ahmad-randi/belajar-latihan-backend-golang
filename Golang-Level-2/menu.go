package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"Golang-Level-2/service"
)

// =======================================================
// SHOW MENU
// =======================================================
// Fungsi ini adalah "loop hidup" aplikasi
// Selama user belum memilih menu 0 (keluar),
// program akan terus berjalan
func ShowMenu() {

	// Reader digunakan untuk membaca input terminal secara aman
	// Lebih stabil dibanding fmt.Scanln
	reader := bufio.NewReader(os.Stdin)

	for {
		// ==============================
		// TAMPILKAN MENU
		// ==============================
		fmt.Println("\n============================")
		fmt.Println("      CRUD USERS CLI")
		fmt.Println("============================")
		fmt.Println("1. Tambah User")
		fmt.Println("2. Lihat Semua User")
		fmt.Println("3. Update User")
		fmt.Println("4. Hapus User")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		// ==============================
		// BACA INPUT MENU
		// ==============================
		// Input dibaca sebagai string agar:
		// - tidak crash
		// - bisa divalidasi dulu
		inputMenu, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Gagal membaca input")
			continue
		}

		// Bersihkan spasi & newline
		inputMenu = strings.TrimSpace(inputMenu)

		// Konversi string ke integer
		menu, err := strconv.Atoi(inputMenu)
		if err != nil {
			fmt.Println("Menu harus berupa angka")
			continue
		}

		// ==============================
		// ARAHKAN KE FUNGSI SESUAI MENU
		// ==============================
		switch menu {
		case 1:
			createUserCLI(reader)
		case 2:
			readUserCLI()
		case 3:
			updateUserCLI(reader)
		case 4:
			deleteUserCLI(reader)
		case 0:
			// Keluar dari aplikasi
			fmt.Println("Program selesai. Terima kasih 🙏")
			return
		default:
			fmt.Println("Menu tidak tersedia")
		}
	}
}

// =======================================================
// CREATE USER (CLI)
// =======================================================
// Tugas fungsi ini:
// 1. Ambil input dari user
// 2. Kirim ke service
// 3. Tampilkan hasil
func createUserCLI(reader *bufio.Reader) {

	// ==============================
	// INPUT DATA
	// ==============================
	fmt.Print("Nama   : ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Umur   : ")
	umurStr, _ := reader.ReadString('\n')
	umurStr = strings.TrimSpace(umurStr)

	// Konversi umur ke int
	umur, err := strconv.Atoi(umurStr)
	if err != nil {
		fmt.Println("Umur harus angka")
		return
	}

	fmt.Print("Alamat : ")
	alamat, _ := reader.ReadString('\n')
	alamat = strings.TrimSpace(alamat)

	// ==============================
	// KIRIM KE SERVICE
	// ==============================
	// Validasi & penyimpanan ada di service
	err = service.CreateUser(nama, umur, alamat)
	if err != nil {
		fmt.Println("Gagal:", err.Error())
		return
	}

	fmt.Println("User berhasil ditambahkan ✅")
}

// =======================================================
// READ USER (CLI)
// =======================================================
// Menampilkan seluruh data user
func readUserCLI() {

	// Ambil data dari service
	users := service.GetAllUsers()

	// Jika belum ada data
	if len(users) == 0 {
		fmt.Println("Belum ada data user")
		return
	}

	// ==============================
	// TAMPILKAN DATA
	// ==============================
	fmt.Println("\n====== DAFTAR USERS ======")
	for _, user := range users {
		fmt.Printf(
			"ID: %d | Nama: %s | Umur: %d | Alamat: %s\n",
			user.ID,
			user.Nama,
			user.Umur,
			user.Alamat,
		)
	}
}

// =======================================================
// UPDATE USER (CLI)
// =======================================================
// Update user berdasarkan ID
func updateUserCLI(reader *bufio.Reader) {

	// ==============================
	// INPUT ID
	// ==============================
	fmt.Print("ID User : ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID harus angka")
		return
	}

	// ==============================
	// INPUT DATA BARU
	// ==============================
	fmt.Print("Nama Baru   : ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Umur Baru   : ")
	umurStr, _ := reader.ReadString('\n')
	umurStr = strings.TrimSpace(umurStr)

	umur, err := strconv.Atoi(umurStr)
	if err != nil {
		fmt.Println("Umur harus angka")
		return
	}

	fmt.Print("Alamat Baru : ")
	alamat, _ := reader.ReadString('\n')
	alamat = strings.TrimSpace(alamat)

	// ==============================
	// KIRIM KE SERVICE
	// ==============================
	err = service.UpdateUser(id, nama, umur, alamat)
	if err != nil {
		fmt.Println("Gagal:", err.Error())
		return
	}

	fmt.Println("User berhasil diupdate ✅")
}

// =======================================================
// DELETE USER (CLI)
// =======================================================
// Hapus user berdasarkan ID
func deleteUserCLI(reader *bufio.Reader) {

	// ==============================
	// INPUT ID
	// ==============================
	fmt.Print("ID User : ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID harus angka")
		return
	}

	// ==============================
	// KIRIM KE SERVICE
	// ==============================
	err = service.DeleteUser(id)
	if err != nil {
		fmt.Println("Gagal:", err.Error())
		return
	}

	fmt.Println("User berhasil dihapus ✅")
}
