package services

import (
	"../entity"
	"../repository"
)

type ClusterService interface {
	Validate(cluster *entity.Cluster) error
	FindAll() ([]entity.Cluster, error)
}


type ClusterServiceObj struct{}

var ClusterRepo repository.ClusterRepository

func NewClusterService(repository repository.ClusterRepositiory) ClusterServiceObj {
	ClusterRepo = repository
	return &ClusterRepo{}
}

func (*ClusterServiceObj) Validate(video *entity.Cluster) error {
	if ClusterRepo == nil{
		err := errors.New("No nodes in the cluster..")
		return err
	}
	return nil
}

func (*ClusterServiceObj) FindAll() ([]entity.Cluster, error) {
	return ClusterRepo.FindAll()
}