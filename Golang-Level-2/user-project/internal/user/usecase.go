package user

import "user-project/internal/user/dto"

// USECASE CREATE
// Mengatur ALUR BISNIS
func CreateUserUsecase(req dto.CreateUserRequest) error {
	if err := ValidateCreate(req); err != nil {
		return err
	}

	Create(req)
	return nil
}

// USECASE UPDATE
func UpdateUserUsecase(req dto.UpdateUserRequest) error {
	return UpdateUser(req)
}

// USECASE DELETE
func DeleteUserUsecase(id int) error {
	return DeleteUser(id)
}
