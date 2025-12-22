package book

import "errors"

/*
Error didefinisikan TERPUSAT
agar:
- konsisten
- mudah dirawat
*/
var (
	ErrBookNotFound        = errors.New("buku tidak ditemukan")
	ErrInvalidBookInput    = errors.New("input buku tidak valid")
	ErrBookAlreadyBorrowed = errors.New("buku sudah dipinjam")
	ErrBookNotBorrowed     = errors.New("buku belum dipinjam")
)
