package repo

import "LifeGuardAlertas/src/samrtwatch/domain/entities"

type SmartwatchRepository interface {
	Create(data *entities.Smartwatch) error
	GetByMacAddress(macAddress string) (*entities.Smartwatch, error)
	UpdateByMacAddress(macAddress string, data *entities.Smartwatch) error
}