package book

/*
Helper berisi fungsi bantu
yang tidak termasuk business logic utama
*/

// counter global untuk ID buku
var lastBookID int

/*
GenerateBookID bertugas membuat ID unik
ID tidak akan bentrok walau data dihapus
*/
func GenerateBookID() int {
	lastBookID++
	return lastBookID
}
