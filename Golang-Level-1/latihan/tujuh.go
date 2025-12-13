package latihan

//Function Hitung String Lebih Panjang dari N
func HitungStringLebihPanjangdariN(data []string, n int) int {
	//Membuat variabel counter
	counter := 0
	//Loop semua isi slice string
	for _, isiData := range data {
		//Jika panjang string > n, maka counter++
		if len(isiData) > n {
			counter++
		}
	}
	//Kembalikan nilai counter
	return counter
}

//Function Ambil Angka Terakhir yang Genap
func AmbilAngkaTerakhiryangGenap(data []int) int {
	//Loop dari belakang
	for i := len(data) - 1; i >= 0; i-- {
		//Cek apakah angka genap
		if data[i]%2 == 0 {
			//jika ya, retrun langsung angka tersebut
			return data[i]
		}
	}
	//Jika selesai loop, return -1
	return -1
}

func HitungTotalHurufVokal(data []string) int {
	//Membuat variabel counter
	counter := 0
	//Loop isi data slice
	for _, isiData := range data {
		//Loop perkarakter dari isi data
		for _, ch := range isiData {
			//Cek apakah ada karakter a,i,u,e,o pada isiData
			if ch == 'a' || ch == 'i' || ch == 'u' || ch == 'e' || ch == 'o' {
				//jika ya, maka counter++
				counter++
			}
		}
	}
	//kembalikan nilai counter
	return counter
}

func ValidasiPolaSliceAngka(data []int) bool {
	// 1️⃣ Validasi awal: jika slice kosong atau hanya 1 elemen → dianggap valid
	if len(data) <= 1 {
		return true
	}

	// 2️⃣ Inisialisasi flag
	// isNaik  → true jika slice monoton naik
	// isTurun → true jika slice monoton turun
	isNaik := true
	isTurun := true

	// 3️⃣ Loop dari index 1 sampai akhir slice
	//    Bandingkan setiap elemen dengan elemen sebelumnya
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			// a. jika elemen sekarang > elemen sebelumnya
			//    → berarti slice tidak monoton turun
			isTurun = false
		} else if data[i] < data[i-1] {
			// b. jika elemen sekarang < elemen sebelumnya
			//    → berarti slice tidak monoton naik
			isNaik = false
		}
		// c. jika sama (data[i] == data[i-1])
		//    → tidak mengubah flag, tetap dianggap valid
	}

	// 4️⃣ Return true jika slice monoton naik atau turun
	return isNaik || isTurun
}
