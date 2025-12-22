package dto

/*
UpdateBookRequest adalah INPUT user
saat ingin mengubah buku

ID wajib karena kita harus tahu buku mana yang diubah
*/
type UpdateBookRequest struct {
	ID          int
	Title       string
	Author      string
	Description string
}
