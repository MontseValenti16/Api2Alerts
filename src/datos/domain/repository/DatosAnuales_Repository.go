package repository

import "LifeGuardAlertas/src/datos/domain/entities"

type DatosAnualesRepository interface {
    GetAll() ([]*entities.DatoAnual, error)
}