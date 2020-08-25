package services

import (
	"../entity"
	"../repository"
)

type PodService interface {
	Validate(pod *entity.Pod) error
	FindAll() ([]entity.Pod, error)
}

type PodServiceObj struct{}

var PodRepo repository.PodRepository

func NewPodService(repository repository.PodRepository) PodService {
	repo = repository
	return &service{}
}

func (*PodServiceObj) Validate(pods *entity.Pod) error {
	if pods == nil{
		err := errors.New("No pods in the namespaec..")
		return err
	}
	return nil
}

func (*PodServiceObj) FindAll() ([]entity.Pod, error) {
	return repo.FindAll()
}