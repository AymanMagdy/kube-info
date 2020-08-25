package repository

import "../entity"

type PodRepository interface {
	FindAll() ([]entity.Pod, error)
}