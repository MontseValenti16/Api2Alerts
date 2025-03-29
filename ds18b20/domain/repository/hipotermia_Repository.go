package repository

import "LifeGuardAlertas/ds18b20/domain/entities"

type Hipotermia interface {
	FindAll() ([]entities.AlertaDs18b20, error)	
	Create(hipotermia *entities.AlertaDs18b20) error
}