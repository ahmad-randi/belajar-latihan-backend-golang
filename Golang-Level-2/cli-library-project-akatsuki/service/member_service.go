package service

import (
	"cli-library-project-akatsuki/dto/member_dto"
	"cli-library-project-akatsuki/entity"
	"cli-library-project-akatsuki/errors"
	"cli-library-project-akatsuki/helper"
	"cli-library-project-akatsuki/repository"
	"cli-library-project-akatsuki/validate"
	"strings"
	"time"
)

var repo = repository.RepositoryMember{}

// Add new member logic
func AddMemberService(req member_dto.CreateMemberRequestUser) error {
	if err := validate.ValidateAddMember(req); err != nil {
		return err
	}

	//Cek duplikate add member
	for _, value := range repo.FindAll() {
		if strings.EqualFold(value.Name, req.Nama) &&
			strings.EqualFold(value.Partner, req.Patner) {
			return errors.ErrDuplicateMember
		}
	}

	now := time.Now()
	member := entity.Member{
		ID:        helper.GenerateMemberID(),
		Name:      req.Nama,
		Rank:      req.Rank,
		Partner:   req.Patner,
		Status:    false,
		CreatedAt: now,
		UpdatedAt: now,
	}
	repo.AddMember(member)
	return nil
}

// View all member service logic
func GetAllMemberService() []entity.Member {
	return repo.FindAll()
}

// Update data memver sevice logic
func UpdateMmeberService(req member_dto.UpdateMemberRequestUser) error {
	if err := validate.ValidateUpdateMember(req); err != nil {
		return err
	}

	member, err := repo.FindByID(req.ID)
	if err != nil {
		return err
	}

	if member.Status {
		return errors.ErrMemberInactive
	}

	member.Name = req.Nama
	member.Partner = req.Patner
	member.Rank = req.Rank
	member.UpdatedAt = time.Now()

	return repo.UpdateMember(*member)
}
