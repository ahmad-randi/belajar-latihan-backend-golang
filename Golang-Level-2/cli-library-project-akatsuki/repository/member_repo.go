package repository

import (
	"cli-library-project-akatsuki/entity"
	"cli-library-project-akatsuki/errors"
)

// Membuat database sementar menampung di slice
type Repository struct {
	members []entity.Member
}

// Add member baru
func (r *Repository) AddMember(member entity.Member) {
	r.members = append(r.members, member)
}

// FindById mencari member berdasarkan ID
func (r *Repository) FindByID(id int) (*entity.Member, error) {
	for i := range r.members {
		if r.members[i].ID == id {
			return &r.members[i], nil
		}
	}
	return nil, errors.ErrMemberNotFound
}

// Update menganti data member
func (r *Repository) UpdateMember(member entity.Member) error {
	for i := range r.members {
		if r.members[i].ID == member.ID {
			r.members[i] = member
			return nil
		}
	}
	return errors.ErrMemberNotFound
}

// Delete menghapus member berdasarkan ID
func (r *Repository) Delete(id int) error {
	for i := range r.members {
		if r.members[i].ID == id {
			r.members = append(r.members[:i], r.members[i+1:]...)
			return nil
		}
	}
	return errors.ErrMemberNotFound
}
