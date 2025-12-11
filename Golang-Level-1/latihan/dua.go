package latihan

// Function: Menghitung jumlah angka genap dalam slice
func TotalBayaknyaAngkaGenap(nums []int) int {

	// 1. Siapkan counter untuk menghitung berapa angka genap
	count := 0

	// 2. Loop semua angka dalam slice nums
	for _, num := range nums {

		// 3. Cek apakah angka genap (sisa bagi 2 = 0)
		if num%2 == 0 {

			// 4. Jika genap → tambahkan counter
			count++
		}
	}

	// 5. Kembalikan total angka genap
	return count
}

// Function: Mencari nilai paling kecil dalam sebuah slice
func CariNilaiTerkecil(nums []int) int {

	// 1. Ambil angka pertama sebagai nilai awal minimum
	min := nums[0]

	// 2. Loop semua angka dalam slice
	for _, num := range nums {

		// 3. Jika angka lebih kecil dari current min → update min
		if num < min {
			min = num
		}
	}

	// 4. Return nilai terkecil
	return min
}

// Function: Menjumlahkan seluruh angka dalam slice
func HitungTotalSemuaAngka(nums []int) int {

	// 1. Siapkan variabel untuk total
	sum := 0

	// 2. Loop semua angka
	for _, num := range nums {

		// 3. Tambahkan angka ke total sum
		sum += num
	}

	// 4. Return total
	return sum
}

// Function: Mengecek apakah target ada dalam slice
func CekAngkaTertentuDalamSlice(nums []int, target int) bool {

	// 1. Loop tiap angka dalam slice
	for _, num := range nums {

		// 2. Jika angka sama dengan target → langsung return true
		if num == target {
			return true
		}
	}

	// 3. Jika loop selesai tanpa ada yang cocok → false
	return false
}

// Function: Menghitung rata-rata angka
func HitungRataRataAngka(nums []int) int {

	// 1. Hitung total semua angka
	total := 0
	for _, num := range nums {
		total += num
	}

	// 2. Rumus rata-rata: total / jumlah elemen
	rata_rata := total / len(nums)

	// 3. Return nilai rata-rata
	return rata_rata
}
