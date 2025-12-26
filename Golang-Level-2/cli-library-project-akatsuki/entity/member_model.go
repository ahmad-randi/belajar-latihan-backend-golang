package entity

import "time"

// Member merepresentasikan anggota Akatsuki
// Data ini akan disimpan di memory (slice) selama program berjalan
type Member struct {
	ID            int       // ID unik member (auto increment)
	Name          string    // Nama member (Itachi, Kisame, dll)
	Partner       string    // Partner satu tim
	Rank          string    // Pemimpin / Member
	Status        string    // Aktif / Mati / Kabur
	TotalMissions int       // Total misi yang pernah dijalankan
	TotalReward   int       // Total reward (Ryō) dari misi
	CreatedAt     time.Time // Waktu dibuat member
	UpdatedAt     time.Time // Waktu diubah member
}
