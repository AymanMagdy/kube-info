package services

import (
	"../entity"
	"../repository"
)

type NodeService interface {
	Validate(node *entity.Node) error
	FindAll() ([]entity.Node, error)
}

type NodeServiceObj struct{}

var NodeServiceRepo repository.NodeRepository

func NewNodeService(repository repository.NodeRepository) NodeService {
	NodeServiceRepo = repository
	return &service{}
}

func (*NodeServiceObj) Validate(node *entity.Node) error {
	if node == nil{
		err := errors.New("No nodes in the cluster..")
		return err
	}
	return nil
}

func (*NodeServiceObj) FindAll() ([]entity.Node, error) {
	return NodeServiceRepo.FindAll()
}