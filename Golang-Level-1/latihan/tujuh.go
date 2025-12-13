package latihan

//Function Hitung String Lebih Panjang dari N
func SoalMudah(data []string, n int) int {
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
func SoalSedang(data []int) int {
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

func SoalSusah(data []string) int {
	counter := 0
	for _, isiData := range data {
		for _, ch := range isiData {

			if ch == 'a' || ch == 'i' || ch == 'u' || ch == 'e' || ch == 'o' {
				counter++
			}
		}
	}
	return counter
}

func SoalSangatSusah(data []int) bool {
	if len(data) <= 1 {
		return true
	}

	isNaik := true
	isTurun := true

	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			isTurun = false
		} else if data[i] < data[i-1] {
			isNaik = false
		}
	}
	return isNaik || isTurun
}
