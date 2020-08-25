package repository

import "../entity"

type ClusterRepository interface {
	FindAll() ([]entity.Cluster, error)
}