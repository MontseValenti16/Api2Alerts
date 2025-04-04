package application

import (
    "LifeGuardAlertas/src/datos/domain/entities"
    "LifeGuardAlertas/src/datos/domain/repository"
)

type GetAllDatosSemanalUseCase struct {
    repo repository.DatosSemanalesRepository
}

func NewGetAllDatosSemanalUseCase(repo repository.DatosSemanalesRepository) *GetAllDatosSemanalUseCase {
    return &GetAllDatosSemanalUseCase{
        repo: repo,
    }
}

func (uc *GetAllDatosSemanalUseCase) Run() ([]*entities.DatoSemanal, error) {
    return uc.repo.GetAll()
}