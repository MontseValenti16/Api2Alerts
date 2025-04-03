package entities

type Alerta struct {
	ID 				int    `json:"id"`
	MacAddress 		string `json:"macaddress"`
	Mensaje 		string `json:"mensaje"`
	Accion			string `json:"accion"`
	Nivel_peligro 	string `json:"nivel_peligro"`
	Fecha 			string `json:"fecha"`
}
