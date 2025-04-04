package application

import (
	"LifeGuardAlertas/src/samrtwatch/domain/entities"
	"LifeGuardAlertas/src/samrtwatch/domain/repo"
)

type UpdateSmartwatchUseCase struct {
	smartwatchRepo repo.SmartwatchRepository
}

func NewUpdateSmartwatchUseCase(smartwatchRepo repo.SmartwatchRepository) *UpdateSmartwatchUseCase {
	return &UpdateSmartwatchUseCase{
		smartwatchRepo: smartwatchRepo,
	}
}

func (u *UpdateSmartwatchUseCase) Run(macAddress string, data *entities.Smartwatch) error {
	return u.smartwatchRepo.UpdateByMacAddress(macAddress, data)
}