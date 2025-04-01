package application

import (
	"LifeGuardAlertas/ds18b20/domain/entities"
	"LifeGuardAlertas/ds18b20/domain/repository"
)

type GetAllHipertermiaUseCase struct {
	hipertermiaRepo repository.Hipertermia
}

func NewGetAllHipertermiaUseCase(hipertermia repository.Hipertermia) *GetAllHipertermiaUseCase {
	return &GetAllHipertermiaUseCase{
		hipertermiaRepo: hipertermia,
	}
}

func (u *GetAllHipertermiaUseCase) Run() ([]*entities.AlertaDs18b20, error) {	
	return u.hipertermiaRepo.FindAll()
}	