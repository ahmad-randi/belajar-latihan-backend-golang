package latihan

// Fungsi HitungPanjangKatadalamSlice menghitung jumlah kata yang panjangnya > 3
func HitungPanjangKatadalamSlice(slice []string) int {
	counter := 0 // Inisialisasi counter untuk menghitung kata panjang > 3

	// Loop setiap kata dalam slice
	for _, kata := range slice {
		// Cek panjang kata
		if len(kata) > 3 {
			counter++ // Jika > 3, tambahkan counter
		}
	}

	// Return total kata panjang > 3
	return counter
}

// Fungsi AmbilElemenStringPertamayangPanjangnyaGenap mengambil kata pertama dengan panjang genap
func AmbilElemenStringPertamayangPanjangnyaGenap(slice []string) string {
	// Loop semua kata dalam slice
	for _, kata := range slice {
		// Cek apakah panjang kata genap
		if len(kata)%2 == 0 {
			return kata // Jika genap, langsung return kata
		}
	}

	// Jika tidak ada kata panjang genap, return string kosong
	return ""
}

// Fungsi FilterSlicedenganAngkaLebihBesardari5 membuat slice baru dengan angka > 5
func FilterSlicedenganAngkaLebihBesardari5(slice []int) []int {
	var result []int // Inisialisasi slice kosong untuk hasil

	// Loop semua angka dalam slice
	for _, num := range slice {
		// Cek apakah angka > 5
		if num > 5 {
			result = append(result, num) // Tambahkan ke slice baru
		}
	}

	// Return slice baru berisi angka > 5
	return result
}

// Fungsi HitungJumlahAngkadanString menghitung jumlah angka dan string dalam slice interface{}
func HitungJumlahAngkadanString(items []interface{}) (int, int) {
	angka := 0 // Counter untuk angka
	teks := 0  // Counter untuk string

	// Loop semua elemen dalam slice
	for _, item := range items {
		switch item.(type) { // Cek tipe data elemen
		case int:
			angka++ // Jika int → tambah counter angka
		case string:
			teks++ // Jika string → tambah counter teks
		}
	}

	// Return jumlah angka dan string
	return angka, teks
}

// Fungsi ReverseSliceCampuran membalik urutan elemen slice interface{}
func ReverseSliceCampuran(slice []interface{}) []interface{} {
	var reversed []interface{} // Inisialisasi slice kosong baru

	// Loop dari index terakhir ke index pertama
	for i := len(slice) - 1; i >= 0; i-- {
		reversed = append(reversed, slice[i]) // Tambahkan elemen ke slice baru
	}

	// Return slice baru yang sudah dibalik
	return reversed
}
