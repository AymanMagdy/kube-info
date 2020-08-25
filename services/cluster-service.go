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

var ClusterRepo repository.ClusterRepositiory

func NewClusterService(repository repository.ClusterRepositiory) ClusterService {
	ClusterRepo = repository
	return &service{}
}

func (*ClusterServiceObj) Validate(video *entity.Cluster) error {
	if cluster == nil{
		err := errors.New("No nodes in the cluster..")
		return err
	}
	return nil
}

func (*ClusterServiceObj) FindAll() ([]entity.Cluster, error) {
	return ClusterRepo.FindAll()
}