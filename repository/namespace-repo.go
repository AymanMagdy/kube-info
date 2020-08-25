package repository

import "../entity"

type NamespaceRepository interface {
	FindAll() ([]entity.Namespace, error)
}