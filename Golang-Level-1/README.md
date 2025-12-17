# 🧩 Golang Level 1 — Fundamental Logika

Level 1 adalah **fondasi utama** dalam proses belajar Backend Golang.

Fokus level ini **bukan framework**, melainkan:

* cara berpikir logis
* memahami alur program
* membiasakan problem solving tanpa helper berlebihan

Jika Level 1 ini kuat, maka naik ke level backend akan jauh lebih mudah.

---

## 🎯 Tujuan Level 1

Setelah menyelesaikan Level 1, diharapkan mampu:

* Menguasai logika dasar Golang
* Berpikir **step by step** sebelum menulis kode
* Mengolah data menggunakan loop, kondisi, dan function
* Menulis kode yang rapi, jelas, dan mudah dibaca

---

## 📂 Struktur Folder Level 1

```
Golang-Level-1/
├── main.go            # Entry point untuk menjalankan latihan
├── go.mod             # Module Golang Level 1
├── latihan/            # Kumpulan soal latihan
│   ├── satu.go        # Max, sum odd, count above avg, reverse, frequency
│   ├── dua.go         # Genap, min, total, cek angka, rata-rata
│   ├── tiga.go        # Total, genap, ambil angka pertama & terakhir, max
│   ├── empat.go       # Genap & ganjil, palindrome, cari index, slice *2
│   ├── lima.go        # Slice campuran (interface{}), filter, reverse
│   ├── enam.go        # Filter & transform data, analisis logika
│   ├── tujuh.go       # Soal bertingkat (mudah → sangat susah)
│   ├── delapan.go     # Analisis string & validasi pola angka
│   ├── sembilan.go    # Evaluasi ulang logika Level 1
│   └── final-level.go # Final Test Level 1
```

---

## 📚 Materi yang Dipelajari

<details>
<summary><strong>🔹 Dasar Bahasa & Struktur</strong></summary>

  * Variabel & tipe data (`int`, `string`, `bool`, dll)
  * Operator dasar (`+ - * / %`)
  * Looping (`for`, `for range`)
  * Conditional (`if`, `else`)
  * Slice & array (`append`, index)
  * Function (parameter & return)
    
</details>

---

<details>
<summary><strong>🔹 Logika Dasar</strong></summary>

  * Sum, max, min
  * Hitung angka genap & ganjil
  * Rata-rata & frequency
  * Reverse slice
  * Ambil elemen pertama & terakhir
  * Cek angka dalam slice
  * Cari index tertentu
  * Palindrome slice
  * Transform data (buat slice baru)

</details>

---

<details>
<summary><strong>🔹 String & Slice Campuran</strong></summary>

  * Olah slice string
  * Filter data berdasarkan kondisi
  * Slice campuran (`[]interface{}`)
  * Hitung jumlah angka & string
  * Reverse slice campuran
  * `type switch`

</details>

---

<details>
<summary><strong>🔹 Latihan Lanjutan (`enam.go`)</strong></summary>

  Fokus pada **alur logika & transformasi data**:
  
  * Filter angka ganjil → transform (perkalian)
  * Ambil angka genap pertama
  * Menentukan nilai maksimum tanpa helper
  * Analisis string per karakter

</details>

---


<details>
<summary><strong>🔹 Soal Bertingkat (`tujuh.go`)</strong></summary>

  Latihan problem solving bertahap:
  
  1. **Mudah** — hitung string dengan panjang > N
  2. **Sedang** — ambil angka genap terakhir
  3. **Susah** — hitung jumlah huruf vokal
  4. **Sangat Susah** — cek urutan angka:
  
     * naik
     * turun
     * tidak beraturan
  
  Menggunakan **flag logic** (`isNaik`, `isTurun`) tanpa sorting.

</details>

---

<details>
<summary><strong>🔹 Analisis String & Pola Angka (`delapan.go`)</strong></summary>

  Latihan ketelitian & logika lanjut:
  
  * Hitung huruf kapital
  * Ambil kata terpanjang
  * Hitung karakter angka
  * Validasi pola **gunung (naik lalu turun)**:
  
    * minimal 3 angka
    * tidak boleh datar
    * tidak boleh naik setelah turun

</details>

---

<details>
<summary><strong>🔹 Evaluasi Logika (`sembilan.go`)</strong></summary>

  Digunakan untuk mengukur kematangan logika:
  
  * Hitung angka genap
  * Hitung string lebih panjang dari N
  * Ambil angka ganjil pertama
  * Cari nilai terbesar
  * Hitung string mengandung vokal

</details>

---

<details>
<summary><strong>🔹 Final Test Level 1 (`final-level.go`)</strong></summary>

  Ujian akhir Level 1 dengan fokus:
  
  * Flag logic
  * Early return
  * Double loop terkontrol
  * Filter → transform → simpan
  * Validasi pola kompleks (lembah)

</details>

Jika seluruh soal bisa dikerjakan & dijelaskan, maka **Level 1 dinyatakan lulus**.

---

## ▶️ Cara Menjalankan

```bash
cd Golang-Level-1
go run main.go
```

---

## ✅ Aturan Belajar (Wajib Diikuti)

* ❌ Tidak langsung lihat jawaban
* ✍️ Tulis alur logika di komentar
* 🧠 Utamakan pemahaman, bukan cepat selesai
* 🔁 Ulangi soal sampai paham

---

🔥 **Catatan Mentor:**
Jika Level 1 ini terasa sulit, itu tandanya logika lu sedang dibentuk. Jangan loncat level sebelum benar-benar paham.
