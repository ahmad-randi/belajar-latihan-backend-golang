# 👤 User Project — Backend Engineering CLI (Industry Style)

Project ini merupakan **simulasi backend user management** berbasis **Command Line Interface (CLI)** yang dirancang menggunakan **arsitektur backend ala industri**.

Walaupun project ini **belum menggunakan database, API, atau framework**, seluruh struktur dan alur logikanya **merepresentasikan backend production-ready**.

📌 Project ini berfungsi sebagai **blueprint / referensi utama** untuk membangun project backend lain pada Level 2 dan level berikutnya.

---

## 🎯 Project Objectives

Tujuan utama project ini adalah melatih:

* mindset backend engineer
* desain arsitektur backend yang rapi dan scalable
* pemisahan tanggung jawab antar layer
* alur data yang jelas dari input hingga penyimpanan

Setelah memahami project ini, kamu diharapkan:

* tidak bingung membaca struktur project backend besar
* paham peran setiap layer
* siap mengembangkan backend berbasis API dan database

---

## 📂 Project Structure

```
user-project/
├── main.go                     # Application entry point
│
├── cmd/
│   └── cli/
│       └── menu.go             # CLI interaction & user input
│
├── internal/
│   └── user/
│       ├── dto/                # Data Transfer Object (input layer)
│       │   ├── create_user.go
│       │   └── update_user.go
│       │
│       ├── model.go            # Domain model (User entity)
│       ├── validator.go        # User input validation rules
│       ├── helper.go           # Reusable helper (FindByID, GenerateID)
│       ├── repository.go       # Data storage (in-memory slice)
│       ├── service.go          # Technical logic
│       ├── usecase.go          # Business flow orchestration
│       └── errors.go           # Domain-specific errors
│
├── pkg/
│   ├── logger/
│   │   └── logger.go           # Centralized logging utility
│   └── response/
│       └── response.go         # Standardized CLI response
```

---

## 🔄 Application Flow

```
CLI (menu.go)
      ↓
DTO (Create / Update User)
      ↓
Usecase (Business Flow)
      ↓
Validator (Input Validation)
      ↓
Service (Technical Logic)
      ↓
Repository (In-Memory Storage)
```

Struktur alur ini **identik dengan backend REST API**, hanya berbeda pada media input (CLI).

---

## 🔍 Why FindByID Exists?

`FindByID` digunakan untuk:

* mencari user secara terpusat dan konsisten
* menghindari duplikasi loop di banyak layer
* memastikan update dan delete hanya terjadi pada data yang valid
* menjaga service dan usecase tetap bersih

📌 Ini adalah **best practice backend engineering**, bukan sekadar helper tambahan.

---

## 🧠 Why DTO Is Used?

DTO (Data Transfer Object) berfungsi untuk:

* memisahkan **data input** dari **domain model**
* menjaga entity `User` tetap bersih
* memudahkan validasi
* mempersiapkan project untuk REST API

Contoh:

* `CreateUserDTO` → data saat create user
* `UpdateUserDTO` → data saat update user

---

## ▶️ How to Run

Masuk ke folder `user-project`, lalu jalankan:

```bash
go run main.go
```

---

## ✅ Learning Checklist

Gunakan checklist ini sebagai indikator kesiapan:

* [ ] Paham fungsi setiap folder
* [ ] Paham alur CRUD user
* [ ] Bisa menjelaskan perbedaan usecase, service, dan repository
* [ ] Bisa menambah field baru pada User
* [ ] Bisa menggunakan struktur ini untuk project lain

---

## 🚀 Future Improvements

Project ini **siap dikembangkan lebih lanjut**, antara lain:

* migrasi dari slice ke database
* perubahan CLI menjadi REST API
* penambahan pagination & search
* penambahan authentication & authorization

---

## 🚦 Mentor Notes

Jika kamu sudah mampu **menjelaskan alur project ini tanpa melihat kode**, maka:

🔥 kamu sudah memiliki **fondasi backend engineering yang kuat**.

Project ini bukan sekadar latihan, melainkan **kerangka berpikir backend profesional**.
