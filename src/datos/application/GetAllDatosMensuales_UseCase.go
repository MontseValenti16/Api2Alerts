package application

import (
    "LifeGuardAlertas/src/datos/domain/entities"
    "LifeGuardAlertas/src/datos/domain/repository"
)

type GetAllDatosMensualesUseCase struct {
    repo repository.DatosMensualRepository
}

func NewGetAllDatosMensualesUseCase(repo repository.DatosMensualRepository) *GetAllDatosMensualesUseCase {
    return &GetAllDatosMensualesUseCase{
        repo: repo,
    }
}

func (uc *GetAllDatosMensualesUseCase) Run() ([]*entities.DatoMensual, error) {
    return uc.repo.GetAll()
}