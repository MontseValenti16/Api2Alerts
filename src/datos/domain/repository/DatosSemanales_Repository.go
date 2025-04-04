package repository

import "LifeGuardAlertas/src/datos/domain/entities"

type DatosSemanalesRepository interface {
    GetAll() ([]*entities.DatoSemanal, error)
}
