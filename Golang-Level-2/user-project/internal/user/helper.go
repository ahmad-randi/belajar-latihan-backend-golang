package user

// GenerateID membuat ID otomatis (in-memory)
func GenerateID(users []User) int {
	return len(users) + 1
}

// FindByID mencari user dan index slice
// Dipakai untuk UPDATE & DELETE
func FindByID(users []User, id int) (*User, int) {
	for i, u := range users {
		if u.ID == id {
			return &u, i
		}
	}
	return nil, -1
}
