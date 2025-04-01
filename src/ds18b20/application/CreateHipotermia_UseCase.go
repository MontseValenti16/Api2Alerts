package application

import (
	"LifeGuardAlertas/ds18b20/domain/entities"
	"LifeGuardAlertas/ds18b20/domain/repository"
)

type CreateHipotermiaUseCase struct {
	hipotermiaRepo repository.Hipotermia
}

func NewCreateHipotermiaUseCase(hipotermiaRepo repository.Hipotermia) *CreateHipotermiaUseCase {
	return &CreateHipotermiaUseCase{
		hipotermiaRepo: hipotermiaRepo,
	}
}

func (u *CreateHipotermiaUseCase) Run(data *entities.AlertaDs18b20) error {
	return u.hipotermiaRepo.Create(data)
}