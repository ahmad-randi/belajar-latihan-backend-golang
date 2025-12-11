package latihan

// Soal jawaban nomor 1
// Fungsi HitungJumlahAngkaGenapdanGanjil menghitung jumlah angka genap dan ganjil dalam slice
func HitungJumlahAngkaGenapdanGanjil(nums []int) (int, int) {
	genap := 0  // Counter untuk angka genap
	ganjil := 0 // Counter untuk angka ganjil

	// Loop semua angka dalam slice
	for _, num := range nums {
		if num%2 == 0 { // Cek apakah angka genap
			genap++ // Tambahkan counter genap
		} else {
			ganjil++ // Tambahkan counter ganjil
		}
	}

	// Return jumlah ganjil dan genap
	return ganjil, genap
}

// Soal jawaban nomor 2
// Fungsi CariIndexDariAngkaTertentu mencari index pertama dari angka target dalam slice
func CariIndexDariAngkaTertentu(nums []int, target int) int {
	// Loop semua angka dengan index
	for index, num := range nums {
		if num == target { // Jika angka sama dengan target
			return index // Kembalikan index
		}
	}
	// Jika tidak ditemukan → return -1
	return -1
}

// Soal jawaban nomor 3
// Fungsi CekSlicePalindrome mengecek apakah slice merupakan palindrome
func CekSlicePalindrome(slice []int) bool {
	i := 0              // Pointer awal
	j := len(slice) - 1 // Pointer akhir

	// Loop selama i < j
	for i < j {
		if slice[i] != slice[j] { // Jika elemen awal ≠ akhir
			return false // Bukan palindrome
		}
		i++ // Geser pointer awal
		j-- // Geser pointer akhir
	}
	// Jika loop selesai tanpa beda → palindrome
	return true
}

// Soal jawaban nomor 4
// Fungsi HitungJumlahAngkaLebihKecildariNilaiTertentu menghitung jumlah angka < batas
func HitungJumlahAngkaLebihKecildariNilaiTertentu(nums []int, batas int) int {
	counter := 0 // Inisialisasi counter

	// Loop semua angka
	for _, num := range nums {
		if num < batas { // Jika angka < batas
			counter++ // Tambahkan counter
		}
	}

	// Return jumlah angka < batas
	return counter
}

// Soal jawaban nomor 5
// Fungsi BuatSliceBarudenganAngkaDikali2 membuat slice baru dengan setiap angka dikali 2
func BuatSliceBarudenganAngkaDikali2(nums []int) []int {
	var result []int // Inisialisasi slice kosong baru

	// Loop semua angka
	for _, num := range nums {
		result = append(result, num*2) // Kalikan 2 dan tambahkan ke slice baru
	}

	// Return slice baru
	return result
}
