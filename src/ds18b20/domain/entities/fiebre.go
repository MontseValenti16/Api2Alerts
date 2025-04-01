package entities

type AlertaDs18b20 struct {
	ID int `json:"id"`
	Mensaje string `json:"mensaje"`
	Nivel_peligro string `json:"nivel_peligro"`
	Fecha string `json:"fecha"`
}
