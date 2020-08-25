package repository

import "../entity"

type NodeRepository interface {
	FindAll() ([]entity.Node, error)
}