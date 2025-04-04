package repository

import "LifeGuardAlertas/src/datos/domain/entities"

type DatosMensualRepository interface {
    GetAll() ([]*entities.DatoMensual, error)
}