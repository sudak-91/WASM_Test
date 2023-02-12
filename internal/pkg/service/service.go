package service

import "github.com/sudak-91/wasm-test/pkg/repository"

type Services interface {
	Registration
}

type Registration interface {
	Service(repository.User) error
}

type BlogService struct {
	repo                repository.Repository
	RegistrationService Registration
}

func NewBlogService(repo repository.Repository) *BlogService {
	return &BlogService{
		repo:                repo,
		RegistrationService: NewRegistrationService(repo),
	}
}
