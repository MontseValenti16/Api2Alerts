package application

import (
    "LifeGuardAlertas/src/datos/domain/entities"
    "LifeGuardAlertas/src/datos/domain/repository"
)

type GetAllDatosDiariosUseCase struct {
    repo repository.DatosDiariosRepository
}

func NewGetAllDatosDiariosUseCase(repo repository.DatosDiariosRepository) *GetAllDatosDiariosUseCase {
    return &GetAllDatosDiariosUseCase{
        repo: repo,
    }
}

func (uc *GetAllDatosDiariosUseCase) Run() ([]*entities.DatoDiario, error) {
    return uc.repo.GetAll()
}