package repository

import "LifeGuardAlertas/src/datos/domain/entities"

type DatosDiariosRepository interface {
    GetAll() ([]*entities.DatoDiario, error)
}