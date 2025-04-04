package application

import (
"LifeGuardAlertas/src/samrtwatch/domain/repo"
"LifeGuardAlertas/src/samrtwatch/domain/entities"
)

type CreateSmartwatchUseCase struct {
	smartwatchRepo repo.SmartwatchRepository
}

func NewCreateSmartwatchUseCase(smartwatchRepo repo.SmartwatchRepository) *CreateSmartwatchUseCase {
	return &CreateSmartwatchUseCase{
		smartwatchRepo: smartwatchRepo,
	}
}

func (u *CreateSmartwatchUseCase) Run(data *entities.Smartwatch) error {
	return u.smartwatchRepo.Create(data)
}