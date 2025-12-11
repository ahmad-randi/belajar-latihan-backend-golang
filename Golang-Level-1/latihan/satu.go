package latihan

// Soal jawaban nomor 1
// Fungsi GetMax mencari nilai maksimum di slice nums
func GetMax(nums []int) int {
	// Anggap elemen pertama sebagai max sementara
	max := nums[0]

	// Loop semua elemen di slice
	for _, num := range nums {
		// Jika elemen saat ini lebih besar dari max, update max
		if num > max {
			max = num
		}
	}
	// Kembalikan nilai max
	return max
}

// Soal jawaban nomor 2
// Fungsi SumOdd menjumlahkan semua angka ganjil di slice nums
func SumOdd(nums []int) int {
	// Inisialisasi total sum
	sum := 0

	// Loop semua elemen di slice
	for _, num := range nums {
		// Cek apakah angka ganjil
		if num%2 != 0 {
			// Jika ganjil, tambahkan ke sum
			sum += num
		}
	}
	// Kembalikan total sum angka ganjil
	return sum
}

// Soal jawaban nomor 3
// Fungsi CountAboveAvg menghitung jumlah elemen yang lebih besar dari rata-rata
func CountAboveAvg(nums []int) int {
	// Hitung total semua elemen
	total := 0
	for _, num := range nums {
		total += num
	}

	// Hitung rata-rata
	rataRata := float64(total) / float64(len(nums))

	// Hitung jumlah elemen di atas rata-rata
	count := 0
	for _, num := range nums {
		if float64(num) > rataRata {
			count++
		}
	}
	// Kembalikan jumlah elemen di atas rata-rata
	return count
}

// Soal jawaban nomor 4
// Fungsi Reverse membalik urutan elemen di slice nums
func Reverse(nums []int) []int {
	// Inisialisasi pointer i di awal dan j di akhir slice
	i := 0
	j := len(nums) - 1

	// Loop selama i < j
	for i < j {
		// Tukar elemen nums[i] dengan nums[j]
		nums[i], nums[j] = nums[j], nums[i]
		// Geser pointer i ke kanan, j ke kiri
		i++
		j--
	}
	// Kembalikan slice yang sudah dibalik
	return nums
}

// Soal jawaban nomor 5
// Fungsi Frequency menghitung frekuensi setiap angka di slice nums
func Frequency(nums []int) map[int]int {
	// Inisialisasi map kosong untuk menyimpan frekuensi
	freq := map[int]int{}

	// Loop semua elemen di slice
	for _, num := range nums {
		// Tambahkan count di map, jika belum ada otomatis jadi 0 + 1
		freq[num]++
	}
	// Kembalikan map frekuensi
	return freq
}
