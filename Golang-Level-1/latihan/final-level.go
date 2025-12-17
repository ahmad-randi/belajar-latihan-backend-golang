package latihan

func FlagLogic(data []int) bool {
	//Cek apakah data kosong
	if len(data) == 0 {
		//Jika kosong false
		return false
	}

	//Membuat flag
	isGenap := false
	isGanjil := false

	//Loop isi data
	for _, value := range data {
		//Cek apakah data genap semua
		if value%2 == 0 {
			//Cek apakah value melanggar
			if isGanjil {
				//Jika melanggar maka false
				return false // melanggar aturan konsisten
			}
			//Jika ya maka true
			isGenap = true
		} else {
			if isGenap {
				//Jika melanggar maka false
				return false // melanggar aturan konsisten
			}
			isGanjil = true
		}
	}
	//return tidak ada pelanggaran jadi true
	return true
}

func EarlyReturn(data []string) string {
	//Loop semua isi data
	for _, value := range data {
		//Loop isi data perkarakter
		for _, ch := range value {
			//Cek apakah value mengandung kapital A-Z
			if ch >= 'A' && ch <= 'Z' { // cek apakah karakter adalah huruf kapital
				//Jika ada return Value
				return value
			}
		}
	}
	//Return default
	return ""
}

func HitungStringDenganMinimalDuaVokal(data []string) int {
	// Jika slice kosong, tidak ada yang bisa dihitung
	if len(data) == 0 {
		return 0
	}

	// Counter utama:
	// Menyimpan jumlah string yang memenuhi syarat (>= 2 vokal)
	totalValid := 0

	// Loop LUAR:
	// Mengambil string satu per satu dari slice
	for _, kata := range data {

		// Counter LOKAL untuk satu string
		// PENTING: di-reset setiap masuk string baru
		jumlahVokal := 0

		// Loop DALAM:
		// Mengecek setiap karakter dalam string
		for _, ch := range kata {

			// Cek apakah karakter adalah huruf vokal
			if ch == 'a' || ch == 'i' || ch == 'u' || ch == 'e' || ch == 'o' {
				// Jika vokal, tambah counter vokal string ini
				jumlahVokal++
			}
		}

		// Setelah loop karakter SELESAI,
		// cek apakah string ini valid (>= 2 vokal)
		if jumlahVokal >= 2 {
			// Jika valid, tambahkan counter utama
			totalValid++
		}
	}

	// Kembalikan total string yang valid
	return totalValid
}
