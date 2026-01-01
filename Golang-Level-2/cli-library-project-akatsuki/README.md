# 🩸 Akatsuki CLI Library Management (In Progress)

CLI berbasis **Golang** dengan tema **Akatsuki Organization**.  
Project ini bertujuan melatih **logika backend**, **alur menu CLI**, dan **arsitektur clean** tanpa database dan tanpa REST API.

> Status: 🚧 **On Progress (In-Memory / CLI Only)**

---

## 🎯 Tujuan Project

- Melatih pola pikir backend melalui CLI
- Memahami alur data tanpa database
- Mengimplementasikan konsep:
  - Entity
  - Repository (In-Memory)
  - Service
  - Menu & Routing
- Membiasakan clean code & struktur project

---

## 🧩 Fitur Utama

### 👤 Member Management
- Add Member
- View All Member
- Update Member
- Deactivate Member
- Delete Member

### 📜 Mission Management
- Create Mission
- Assign Mission ke Member
- Update Status Mission
- View All Missions

### 📊 Organization Stats
- Total Member Aktif
- Total Mission Aktif
- Total Mission Selesai
- Top Member berdasarkan Reward

---

## 🗂️ Contoh Data (In-Memory)

### 📌 Member Data

| ID | Nama    | Partner | Rank      | Status | Total Misi | Total Reward |
|----|---------|---------|-----------|--------|------------|--------------|
| 1  | Itachi  | Kisame  | Executive | Aktif  | 1          | 50000        |
| 2  | Kisame  | Itachi  | Member    | Aktif  | 0          | 0            |
| 3  | Deidara | Sasori  | Member    | Aktif  | 0          | 0            |

---

### 📌 Mission Data

| ID | Nama Misi            | Level | Target      | Status  | Assigned | Reward |
|----|----------------------|-------|-------------|---------|----------|--------|
| 1  | Capture Tailed Beast | S     | Jinchuriki  | Selesai | Itachi   | 50000  |
| 2  | Assassination        | A     | Kage Escort | Aktif   | -        | 30000  |

---

## ▶️ Preview Tampilan CLI (FULL)

### 🩸 Program Start

```
========================================
      🩸 AKATSUKI ORGANIZATION SYSTEM 🩸
========================================
"Rasa sakit adalah guru dunia"

1. Member Management
2. Mission Management
3. Organization Stats
0. Exit

Pilih menu: _
```

---

## 👤 MEMBER MANAGEMENT

```
===== Menu Member Akatsuki =====
1. Add Member
2. View All Member
3. Update Member
4. Deactivate Member
5. Hapus Data Member
0. Back to Main Menu

Pilih menu member: _
```

### ➕ Add Member

```
Masukkan Nama Member   : Itachi
Masukkan Partner       : Kisame
Masukkan Rank [Pain/Executive/Member]: Member

✅ Member berhasil ditambahkan!
```

---

### 📄 View All Member

```
ID  Nama     Partner   Rank        Status   TotalMisi  TotalReward
1   Itachi   Kisame    Member      Aktif    0          0
2   Kisame   Itachi    Executive   Aktif    0          0
```

---

### ✏️ Update Member

```
Masukkan ID Member : 1
Masukkan Rank Baru [Pain/Executive/Member]: Executive

✅ Data member berhasil diupdate
```

---

### 💀 Deactivate Member

```
Masukkan ID Member : 2

💀 Member Kisame telah dinonaktifkan
```

---

## 📜 MISSION MANAGEMENT

```
===== Menu Mission Akatsuki =====
1. Create Mission
2. Assign Mission
3. Update Mission Status
4. View All Missions
0. Back to Main Menu

Pilih menu mission: _
```

### 🆕 Create Mission

```
Nama Misi   : Capture Tailed Beast
Level       : S
Target      : Jinchuriki
Reward      : 50000

✅ Misi berhasil dibuat
```

---

### 🔗 Assign Mission

```
Masukkan ID Misi   : 1
Masukkan ID Member : 1

✅ Misi berhasil diberikan ke Itachi
```

---

### ✅ Update Mission Status

```
Masukkan ID Misi : 1
Status Baru [Aktif/Selesai/Gagal]: Selesai

✅ Misi selesai!
Reward 50000 Ryō ditambahkan ke Itachi
```

---

### 📄 View All Missions

```
ID  Nama Misi                Level  Target        Status    Assigned  Reward
1   Capture Tailed Beast     S      Jinchuriki    Selesai   Itachi    50000
```

---

## 📊 ORGANIZATION STATS

```
===== Organization Statistics =====
Total Member Aktif        : 5
Total Misi Aktif          : 2
Total Misi Selesai        : 7
Top Member (Reward)       : Itachi (50000 Ryō)
```

---

## ❌ EXIT

```
"Perdamaian hanyalah ilusi..."
Program CLI ditutup, terima kasih by Bedul 👋
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

- Semua data disimpan **di memory (slice)**
- Data akan **hilang saat program ditutup**
- Menu back menggunakan `return` (bukan recursive)
- `bufio.Reader` dibuat **1 kali**, dipassing ke semua menu
- Fokus utama: **alur logic & clean architecture**

---

## 👤 Author

**Bedul**  
Backend Golang Learner 🧠🔥

---

> _"Perdamaian hanyalah ilusi..." — Akatsuki_
