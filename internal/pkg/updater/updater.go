package updater

import (
	"log"

	pubrep "github.com/sudak-91/wasm-test/pkg/repository"
)

type Updater struct {
	repoUser pubrep.Users
}

type Update struct {
	Type   string       `json:"type"`
	SignIn *pubrep.User `json:"signin,omitempty"`
}

func NewUpdater(repository pubrep.Users) Updater {
	return Updater{
		repoUser: repository,
	}
}

func (u Updater) Controler(update Update) error {
	switch update.Type {
	case "login":
	case "registration":

	default:
		log.Println("Default")

	}
	return nil
}
