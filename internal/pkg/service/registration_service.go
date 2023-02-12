package service

import (
	"github.com/sudak-91/wasm-test/pkg/repository"
)

type RegistrationService struct {
	UserRepo repository.Users
}

type DuplicateUserError struct {
	Message string
}

func (e *DuplicateUserError) Error() string {
	return "UserDuplicate"
}
func NewRegistrationService(userRepo repository.Users) *RegistrationService {
	return &RegistrationService{
		UserRepo: userRepo,
	}
}

func (r *RegistrationService) Service(user repository.User) error {
	err := r.duplicateChecker(user.Email)
	if err != nil {
		return err
	}
	user.IsTemporary = true
	err = r.UserRepo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (r *RegistrationService) duplicateChecker(email string) error {
	users, err := r.UserRepo.ReadUserByEmail(email)
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return nil
	}
	return &DuplicateUserError{}
}
