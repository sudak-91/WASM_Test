package updater

import (
	"fmt"

	pubrep "github.com/sudak-91/wasm-test/pkg/repository"
	pubupdater "github.com/sudak-91/wasm-test/pkg/updater"
)

func (u *Updater) registrationUpdater(data any) error {
	user, ok := data.(pubrep.User)
	if !ok {
		return pubupdater.NewInvalidData()
	}
	userList, err := u.repo.ReadUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if len(userList) != 0 {
		return DuplicateUser{}
	}
	err = u.repo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

type DuplicateUser struct {
}

func (d DuplicateUser) Error() string {
	return fmt.Sprint("User has duplicate")
}
