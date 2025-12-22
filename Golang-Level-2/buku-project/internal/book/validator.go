package book

import (
	"buku-project/internal/book/dto"
	"errors"
	"strings"
)

/*
Validator bertugas:
- memastikan INPUT USER benar
- mencegah BUG LOGIKA
*/

func ValidateCreateBook(req dto.CreateBookRequest) error {
	if strings.TrimSpace(req.Title) == "" {
		return errors.New("title wajib diisi")
	}
	if strings.TrimSpace(req.Author) == "" {
		return errors.New("author wajib diisi")
	}
	if len(req.Description) < 20 {
		return errors.New("description minimal 20 karakter")
	}

	// Cegah duplikat buku
	for _, b := range repo.FindAll() {
		if strings.EqualFold(b.Title, req.Title) &&
			strings.EqualFold(b.Author, req.Author) {
			return errors.New("buku sudah terdaftar")
		}
	}
	return nil
}

func ValidateUpdateBook(req dto.UpdateBookRequest) error {
	if req.ID <= 0 {
		return ErrInvalidBookInput
	}
	if strings.TrimSpace(req.Title) == "" {
		return errors.New("title wajib diisi")
	}
	return nil
}
