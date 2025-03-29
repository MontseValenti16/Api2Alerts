package repository

import "LifeGuardAlertas/ds18b20/domain/entities"

type Fiebre interface {
	FindAll() ([]entities.AlertaDs18b20, error)
	Create(alertasDs18b20 *entities.AlertaDs18b20) error
}