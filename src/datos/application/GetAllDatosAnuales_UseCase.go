package application

import (
    "LifeGuardAlertas/src/datos/domain/entities"
    "LifeGuardAlertas/src/datos/domain/repository"
)

type GetAllDatosAnualesUseCase struct {
    repo repository.DatosAnualesRepository
}

func NewGetAllDatosAnualesUseCase(repo repository.DatosAnualesRepository) *GetAllDatosAnualesUseCase {
    return &GetAllDatosAnualesUseCase{
        repo: repo,
    }
}

func (uc *GetAllDatosAnualesUseCase) Run() ([]*entities.DatoAnual, error) {
    return uc.repo.GetAll()
}