package user

import (
	"errors"
	"strings"
	"user-project/internal/user/dto"
)

// Validasi data CREATE user
func ValidateCreate(req dto.CreateUserRequest) error {
	if strings.TrimSpace(req.Nama) == "" {
		return errors.New("nama wajib diisi")
	}
	if req.Umur < 17 {
		return errors.New("umur minimal 17")
	}
	if strings.TrimSpace(req.Alamat) == "" {
		return errors.New("alamat wajib diisi")
	}
	return nil
}
