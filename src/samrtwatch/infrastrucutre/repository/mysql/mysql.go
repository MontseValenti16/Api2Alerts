package mysql

import (
	"LifeGuardAlertas/src/core/db"
	"LifeGuardAlertas/src/samrtwatch/domain/entities"
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

func (mysql *MySQL) Create(data *entities.Smartwatch) error {
	query := "INSERT INTO smartwatch (macaddress, id_persona) VALUES (?, ?)"
	result, err := mysql.conn.DB.Exec(query, data.MACADDRESS, data.IDPERSONA)
	if err != nil {
		return fmt.Errorf("error al guardar el dato en la base de datos: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error al obtener ID insertado: %w", err)
	}
	data.ID = int(id)

	return nil
}

func (mysql *MySQL) GetByMacAddress(macAddress string) (*entities.Smartwatch, error) {
	query := "SELECT id_persona FROM smartwatch WHERE macaddress = ?"
	row := mysql.conn.DB.QueryRow(query, macAddress)

	var smartwatch entities.Smartwatch
	err := row.Scan(&smartwatch.IDPERSONA)
	if err != nil {
		return nil, fmt.Errorf("error al obtener smartwatch: %w", err)
	}

	return &smartwatch, nil
}

func (mysql *MySQL) UpdateByMacAddress(macAddress string, data *entities.Smartwatch) error {
	query := "UPDATE smartwatch SET id_persona = ? WHERE macaddress = ?"
	_, err := mysql.conn.DB.Exec(query, data.IDPERSONA, macAddress)
	if err != nil {
		return fmt.Errorf("error al actualizar smartwatch: %w", err)
	}

	return nil
}
