package repository

import "LifeGuardAlertas/src/alerta/domain/entities"

type AlertaRepository interface {
	Create(data *entities.Alerta) error
}
