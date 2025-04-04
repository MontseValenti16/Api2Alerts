package entities

type DatoAnual struct {
    ID         int     `json:"id_dato_anual"`
    IDPersona  int     `json:"id_persona"`
    Anio       int     `json:"anio"`
    Promediopasos  float64 `json:"promediopasos"`
    Promediobpm   float64 `json:"promediobpm"`
    Promediospo2   float64 `json:"promediospo2"`
    Promediotemperatura   float64 `json:"promediotemperatura"`
}