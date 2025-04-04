package entities

type DatoDiario struct {
    ID         int     `json:"id_dato_diario"`
    IDPersona  int     `json:"id_persona"`
    Fecha      string  `json:"fecha"` // o usar time.Time si prefieres
    HoraInicio string  `json:"hora_inicio"`
    HoraFin    string  `json:"hora_fin"`
    Promedio   float64 `json:"promedio"`
}