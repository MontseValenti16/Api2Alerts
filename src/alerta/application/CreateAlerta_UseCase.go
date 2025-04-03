package application

import (
	"LifeGuardAlertas/src/alerta/domain/entities"
	"LifeGuardAlertas/src/alerta/domain/repository"
)

type CreateAlertaUseCase struct {
	alertaRepo repository.AlertaRepository
}

func NewCreateAlertaUseCase(alertaRepo repository.AlertaRepository) *CreateAlertaUseCase {
	return &CreateAlertaUseCase{
		alertaRepo: alertaRepo,
	}
}

func (u *CreateAlertaUseCase) Run(data *entities.Alerta) error {
	return u.alertaRepo.Create(data)
}