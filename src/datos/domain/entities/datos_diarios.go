package entities

type DatoDiario struct {
    ID         int     `json:"id_dato_diario"`
    IDPersona  int     `json:"id_persona"`
    Fecha      string  `json:"fecha"` // o usar time.Time si prefieres
    HoraInicio string  `json:"hora_inicio"`
    HoraFin    string  `json:"hora_fin"`
    Promediopasos  float64 `json:"promediopasos"`
    Promediobpm   float64 `json:"promediobpm"`
    Promediospo2   float64 `json:"promediospo2"`
    Promediotemperatura   float64 `json:"promediotemperatura"`
}