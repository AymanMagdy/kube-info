package services

import (
	"../logger"
	"../entity"
	"../repository"
)

type NamespaceService interface {
	Validate(node *entity.Namespace) error
	FindAll() ([]entity.Namespace, error)
}

type NamespaceServiceObj struct{}

var namespaceRepo repository.NamespaceRepository

func FindAll() ([]entity.Namespace, error) {
	return namespaceRepo.FindAll()
}

func NewNamespaceService(repository repository.NamespaceRepository) NamespaceService {
	namespaceRepo = repository
	return &namespaceRepo{}
}

func (*NamespaceServiceObj) Validate(node *entity.Namespace) error {
	if namespaceRepo == nil{
		err := logger.ServiceError(
			"No found namespaces",
			404,
		)
		return err
	}
	return nil
}
