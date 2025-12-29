package service

import (
	"cli-library-project-akatsuki/dto/member_dto"
	"cli-library-project-akatsuki/entity"
	"cli-library-project-akatsuki/helper"
	"cli-library-project-akatsuki/repository"
	"cli-library-project-akatsuki/validate"
	"errors"
	"strings"
	"time"
)

var repo = repository.RepositoryMember{}

// Logik Add new member
func AddMember(req member_dto.CreateMemberRequestUser) error {
	if err := validate.ValidateAddMember(req); err != nil {
		return err
	}

	//Cek duplikate add member
	for _, value := range repo.FindAll() {
		if strings.EqualFold(value.Name, req.Nama) &&
			strings.EqualFold(value.Partner, req.Patner) {
			return errors.New("member sudah terdaftar")
		}
	}

	now := time.Now()
	member := entity.Member{
		ID:        helper.GenerateMemberID(),
		Name:      req.Nama,
		Rank:      req.Rank,
		Status:    false,
		CreatedAt: now,
		UpdatedAt: now,
	}
	repo.AddMember(member)
	return nil
}
