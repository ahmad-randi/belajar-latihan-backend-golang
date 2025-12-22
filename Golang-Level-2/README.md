# 🚀 Golang Level 2 — Fundamental Backend Engineering (CLI)

Level 2 adalah tahap **transisi penting** dari *belajar syntax Golang (Level 1)* menuju **cara berpikir sebagai Backend Engineer**.

Pada level ini kamu **belum fokus ke API, database, atau framework**, melainkan ke:

* struktur project backend yang rapi
* alur data yang jelas dan terkontrol
* pemisahan tanggung jawab kode (layering)
* logika bisnis yang aman dan mudah dikembangkan

Semua studi kasus dibuat dalam bentuk **CLI (Command Line Interface)** agar alur logika terlihat **jelas & eksplisit**, tanpa bantuan framework.

---

## 🎯 Tujuan Level 2

Setelah menyelesaikan Level 2, kamu diharapkan mampu:

* Mendesain **struktur folder backend yang scalable**
* Memahami **alur data backend dari input sampai penyimpanan**
* Menerapkan **validasi & error handling**
* Membuat **CRUD lengkap (Create, Read, Update, Delete)**
* Menulis kode yang **mudah dibaca, diuji, dan dikembangkan**

Level ini adalah **pondasi sebelum masuk ke materi lanjutan** seperti:

> Goroutines • Context • Database • REST API • Framework (Gin / Fiber)

---

## 📂 Struktur Umum Golang-Level-2

```
Golang-Level-2/
├── Pengenalan/                 # Materi dasar Level 2
│
├── user-project/               # ⭐ Project utama (industry style)
├── buku-project/               # ⭐ Project kedua (learn industry style)
├── kasir-project/              # (menyusul)
├── inventory-project/          # (menyusul)
│
└── README.md                   # Dokumentasi Level 2
```

---

## ⭐ Project Utama: user-project

`user-project` adalah **project inti Level 2**.

Walaupun:

* ❌ belum pakai database
* ❌ belum pakai API
* ❌ belum pakai framework

Namun project ini:

* menggunakan **struktur backend industri**
* menerapkan **layering yang benar**
* siap di-upgrade ke REST API & database

📌 **Project ini menjadi template backend kamu ke depan.**

---

## 📚 Fokus Materi Level 2

### 🔹 Backend Structure & Layering

* `cmd` → input user (CLI)
* `internal` → logic inti aplikasi
* `usecase` → alur bisnis
* `service` → logic teknis
* `repository` → penyimpanan data
* `dto` → data masuk (request)
* `helper` → fungsi reusable (FindByID, GenerateID)

### 🔹 CRUD & Data Flow

* Create
* Read
* Update
* Delete
* Validasi data
* Error handling
* Helper function

### 🔹 Backend Mindset

* kode tidak saling bergantung
* logic mudah ditelusuri
* struktur siap berkembang
* konsisten dan profesional

---

## ▶️ Cara Menjalankan Project

Masuk ke folder project, lalu jalankan:

```bash
go run main.go
```

Contoh:

```bash
cd Golang-Level-2/user-project
go run main.go
```

---

## ✅ Aturan Belajar (WAJIB)

* ❌ Jangan hanya copy–paste
* ✍️ Baca komentar di setiap file
* 🧠 Pahami fungsi tiap folder
* 🔁 Coba modifikasi sendiri

---

## 🚦 Catatan Mentor

Jika kamu sudah **paham struktur & alur user-project**, berarti:

* mindset backend kamu **sudah terbentuk**
* kamu siap naik ke **Level 3**

🔥 Level 2 bukan soal cepat, tapi **kuat & matang**.
