package updater

import (
	"log"

	pubrep "github.com/sudak-91/wasm-test/pkg/repository"
	pubupdater "github.com/sudak-91/wasm-test/pkg/updater"
)

type Updater struct {
	repo pubrep.Repository
}

func NewUpdater(repository pubrep.Repository) Updater {
	return Updater{
		repo: repository,
	}
}

func (u Updater) Controler(update pubupdater.Update) error {
	switch update.Type {
	case "login":
		return nil
	case "registration":
		err := u.registrationUpdater(update.Data)
		if err != nil {
			return err
		}
		return nil
	default:
		log.Println("Default")
		return nil

	}
}
