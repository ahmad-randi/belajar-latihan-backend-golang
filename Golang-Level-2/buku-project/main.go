package main

import (
	"buku-project/cmd/cli_book"
)

/*
main.go adalah TITIK MASUK (entry point) aplikasi

Alur besar program:
1. Program dijalankan
2. Fungsi main() dipanggil otomatis oleh Go
3. main() memanggil cli.Run()
4. Seluruh kontrol aplikasi berpindah ke menu CLI
*/
func main() {
	// Menjalankan aplikasi CLI book-project
	cli_book.Run()
}
