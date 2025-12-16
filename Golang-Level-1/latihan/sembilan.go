package latihan

//Membuat function soal ke-1 mudah
func MenghitungBerapaAngkaGenap(data []int) int {
	//Membuat variabel counter
	counter := 0

	//Loop isi slice
	for _, value := range data {
		//Cek apakah genap?
		if value%2 == 0 {
			//Jika ya, maka counter++
			counter++
		}
	}
	//Kembalikan nilai counter
	return counter
}

//Membuat function soal ke-2 mudah
func JumlahStringYangPanjangLebihDariN(data []string, n int) int {
	//Membuat variabel counter
	counter := 0

	//Loop isi data slice
	for _, value := range data {
		//Cek panjang isi data slice
		//Jika lebih besar dari n
		if len(value) > n {
			//Maka counter++
			counter++
		}
	}
	//Kembalikan nilai counter
	return counter
}

//Membuat function soal ke-1 sedang
func AmbilAngkaGajilPertamaSlice(data []int) int {
	//Loop isi data slice
	for _, value := range data {
		//Cek apakah ganjil?
		if value%2 != 0 {
			//Jika ya, maka retrun value
			return value
		}
	}
	//Kembalikan nilai, jika loop tidak menemukan angka ganjil
	return -1
}

//Membuat function soal ke-2 mudah
func NilaiTerbesarDiSlice(data []int) int {
	//Membuat variable slice pertama sebagai nilai terbesar
	result := data[0]

	//Loop semua isi slice
	for _, value := range data {
		//Cek apakah data lebih besar dari result
		if value > result {
			//Jika ya maka update baru isi result
			result = value
		}
	}
	//Kembalikan nilai result
	return result
}

func BerapaStringYangMengandungHurufVokal(data []string) int {
	//Membuat variabel counter
	counter := 0
	//Loop isi data slice
	for _, value := range data {
		//Loop isi data slice perkarakter
		for _, ch := range value {
			//Cek apakah ada huruf vokal (a, i, u, e, o)
			if ch == 'a' || ch == 'i' || ch == 'u' || ch == 'e' || ch == 'o' {
				//Jika ya, counter, ketemu lalu break dan lanjutkan kembali loop
				counter++
				break
			}
		}
	}
	//Kembalikan nilai counter
	return counter
}
