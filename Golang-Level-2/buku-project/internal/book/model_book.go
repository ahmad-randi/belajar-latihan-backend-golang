package book

import "time"

/*
Book adalah REPRESENTASI DATA BUKU di sistem

Ini bukan input user,
tapi bentuk data FINAL yang disimpan di repository
*/
type Book struct {
	ID          int       // ID unik buku
	Title       string    // Judul buku
	Author      string    // Penulis
	Description string    // Deskripsi singkat
	IsBorrowed  bool      // Status: sedang dipinjam atau tidak
	CreatedAt   time.Time // Waktu buku dibuat
	UpdatedAt   time.Time // Waktu terakhir diubah
}
