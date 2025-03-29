package application

import (
	"LifeGuardAlertas/ds18b20/domain/entities"
	"LifeGuardAlertas/ds18b20/domain/repository"
)

type CreateFiebreUseCase struct {
	fiebreRepo repository.Fiebre
}

func NewCreateFiebreUseCase(fiebreRepo repository.Fiebre) *CreateFiebreUseCase {
	return &CreateFiebreUseCase{
		fiebreRepo: fiebreRepo,
	}
}

func (u *CreateFiebreUseCase) Run(data *entities.AlertaDs18b20) error {
	return u.fiebreRepo.Create(data)
}