package mysql

import (
	"LifeGuardAlertas/src/alerta/domain/entities"
	"LifeGuardAlertas/src/core/db"
	"fmt"
	"log"
)

type MySQL struct {
	conn *db.ConnMySQL
}

func NewMySQL() *MySQL {
	conn := db.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al conectar a la base de datos: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) Create(data *entities.Alerta) error {
	query := "INSERT INTO alertas (macaddress, mensaje, accion, nivel_peligro, fecha) VALUES (?, ?, ?, ?, ?)"
	_, err := mysql.conn.DB.Exec(query, data.MacAddress, data.Mensaje, data.Accion, data.Nivel_peligro, data.Fecha)
	if err != nil {
		return fmt.Errorf("error al guardar el dato en la base de datos: %w", err)
	}
	return nil
}
