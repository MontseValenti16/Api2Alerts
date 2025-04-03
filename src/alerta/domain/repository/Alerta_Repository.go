package repository

import "LifeGuardAlertas/src/alerta/domain/entities"

type AlertaRepository interface {
	Create(data *entities.Alerta) error
	GetAll() ([]entities.Alerta, error)
	GetByMacAddress(macAddress string) ([]entities.Alerta, error)
}
