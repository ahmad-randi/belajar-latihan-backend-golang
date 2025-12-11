package main

import (
	"Golang-Level-1/latihan"
	"fmt"
)

func main() {
	// latihan satu.go
	fmt.Println("Hasil Nomor 1 :", latihan.GetMax([]int{3, 9, 2, 5, 7}))
	fmt.Println("Hasil Nomor 2 :", latihan.SumOdd([]int{1, 4, 7, 10, 13}))
	fmt.Println("Hasil Nomor 3 :", latihan.CountAboveAvg([]int{2, 4, 6, 8}))
	fmt.Println("Hasil Nomor 4 :", latihan.Reverse([]int{1, 2, 3, 4}))
	fmt.Println("Hasil Nomor 5 :", latihan.Frequency([]int{1, 2, 1, 3, 2, 1}))

	fmt.Println("===========================================================")

	// latihan dua.go
	fmt.Println("Hasil soal nomor 1 :", latihan.TotalBayaknyaAngkaGenap([]int{2, 5, 8, 11, 14, 3}))
	fmt.Println("Hasil soal nomor 2 :", latihan.CariNilaiTerkecil([]int{7, 2, 9, 1, 5}))
	fmt.Println("Hasil soal nomor 3 :", latihan.HitungTotalSemuaAngka([]int{3, 5, 2, 8}))
	fmt.Println("Hasil soal nomor 4 :", latihan.CekAngkaTertentuDalamSlice([]int{4, 9, 1, 3}, 1))
	fmt.Println("Hasil soal nomor 5 :", latihan.HitungRataRataAngka([]int{10, 20, 30}))

	fmt.Println("===========================================================")

	//latihan tiga.go
	fmt.Println("Hasil soal nomor 1 : ", latihan.HitungJumlahSemuaAngkaSlice([]int{4, 2, 7}))
	fmt.Println("Hasil soal nomor 2 : ", latihan.HitungBerapaBanyakAngkaGenap([]int{1, 2, 3, 4, 6}))
	fmt.Println("Hasil soal nomor 3 : ", latihan.AmbilAngkaPertamaDariSlice([]int{9, 4, 7}))
	fmt.Println("Hasil soal nomor 4 : ", latihan.AmbilAngkaTerakhirDariSlice([]int{3, 8, 1, 6}))
	fmt.Println("Hasil soal nomor 4 : ", latihan.CariNilaiTerbesar([]int{5, 8, 2, 10, 4}))

	fmt.Println("===========================================================")

	//latihan empat.go
	fmt.Println(func() string {
		g, e := latihan.HitungJumlahAngkaGenapdanGanjil([]int{2, 5, 8, 11, 14})
		return fmt.Sprintf("Hasil soal nomor 1 : Ganjil: %d, Genap: %d", g, e)
	}())
	fmt.Println("Hasil soal nomor 2 :", latihan.CariIndexDariAngkaTertentu([]int{4, 7, 1, 9}, 1))
	fmt.Println("Hasil soal nomor 3 : ", latihan.CekSlicePalindrome([]int{1, 3, 3, 1}))
	fmt.Println("Hasil soal nomor 4 : ", latihan.HitungJumlahAngkaLebihKecildariNilaiTertentu([]int{2, 5, 8, 1}, 5))
	fmt.Println("Hasil soal nomor 5 : ", latihan.BuatSliceBarudenganAngkaDikali2([]int{1, 3, 5, 7}))

	fmt.Println("===========================================================")

	// latihan lima.go
	fmt.Println("Hasil soal nomor 1 : ", latihan.HitungPanjangKatadalamSlice([]string{"go", "java", "py", "python"}))
	fmt.Println("Hasil soal nomor 2 : ", latihan.AmbilElemenStringPertamayangPanjangnyaGenap([]string{"hi", "go", "java", "py"}))
	fmt.Println("Hasil soal nomor 3 : ", latihan.FilterSlicedenganAngkaLebihBesardari5([]int{2, 7, 4, 9, 1}))
	jumlahAngka, jumlahString := latihan.HitungJumlahAngkadanString([]interface{}{1, "go", 2, "python", 3})
	fmt.Println("Hasil soal nomor 4 :", jumlahAngka, jumlahString)
	fmt.Println("Hasil soal nomor 5 : ", latihan.ReverseSliceCampuran([]interface{}{1, "go", 3, "hi"}))

}
