package entities

import "time"

type DatoMensual struct {
    ID          int       `json:"id_dato_mensual"`
    IDPersona   int       `json:"id_persona"`
    FechaInicio time.Time `json:"fecha_inicio"`
    FechaFin    time.Time `json:"fecha_fin"`
    Promedio    float64   `json:"promedio"`
}