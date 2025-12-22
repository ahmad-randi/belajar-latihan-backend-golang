package cli_book

import (
	"bufio"
	"buku-project/internal/book"
	"buku-project/internal/book/dto"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Run adalah FUNGSI UTAMA CLI

Tugas fungsi ini:
- Menampilkan menu
- Menerima input user
- Mengarahkan ke fungsi yang sesuai

❗ Tidak ada LOGIKA BISNIS di sini
❗ Semua logika ada di service layer
*/
func Run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// ===== Tampilan Menu =====
		fmt.Println("\n===== BOOK PROJECT CLI =====")
		fmt.Println("1. Tambah Buku")
		fmt.Println("2. Lihat Semua Buku")
		fmt.Println("3. Ubah Buku")
		fmt.Println("4. Pinjam Buku")
		fmt.Println("5. Kembalikan Buku")
		fmt.Println("6. Hapus Buku")
		fmt.Println("0. Keluar")

		fmt.Print("Pilih menu: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		// ===== Routing Menu =====
		switch choice {
		case "1":
			menuCreateBook(reader)
		case "2":
			menuListBooks()
		case "3":
			menuUpdateBook(reader)
		case "4":
			menuBorrowBook(reader)
		case "5":
			menuReturnBook(reader)
		case "6":
			menuDeleteBook(reader)
		case "0":
			fmt.Println("Terima kasih, program selesai 👋")
			return
		default:
			fmt.Println("Menu tidak valid")
		}
	}
}

/*
menuCreateBook menangani input user
untuk MENAMBAH buku baru
*/
func menuCreateBook(reader *bufio.Reader) {
	fmt.Print("Judul Buku: ")
	title, _ := reader.ReadString('\n')

	fmt.Print("Author: ")
	author, _ := reader.ReadString('\n')

	fmt.Print("Deskripsi: ")
	description, _ := reader.ReadString('\n')

	req := dto.CreateBookRequest{
		Title:       strings.TrimSpace(title),
		Author:      strings.TrimSpace(author),
		Description: strings.TrimSpace(description),
	}

	if err := book.CreateBook(req); err != nil {
		fmt.Println("Gagal:", err.Error())
		return
	}

	fmt.Println("Buku berhasil ditambahkan ✅")
}

/*
menuListBooks menampilkan seluruh buku
tanpa input user
*/
func menuListBooks() {
	books := book.GetAllBooks()

	if len(books) == 0 {
		fmt.Println("Belum ada buku")
		return
	}

	for _, b := range books {
		status := "Tersedia"
		if b.IsBorrowed {
			status = "Dipinjam"
		}

		fmt.Printf(
			"ID: %d | %s | %s | %s\n",
			b.ID,
			b.Title,
			b.Author,
			status,
		)
	}
}

/*
menuUpdateBook menangani proses UPDATE buku
*/
func menuUpdateBook(reader *bufio.Reader) {
	fmt.Print("ID Buku: ")
	inputID, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(inputID))
	if err != nil {
		fmt.Println("ID harus angka")
		return
	}

	fmt.Print("Judul Baru: ")
	title, _ := reader.ReadString('\n')

	fmt.Print("Author Baru: ")
	author, _ := reader.ReadString('\n')

	fmt.Print("Deskripsi Baru: ")
	description, _ := reader.ReadString('\n')

	req := dto.UpdateBookRequest{
		ID:          id,
		Title:       strings.TrimSpace(title),
		Author:      strings.TrimSpace(author),
		Description: strings.TrimSpace(description),
	}

	if err := book.UpdateBook(req); err != nil {
		fmt.Println("Gagal:", err.Error())
		return
	}

	fmt.Println("Buku berhasil diubah ✏️")
}

/*
menuBorrowBook menangani peminjaman buku
*/
func menuBorrowBook(reader *bufio.Reader) {
	fmt.Print("ID Buku: ")
	input, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("ID harus angka")
		return
	}

	if err := book.BorrowBook(id); err != nil {
		fmt.Println("Gagal:", err.Error())
		return
	}

	fmt.Println("Buku berhasil dipinjam 📚")
}

/*
menuReturnBook menangani pengembalian buku
*/
func menuReturnBook(reader *bufio.Reader) {
	fmt.Print("ID Buku: ")
	input, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("ID harus angka")
		return
	}

	if err := book.ReturnBook(id); err != nil {
		fmt.Println("Gagal:", err.Error())
		return
	}

	fmt.Println("Buku berhasil dikembalikan 🔄")
}

/*
menuDeleteBook menangani penghapusan buku
*/
func menuDeleteBook(reader *bufio.Reader) {
	fmt.Print("ID Buku: ")
	input, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("ID harus angka")
		return
	}

	if err := book.DeleteBook(id); err != nil {
		fmt.Println("Gagal:", err.Error())
		return
	}

	fmt.Println("Buku berhasil dihapus 🗑️")
}
