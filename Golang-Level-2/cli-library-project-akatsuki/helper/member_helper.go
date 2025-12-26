package helper

/*
Helper berisi fungsi bantu
yang tidak termasuk business logic utama
*/

// counter global untuk ID member
var lastMemberID int

/*
GenerateMemberID bertugas membuat ID unik
ID tidak akan bentrok walau data dihapus
*/

func GenerateMemberID() int {
	lastMemberID++
	return lastMemberID
}
