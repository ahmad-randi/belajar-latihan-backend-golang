package validate

import (
	"cli-library-project-akatsuki/dto/member_dto"
	"errors"
	"strings"
)

/*
	Validasi input user member
	- memastikan input user benar
	- mencegah bug logika
*/

// Validasi add new member
func ValidateAddMember(req member_dto.CreateMemberRequestUser) error {
	if strings.TrimSpace(req.Nama) == "" {
		return errors.New("nama tidak boleh kosong")
	}
	if strings.TrimSpace(req.Patner) == "" {
		return errors.New("patner tidak boleh kosong")
	}
	if strings.TrimSpace(req.Rank) == "" {
		return errors.New("patner tidak boleh kosong")
	}
	return nil

}

// Validasi update member
// func ValidateUpdateMember(rqq member_dto.UpdateMemberRequestUser) error {

// }
