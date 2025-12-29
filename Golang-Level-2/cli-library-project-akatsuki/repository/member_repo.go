package repository

import (
	"cli-library-project-akatsuki/entity"
	"cli-library-project-akatsuki/errors"
)

// Membuat database sementar menampung di slice
type RepositoryMember struct {
	members []entity.Member
}

// Add member baru
func (r *RepositoryMember) AddMember(member entity.Member) {
	r.members = append(r.members, member)
}

/*
FindAll data member
*/
func (r *RepositoryMember) FindAll() []entity.Member {
	return r.members
}

// FindById mencari member berdasarkan ID
func (r *RepositoryMember) FindByID(id int) (*entity.Member, error) {
	for i := range r.members {
		if r.members[i].ID == id {
			return &r.members[i], nil
		}
	}
	return nil, errors.ErrMemberNotFound
}

// Update menganti data member
func (r *RepositoryMember) UpdateMember(member entity.Member) error {
	for i := range r.members {
		if r.members[i].ID == member.ID {
			r.members[i] = member
			return nil
		}
	}
	return errors.ErrMemberNotFound
}

// Delete menghapus member berdasarkan ID
func (r *RepositoryMember) Delete(id int) error {
	for i := range r.members {
		if r.members[i].ID == id {
			r.members = append(r.members[:i], r.members[i+1:]...)
			return nil
		}
	}
	return errors.ErrMemberNotFound
}
