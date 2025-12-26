package member_dto

/*
CreateMemberRequest adalah DATA INPUT dari user
saat menambahkan member baru

DTO hanya berisi data yang DIMASUKKAN USER
*/
type CreateMemberRequestUser struct {
	Nama   string
	Patner string
	Rank   string
}
