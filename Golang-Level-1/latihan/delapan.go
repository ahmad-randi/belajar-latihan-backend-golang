package latihan

// HitungJumlahHurufKapital menghitung total huruf kapital di setiap string dalam slice
func HitungJumlahHurufKapital(data []string) int {
	counter := 0                   // inisialisasi counter untuk menyimpan jumlah huruf kapital
	for _, isiData := range data { // loop setiap string dalam slice
		for _, ch := range isiData { // loop setiap karakter di string
			if ch >= 'A' && ch <= 'Z' { // cek apakah karakter adalah huruf kapital
				counter++ // jika iya, tambahkan counter
			}
		}
	}
	return counter // kembalikan total huruf kapital
}

// AmbilKataTerpanjang mengambil string terpanjang dari slice
func AmbilKataTerpanjang(data []string) string {
	max := data[0]                   // inisialisasi max dengan string pertama
	for i := 0; i < len(data); i++ { // loop semua string dalam slice
		if len(data[i]) > len(max) { // jika panjang string saat ini lebih dari max
			max = data[i] // update max dengan string baru
		}
	}
	return max // kembalikan string terpanjang
}

// HitungKataYangMengandungAngka menghitung jumlah karakter angka dalam slice string
func HitungKataYangMengandungAngka(data []string) int {
	counter := 0                   // inisialisasi counter untuk menyimpan jumlah angka
	for _, isiData := range data { // loop setiap string dalam slice
		for _, ch := range isiData { // loop setiap karakter di string
			if ch >= '0' && ch <= '9' { // cek apakah karakter adalah angka
				counter++ // jika iya, tambahkan counter
			}
		}
	}
	return counter // kembalikan total angka
}

func ValidasiPolaNaikTurunGunung(data []int) bool {
	// Minimal 3 angka biar bisa naik lalu turun
	if len(data) < 3 {
		return false
	}

	//inisialisasi flag
	naik := false
	turun := false

	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			if turun {
				return false
			}
			naik = true
		} else if data[i] < data[i-1] {
			if !naik {
				return false
			}
			turun = true
		} else {
			return false
		}
	}
	return naik && turun
}
