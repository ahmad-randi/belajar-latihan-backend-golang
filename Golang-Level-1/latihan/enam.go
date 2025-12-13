package latihan

func FilterdanTransform(nums []int) []int {
	var hasil []int // Siapkan slice kosong hasil
	//Loop semua angka
	for _, num := range nums {
		//Cek apakah ganjil
		if num%2 != 0 {
			// Jika ya kalikan 3
			hasil = append(hasil, num*3)
		}
	}
	//Balikan nilai slice
	return hasil
}

func GenapPertama(nums []int) int {
	//Loop dari awal slice
	for _, num := range nums {
		//Cek apakah angka genap
		if num%2 == 0 {
			//Jika ya, maka langsung retrun nilai genap
			return num
		}
	}
	//kembalikan nilai jika tidak genap
	return -1
}

func VariabelSliceKosong(nums []int) int {
	//Cek panjang slice
	if len(nums) == 0 {
		//Jika 0, return 0
		return 0
	}

	//Ambil element pertama sebagai max
	max := nums[0]

	//Bandingkan semua element
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	//Kembalikan nilai max
	return max
}

func LogikaString(data []string) int {
	//Membuat variabel counter 0
	counter := 0

	//Loop semua string
	for _, str := range data {
		//Loop tiap character string
		for _, ch := range str {
			//Jika ketemu 'a', counter++ stop cek string itu
			if ch == 'a' {
				counter++
				break
			}
		}
	}
	//Kembalikan nilai counter
	return counter
}

func SliceCampuran(data []interface{}) []interface{} {
	//Buat dua variabel counter angka, dan text
	angka := 0
	text := 0

	//Loop semua item element
	for _, item := range data {
		//Cek tipe data
		switch item.(type) { // Cek tipe data elemen
		//Jika ya, tambah type int counter angka++
		case int:
			angka++ // Jika int → tambah counter angka
		case string:
			//Jika ya, tambah type string counter text++
			text++ // Jika string → tambah counter teks
		}
	}

	//Menginisialkan variabel hasil sebagai campuran
	hasil := []interface{}{angka, text}
	//Kembalikan nilai hasil
	return hasil
}
