package application

import (
	"LifeGuardAlertas/ds18b20/domain/entities"
	"LifeGuardAlertas/ds18b20/domain/repository"
)

type GetAllHipotermiaUseCase struct {
	hipotermiaRepo *repository.Hipotermia
}

func NewGetAllHipotermiaUseCase(hipotermia repository.Hipotermia) *GetAllHipotermiaUseCase {
	return &GetAllHipotermiaUseCase{
		hipotermiaRepo: hipotermia,
	}
}

func (u *GetAllHipotermiaUseCase) Run() ([]*entities.AlertaDs18b20, error) {
	return u.hipotermiaRepo.FindAll()
}