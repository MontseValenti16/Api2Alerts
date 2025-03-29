package application

import (
	"LifeGuardAlertas/ds18b20/domain/entities"
	"LifeGuardAlertas/ds18b20/domain/repository"
)

type CreateHipertermiaUseCase struct {
	hipertermiaRepo repository.Hipertermia
}

func NewCreateHipertermiaUseCase(hipertermiaRepo repository.Hipertermia) *CreateHipertermiaUseCase {
	return &CreateHipertermiaUseCase{
		hipertermiaRepo: hipertermiaRepo,
	}
}

func (u *CreateHipertermiaUseCase) Run(data *entities.AlertaDs18b20) error {
	return u.hipertermiaRepo.Create(data)
}