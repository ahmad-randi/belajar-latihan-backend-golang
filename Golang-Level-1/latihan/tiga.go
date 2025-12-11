package latihan

// Soal jawaban nomor 1
// Fungsi HitungJumlahSemuaAngkaSlice menjumlahkan semua angka di slice nums
func HitungJumlahSemuaAngkaSlice(nums []int) int {
	sum := 0                   // Inisialisasi total sum
	for _, num := range nums { // Loop semua elemen di slice
		sum += num // Tambahkan elemen ke sum
	}
	return sum // Kembalikan total sum
}

// Soal jawaban nomor 2
// Fungsi HitungBerapaBanyakAngkaGenap menghitung jumlah angka genap di slice nums
func HitungBerapaBanyakAngkaGenap(nums []int) int {
	count := 0                 // Inisialisasi counter
	for _, num := range nums { // Loop semua elemen
		if num%2 == 0 { // Cek apakah angka genap
			count++ // Jika genap, tambahkan counter
		}
	}
	return count // Kembalikan jumlah angka genap
}

// Soal jawaban nomor 3
// Fungsi AmbilAngkaPertamaDariSlice mengambil elemen pertama dari slice nums
func AmbilAngkaPertamaDariSlice(nums []int) int {
	if len(nums) == 0 { // Cek apakah slice kosong
		return 0 // Jika kosong, return 0
	}
	return nums[0] // Return elemen pertama
}

// Soal jawaban nomor 4
// Fungsi AmbilAngkaTerakhirDariSlice mengambil elemen terakhir dari slice nums
func AmbilAngkaTerakhirDariSlice(nums []int) int {
	if len(nums) > 0 { // Cek apakah slice tidak kosong
		indexTerakhir := len(nums) - 1 // Hitung index terakhir
		return nums[indexTerakhir]     // Return elemen terakhir
	}
	return 0 // Jika slice kosong, return 0
}

// Soal jawaban nomor 5
// Fungsi CariNilaiTerbesar mencari nilai terbesar di slice nums
func CariNilaiTerbesar(nums []int) int {
	max := nums[0]                   // Anggap elemen pertama sebagai max sementara
	for i := 1; i < len(nums); i++ { // Loop dari index kedua sampai terakhir
		if nums[i] > max { // Bandingkan elemen saat ini dengan max
			max = nums[i] // Update max jika lebih besar
		}
	}
	return max // Kembalikan nilai max
}
