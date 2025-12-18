package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"user-project/internal/user"
	"user-project/internal/user/dto"
)

// ShowMenu adalah loop utama aplikasi CLI
// Program terus berjalan sampai user memilih menu 0
func ShowMenu() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n===== CRUD USER CLI =====")
		fmt.Println("1. Tambah User")
		fmt.Println("2. Lihat Semua User")
		fmt.Println("3. Update User")
		fmt.Println("4. Hapus User")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		input, _ := reader.ReadString('\n')
		menu, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("Menu harus angka")
			continue
		}

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
			return
		default:
			fmt.Println("Menu tidak tersedia")
		}
	}
}

// ================= CREATE =================
func createUser(reader *bufio.Reader) {
	fmt.Print("Nama: ")
	nama, _ := reader.ReadString('\n')

	fmt.Print("Umur: ")
	umurStr, _ := reader.ReadString('\n')
	umur, _ := strconv.Atoi(strings.TrimSpace(umurStr))

	fmt.Print("Alamat: ")
	alamat, _ := reader.ReadString('\n')

	req := dto.CreateUserRequest{
		Nama:   strings.TrimSpace(nama),
		Umur:   umur,
		Alamat: strings.TrimSpace(alamat),
	}

	if err := user.CreateUserUsecase(req); err != nil {
		fmt.Println("Gagal:", err.Error())
		return
	}

	fmt.Println("User berhasil ditambahkan")
}

// ================= READ =================
func readUsers() {
	users := user.GetAllUsers()

	if len(users) == 0 {
		fmt.Println("Belum ada data user")
		return
	}

	for _, u := range users {
		fmt.Printf("ID:%d | %s | %d | %s\n", u.ID, u.Nama, u.Umur, u.Alamat)
	}
}

// ================= UPDATE =================
func updateUser(reader *bufio.Reader) {
	fmt.Print("ID: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Nama baru: ")
	nama, _ := reader.ReadString('\n')

	fmt.Print("Umur baru: ")
	umurStr, _ := reader.ReadString('\n')
	umur, _ := strconv.Atoi(strings.TrimSpace(umurStr))

	fmt.Print("Alamat baru: ")
	alamat, _ := reader.ReadString('\n')

	req := dto.UpdateUserRequest{
		ID:     id,
		Nama:   strings.TrimSpace(nama),
		Umur:   umur,
		Alamat: strings.TrimSpace(alamat),
	}

	if err := user.UpdateUserUsecase(req); err != nil {
		fmt.Println("Gagal:", err.Error())
		return
	}

	fmt.Println("User berhasil diupdate")
}

// ================= DELETE =================
func deleteUser(reader *bufio.Reader) {
	fmt.Print("ID: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	if err := user.DeleteUserUsecase(id); err != nil {
		fmt.Println("Gagal:", err.Error())
		return
	}

	fmt.Println("User berhasil dihapus")
}
