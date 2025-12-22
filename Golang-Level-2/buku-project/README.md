# 📚 Book Project (CLI)

Book Project adalah aplikasi **CRUD buku berbasis CLI (Command Line Interface)** yang dibuat menggunakan **Golang**.

Project ini dirancang sebagai **latihan backend fundamental** dengan fokus pada:

* Struktur kode yang rapi
* Pemisahan tanggung jawab (layered architecture)
* Business logic yang aman dari bug logika
* Kode yang mudah dipahami oleh pemula

---

## 🎯 Tujuan Project

Project ini dibuat untuk melatih:

* Cara berpikir backend developer
* Alur data dari input → validasi → proses → penyimpanan
* Penulisan kode yang **jelas, aman, dan scalable**

> ❗ Project ini **TIDAK menggunakan database** dan **TIDAK menggunakan framework**.
> Semua data disimpan sementara di memory (slice).

---

## 🧩 Fitur

* ✅ Tambah buku
* ✅ Lihat semua buku
* ✅ Lihat detail buku
* ✅ Ubah data buku
* ✅ Pinjam buku (dengan validasi status)
* ✅ Kembalikan buku
* ✅ Hapus buku

Business rule utama:

* Buku hanya bisa dipinjam **jika tersedia**
* Buku hanya bisa dikembalikan **jika sedang dipinjam**
* Buku tidak boleh duplikat (title + author)

---

## 🗂️ Struktur Folder

```
book-project/
├── main.go
├── cmd/
│   └── cli/
│       └── menu.go
│
├── internal/
│   └── book/
│       ├── book_model.go    # Model / Entity
│       ├── dto/             # Data Transfer Object (input user)
│       │   ├── create_book_request.go
│       │   └── update_book_request.go
│       ├── repository.go    # Penyimpanan data (in-memory)
│       ├── service.go       # Business logic
│       ├── validator.go     # Validasi input & aturan bisnis
│       ├── helper.go        # Helper (Generate ID)
│       └── errors.go        # Custom error
```

---

## 🧠 Penjelasan Arsitektur

### 1️⃣ `main.go`

Entry point aplikasi.

Tugas:

* Menjalankan CLI
* Tidak berisi logika apapun

---

### 2️⃣ `menu.go` (CLI Layer)

Tugas:

* Menampilkan menu
* Menerima input user
* Mengarahkan ke service

❌ Tidak boleh ada logika bisnis di sini

---

### 3️⃣ `service.go` (Business Logic Layer)

Tugas:

* Menjalankan aturan bisnis
* Mengatur alur proses aplikasi

Contoh aturan:

* Buku tidak boleh dipinjam dua kali
* Buku harus ada sebelum dihapus

---

### 4️⃣ `repository.go` (Data Layer)

Tugas:

* Menyimpan data
* Mengambil data

❌ Tidak melakukan validasi
❌ Tidak tahu aturan bisnis

---

### 5️⃣ `validator.go`

Tugas:

* Memastikan input user valid
* Mencegah bug logika sejak awal

---

## 🔐 Keamanan & Anti Bug Logika

Project ini aman dari bug logika umum seperti:

* Buku dipinjam berkali-kali
* Buku dikembalikan padahal belum dipinjam
* ID tidak valid
* Data duplikat

> Walaupun CLI, pola ini **siap dipakai untuk REST API**

---

## ▶️ Cara Menjalankan

Pastikan Go sudah terinstall.

```bash
go run main.go
```

---

## 📈 Roadmap Pengembangan

Jika project ini dikembangkan lebih lanjut:

* 🔜 Logging aktivitas
* 🔜 Riwayat peminjaman
* 🔜 Multi-user simulation
* 🔜 Migrasi ke REST API

---

## 👨‍💻 Catatan Developer

Project ini dibuat sebagai **latihan backend level fundamental–menengah**.

Fokus utama:

* Bukan banyak fitur
* Tapi **alur logika yang benar**

---

## ⭐ Penutup

Jika kamu pemula di backend Golang, project ini cocok sebagai:

* Template belajar
* Bahan portofolio
* Dasar sebelum masuk ke REST API

Happy coding 🚀
