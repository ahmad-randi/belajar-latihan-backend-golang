# 🩸 Akatsuki CLI Library Management (In Progress)

CLI berbasis **Golang** dengan tema **Akatsuki Organization**.
Project ini bertujuan melatih **logika backend**, **alur menu CLI**, dan **arsitektur clean** tanpa database dan tanpa REST API.

> Status: 🚧 **On Progress (In-Memory / CLI Only)**

---

## 🎯 Tujuan Project

* Melatih pola pikir backend melalui CLI
* Memahami alur data tanpa database
* Mengimplementasikan konsep:

  * Entity
  * Repository (In-Memory)
  * Service
  * Menu & Routing
* Membiasakan clean code & struktur project

---

## 🧩 Fitur Utama

### 👤 Member Management

* Add Member
* View All Member
* Update Member
* Deactivate Member
* Delete Member

### 📜 Mission Management

* Create Mission
* Assign Mission ke Member
* Update Status Mission
* View All Missions

### 📊 Organization Stats

* Total Member Aktif
* Total Mission Aktif
* Total Mission Selesai
* Top Member berdasarkan Reward

---

## 🗂️ Contoh Data (In-Memory)

### 📌 Member Data

| ID | Nama    | Partner | Rank      | Status | Total Misi | Total Reward |
| -- | ------- | ------- | --------- | ------ | ---------- | ------------ |
| 1  | Itachi  | Kisame  | Executive | Aktif  | 1          | 50000        |
| 2  | Kisame  | Itachi  | Member    | Aktif  | 0          | 0            |
| 3  | Deidara | Sasori  | Member    | Aktif  | 0          | 0            |

---

### 📌 Mission Data

| ID | Nama Misi            | Level | Target      | Status  | Assigned | Reward |
| -- | -------------------- | ----- | ----------- | ------- | -------- | ------ |
| 1  | Capture Tailed Beast | S     | Jinchuriki  | Selesai | Itachi   | 50000  |
| 2  | Assassination        | A     | Kage Escort | Aktif   | -        | 30000  |

---

## ▶️ Preview Tampilan CLI

### 🩸 Main Menu

```
========================================
      🩸 AKATSUKI ORGANIZATION SYSTEM 🩸
========================================
"Rasa sakit adalah guru dunia"

1. Member Management
2. Mission Management
3. Organization Stats
0. Exit

Pilih menu:
```

---

### 👤 Menu Member

```
===== Menu Member Akatsuki =====
1. Add Member
2. View All Member
3. Update Member
4. Deactivate Member
5. Hapus Data Member
0. Back to Main Menu
```

---

### 📜 Menu Mission

```
===== Menu Mission Akatsuki =====
1. Create Mission
2. Assign Mission
3. Update Mission Status
4. View All Missions
0. Back to Main Menu
```

---

### 📊 Organization Stats

```
===== Organization Statistics =====
Total Member Aktif  : 3
Total Misi Aktif    : 1
Total Misi Selesai  : 1
Top Member          : Itachi (50000 Ryō)
```

---

## 🧠 Arsitektur Project (Sederhana)

```
cli-library-project-akatsuki/
├── cmd/
│   └── run.go
├── menu/
│   ├── member_menu.go
│   ├── mission_menu.go
│   └── stats_menu.go
├── entity/
│   ├── member.go
│   └── mission.go
├── repository/
│   ├── member_repository.go
│   └── mission_repository.go
├── service/
│   ├── member_service.go
│   └── mission_service.go
├── util/
│   ├── input.go
│   └── banner.go
└── main.go
```

---

## 🚀 Cara Menjalankan

```bash
go run main.go
```

---

## 📌 Catatan

* Semua data disimpan **di memory (slice)**
* Data akan **hilang saat program ditutup**
* Fokus utama: **alur logic & clean architecture**

---

## 👤 Author

**Bedul**
Backend Golang Learner 🧠🔥

---

> "Perdamaian hanyalah ilusi" — Akatsuki
