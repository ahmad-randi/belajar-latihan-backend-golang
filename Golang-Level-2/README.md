# 🚀 Golang Level 2 — Fundamental Backend (CLI Based)

Level 2 adalah **jembatan penting** antara *logika dasar (Level 1)* dan *backend production-ready*.
Pada level ini, fokus utama adalah **cara berpikir backend engineer**, bukan framework.

Semua studi kasus dibuat dalam bentuk **CLI (Command Line Interface)** agar kamu:

* memahami alur data secara eksplisit
* terbiasa menulis kode terstruktur & aman
* tidak bergantung pada magic framework

---

## 🎯 Tujuan Level 2

Setelah menyelesaikan Level 2, kamu diharapkan mampu:

* Mendesain **struct & layer sederhana** (model, service)
* Mengelola data menggunakan **slice sebagai in-memory database**
* Menerapkan **validasi data & secure coding dasar**
* Membuat **CRUD lengkap (Create, Read, Update, Delete)**
* Menulis program **CLI interaktif** dengan alur jelas
* Memahami *kenapa backend code ditulis seperti ini*

Level ini adalah **pondasi sebelum masuk ke:**

> goroutines, context, database, REST API, framework (Fiber/Gin)

---

## 📂 Struktur Folder Level 2

```
Golang-Level-2/
├── Pengenalan/                # Materi dasar Level 2
│   ├── model/                 # Struct data (User, Product, dll)
│   │   └── user.go
│   ├── service/               # Logic & validasi
│   │   └── user_service.go
│   └── main.go                # Entry point latihan
│
├── Projects/                  # Studi kasus & mini project
│   ├── CRUD-Users/            # CRUD Users via CLI
│   │   ├── model/
│   │   ├── service/
│   │   └── main.go
│   │
│   ├── Kasir-CLI/              # Mini project kasir
│   │   ├── model/
│   │   ├── service/
│   │   └── main.go
│   │
│   └── Inventory-CLI/          # Mini project inventory barang
│       ├── model/
│       ├── service/
│       └── main.go
│
├── go.mod                      # Module Golang Level 2
└── README.md                   # Dokumentasi Level 2
```

---

## 📚 Materi yang Dipelajari

<details>
<summary><strong>🔹 Pengenalan Struct & Slice</strong></summary>

 * Membuat struct yang rapi & konsisten (`User`, `Product`)
 * Pemisahan tanggung jawab (`model` vs `service`)
 * Slice sebagai **database sementara (in-memory)**
 * Konsep *fail fast* & *early return*
 * Validasi data dasar
 * Multiple error menggunakan `[]error`

</details>

<details>
<summary><strong>🔹 CRUD CLI (Tanpa Database)</strong></summary>

 * **Create** → menambahkan data ke slice
 * **Read** → menampilkan data (by ID / all)
 * **Update** → update data dengan validasi
 * **Delete** → hapus data menggunakan slice baru
 * Menu CLI interaktif (`fmt.Scanln`)
 
 > Fokus: alur data & logika backend, bukan UI

</details>

<details>
<summary><strong>🔹 Mini Project / Studi Kasus</strong></summary>

 * **CRUD Users CLI**
 
   * Simulasi backend user management
   * Validasi input & error handling
 
 * **Kasir CLI (Belum Tuntas)**
 
   * Tambah produk
   * Hitung total belanja
   * Checkout
 
 * **Inventory CLI (Belum Tuntas)**
 
   * Tambah barang
   * Update stok
   * Hapus barang
 
 Semua project menggunakan **pola logika yang sama** seperti backend sungguhan.

</details>

---

## ▶️ Cara Menjalankan

### 1️⃣ Pengenalan

```bash
cd Golang-Level-2/Pengenalan
go run main.go
```

### 2️⃣ CRUD Users

```bash
cd Golang-Level-2/Projects/CRUD-Users
go run main.go
```

### 3️⃣ Kasir CLI

```bash
cd Golang-Level-2/Projects/Kasir-CLI
go run main.go
```

### 4️⃣ Inventory CLI

```bash
cd Golang-Level-2/Projects/Inventory-CLI
go run main.go
```

---

## ✅ Aturan Belajar (Wajib)

* ❌ Jangan langsung copy jawaban
* ✍️ Tulis alur logika di komentar
* 🧠 Pahami *kenapa* kode dibuat seperti itu
* 🔁 Ulangi sampai bisa jelasin tanpa lihat kode

---

## 🚦 Catatan Mentor

Jika Level 2 ini sudah terasa **masuk akal & rapi**, berarti:

* mindset backend kamu **sudah terbentuk**
* kamu siap naik ke **Level 3 (Concurrency, DB, API)**

> Backend bukan soal framework cepat, tapi **alur & konsistensi logika**.

🔥 Selamat datang di Level 2 — mulai belajar sebagai backend engineer.
