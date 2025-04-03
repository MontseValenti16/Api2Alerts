package application

import (
	"LifeGuardAlertas/src/alerta/domain/entities"
	"LifeGuardAlertas/src/alerta/domain/repository"
)

type GetAlertasUseCase struct {
	alertaRepo repository.AlertaRepository
}

func NewGetAlertasUseCase(alertaRepo repository.AlertaRepository) *GetAlertasUseCase {
	return &GetAlertasUseCase{
		alertaRepo: alertaRepo,
	}
}

func (u *GetAlertasUseCase) Run() ([]entities.Alerta, error) {
	return u.alertaRepo.GetAll()
}

func (u *GetAlertasUseCase) RunByMacAddress(macAddress string) ([]entities.Alerta, error) {
	return u.alertaRepo.GetByMacAddress(macAddress)
}