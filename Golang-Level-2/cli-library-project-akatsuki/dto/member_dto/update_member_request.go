package member_dto

/*
UpdateMemberRequestUser adalah INPUT user
saat ingin mengubah member

ID wajib karena kita harus tahu member mana yang diubah
*/
type UpdateMemberRequestUser struct {
	ID     int
	Nama   string
	Patner string
	Rank   string
}
