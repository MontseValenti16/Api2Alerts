package application

import (
	"LifeGuardAlertas/ds18b20/domain/entities"
	"LifeGuardAlertas/ds18b20/domain/repository"
)

type GetAllFiebreUseCase struct {
	fiebreRepo repository.Fiebre
}

func NewGetAllFiebreUseCase(fiebre repository.Fiebre) *GetAllFiebreUseCase {
	return &GetAllFiebreUseCase{
		fiebreRepo: fiebre,
	}
}

func (u *GetAllFiebreUseCase) GetAllFiebre() ([]*entities.AlertaDs18b20, error) {
	return u.fiebreRepo.FindAll()
}