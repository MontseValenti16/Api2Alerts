package application

import (
	"LifeGuardAlertas/src/samrtwatch/domain/entities"
	"LifeGuardAlertas/src/samrtwatch/domain/repo"
)

type GetSmartwatchUseCase struct {
	smartwatchRepo repo.SmartwatchRepository
}

func NewGetSmartwatchUseCase(smartwatchRepo repo.SmartwatchRepository) *GetSmartwatchUseCase {
	return &GetSmartwatchUseCase{
		smartwatchRepo: smartwatchRepo,
	}
}

func (u *GetSmartwatchUseCase) Run(macAddress string) (*entities.Smartwatch, error) {
	return u.smartwatchRepo.GetByMacAddress(macAddress)
}
