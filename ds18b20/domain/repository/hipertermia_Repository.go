package repository

import "LifeGuardAlertas/ds18b20/domain/entities"

type Hipertermia interface {
	FindAll() ([]entities.AlertaDs18b20, error)
	Create(hipertermia *entities.AlertaDs18b20) error
}