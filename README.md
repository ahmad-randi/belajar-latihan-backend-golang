# Belajar & Latihan Backend Golang

Repository ini berisi latihan Golang untuk membangun **fundamental logika**, serta **pola berpikir backend dasar**.
Materi disusun **per level** dan bertahap, dimulai dari logika sederhana hingga problem solving yang lebih kompleks sebagai persiapan menuju **Programmer Backend Golang**.

---

## 📂 Struktur Repository

```
belajar-latihan-backend-golang/
├── Golang-Level-1/                # Latihan Level 1
│   ├── main.go                    # File utama untuk menjalankan latihan Level 1
│   ├── go.mod                     # Module Golang-Level-1
│   └── latihan/                   # Folder berisi soal latihan Level 1
│       ├── satu.go                # Max, sum odd, count above avg, reverse, frequency
│       ├── dua.go                 # Genap, min, total, cek angka, rata-rata
│       ├── tiga.go                # Total, genap, ambil angka pertama & terakhir, max
│       ├── empat.go               # Genap & ganjil, cari index, palindrome, angka < batas, slice *2
│       ├── lima.go                # String & int campuran (interface{}), filter, reverse
│       ├── enam.go                # Filter & transform data, logika genap, string, slice campuran
│       ├── tujuh.go               # Soal bertingkat: mudah, sedang, susah, sangat susah
│       ├── delapan.go             # Analisis string lanjutan & validasi pola angka
│       └── sembilan.go            # Evaluasi logika dasar (mudah–sedang–susah)
├── Golang-Level-2/                # Latihan Level 2 (belum ada)
└── README.md                      # Dokumentasi project & roadmap belajar
```

---

## 📚 Materi Golang Level 1

<details>
<summary><strong>🔹 Dasar Bahasa & Struktur</strong></summary>

* Variabel & tipe data dasar (`int`, `float64`, `string`, `bool`)
* Operator dasar (`+`, `-`, `*`, `/`, `%`, `>`, `<`)
* Looping (`for`, `for range`)
* Kondisi (`if`, `else`)
* Slice & array (akses index, `append`)
* Function (parameter & return value)

</details>

<details>
<summary><strong>🔹 Logika Dasar</strong></summary>

* Sum, max, min, count
* Rata-rata & frequency
* Reverse slice
* Ambil angka pertama & terakhir
* Cek angka dalam slice
* Hitung angka genap & ganjil
* Cari index angka tertentu
* Cek slice palindrome
* Hitung angka lebih kecil dari batas
* Membuat slice baru dengan transformasi data

</details>

<details>
<summary><strong>🔹 String & Slice Campuran</strong></summary>

* Hitung panjang kata dalam slice string
* Ambil elemen string pertama dengan panjang genap
* Filter slice angka berdasarkan kondisi
* Mengolah slice campuran (`[]interface{}`)
* Menghitung jumlah angka & string
* Reverse slice campuran

</details>

<details>
<summary><strong>🔹 Latihan Lanjutan (enam.go)</strong></summary>

Fokus pada **alur logika dan transformasi data**:

* Filter angka ganjil lalu transform (perkalian)
* Mengambil angka genap pertama
* Menentukan nilai maksimum tanpa helper bawaan
* Analisis string per karakter
* Penggunaan `type switch` pada slice campuran

</details>

<details>
<summary><strong>🔹 Soal Bertingkat (tujuh.go)</strong></summary>

Latihan problem solving bertahap:

1. **Soal Mudah**
   Menghitung jumlah string dengan panjang lebih dari N

2. **Soal Sedang**
   Mengambil angka genap terakhir dari slice

3. **Soal Susah**
   Menghitung jumlah huruf vokal dari kumpulan string

4. **Soal Sangat Susah**
   Mengecek apakah slice angka:

   * terurut naik
   * terurut turun
   * atau tidak beraturan

   Menggunakan flag logika (`isNaik`, `isTurun`) tanpa sorting.

</details>

<details>
   <summary><strong>🔹 Analisis String & Pola Angka (delapan.go)</strong></summary>
   
   Latihan fokus **ketelitian & logika tingkat lanjut**:
   
   * Menghitung total huruf kapital dalam slice string
   * Mengambil kata terpanjang dari kumpulan string
   * Menghitung karakter angka dalam string
   * Validasi pola angka **naik lalu turun (gunung)**:
     * minimal 3 angka
     * tidak boleh datar
     * tidak boleh naik setelah turun

</details>

<details>
   <summary><strong>🔹 Evaluasi Logika (sembilan.go)</strong></summary>
   
   Latihan evaluasi ulang level 1:
   
   * Hitung jumlah angka genap
   * Hitung jumlah string lebih panjang dari N
   * Ambil angka ganjil pertama
   * Cari nilai terbesar di slice
   * Hitung jumlah string yang mengandung huruf vokal
   
   Digunakan untuk mengukur **kematangan logika dasar**.

</details>

---

## 💡 Cara Menjalankan Level 1

1. Masuk ke folder Level 1:

```bash
cd Golang-Level-1
```

2. Jalankan program:

```bash
go run main.go
```

---

## 🎯 Tujuan Latihan Level 1

* Menguasai logika dasar Golang
* Terbiasa berpikir **step by step**
* Memahami alur logika tanpa bergantung pada helper bawaan
* Membangun pondasi kuat sebelum masuk Level 2
* Menulis kode yang rapi, modular, dan mudah dibaca

---

## 🚀 Roadmap Belajar Backend Golang

| Level | Fokus                     | Skill Output                                   |
| ----- | ------------------------- | ---------------------------------------------- |
| 1     | Dasar logika & syntax     | Loop, slice, function, palindrome, interface{} |
| 2     | Logika menengah           | Nested loop, map, string, sorting, filter      |
| 3     | Problem solving           | Recursion, struct, modular function, testing   |
| 4     | Backend dasar             | HTTP server, JSON, CRUD memory                 |
| 5     | Junior backend siap kerja | REST API, database, auth, unit test, Docker    |

---

## 📝 Contact

**Ahmad Randi**
Instagram: [@ahmadrandy_06](https://instagram.com/ahmadrandy_06)
