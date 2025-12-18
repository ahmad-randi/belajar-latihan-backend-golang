package user

import "user-project/internal/user/dto"

// Logic teknis CREATE
func Create(req dto.CreateUserRequest) {
	user := User{
		ID:     GenerateID(users),
		Nama:   req.Nama,
		Umur:   req.Umur,
		Alamat: req.Alamat,
	}
	Save(user)
}

// Logic teknis UPDATE
func UpdateUser(req dto.UpdateUserRequest) error {
	_, index := FindByID(users, req.ID)
	if index == -1 {
		return ErrUserNotFound
	}

	Update(index, User{
		ID:     req.ID,
		Nama:   req.Nama,
		Umur:   req.Umur,
		Alamat: req.Alamat,
	})
	return nil
}

// Logic teknis DELETE
func DeleteUser(id int) error {
	_, index := FindByID(users, id)
	if index == -1 {
		return ErrUserNotFound
	}

	Delete(index)
	return nil
}
