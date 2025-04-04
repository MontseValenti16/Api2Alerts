package entities

import "time"

type DatoSemanal struct {
    ID          int       `json:"id_dato_semanal"`
    IDPersona   int       `json:"id_persona"`
    FechaInicio time.Time `json:"fecha_inicio"`
    FechaFin    time.Time `json:"fecha_fin"`
    Promediopasos  float64 `json:"promediopasos"`
    Promediobpm   float64 `json:"promediobpm"`
    Promediospo2   float64 `json:"promediospo2"`
    Promediotemperatura   float64 `json:"promediotemperatura"`
}