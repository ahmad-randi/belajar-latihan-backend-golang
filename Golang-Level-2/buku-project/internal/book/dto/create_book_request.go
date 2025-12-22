package dto

/*
CreateBookRequest adalah DATA INPUT dari user
saat menambahkan buku baru

DTO hanya berisi data yang DIMASUKKAN USER
*/
type CreateBookRequest struct {
	Title       string
	Author      string
	Description string
}
