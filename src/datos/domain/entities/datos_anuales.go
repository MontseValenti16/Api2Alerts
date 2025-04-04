package entities

type DatoAnual struct {
    ID         int     `json:"id_dato_anual"`
    IDPersona  int     `json:"id_persona"`
    Anio       int     `json:"anio"`
    Promedio   float64 `json:"promedio"`
}