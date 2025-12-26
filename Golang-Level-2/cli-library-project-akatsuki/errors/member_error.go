package errors

import "errors"

var (
	ErrMemberNotFound  = errors.New("member tidak ditemukan")
	ErrMemberInactive  = errors.New("member sudah tidak aktif")
	ErrInvalidInput    = errors.New("input tidak valid")
	ErrDuplicateMember = errors.New("member sudah terdaftar")
)
