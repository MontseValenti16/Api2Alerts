// infrastructure/repository/mysql/mysql.go
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
	
	// Enviar a RabbitMQ después de guardar en DB
	err = SendToRabbitMQ(data)
	if err != nil {
		log.Printf("Error enviando a RabbitMQ: %v", err)
	}
	
	return nil
}

func (mysql *MySQL) GetAll() ([]entities.Alerta, error) {
	query := "SELECT id, macaddress, mensaje, accion, nivel_peligro, fecha FROM alertas ORDER BY fecha DESC"
	rows, err := mysql.conn.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener alertas: %w", err)
	}
	defer rows.Close()

	var alertas []entities.Alerta
	for rows.Next() {
		var a entities.Alerta
		err := rows.Scan(&a.ID, &a.MacAddress, &a.Mensaje, &a.Accion, &a.Nivel_peligro, &a.Fecha)
		if err != nil {
			return nil, fmt.Errorf("error al escanear alerta: %w", err)
		}
		alertas = append(alertas, a)
	}

	return alertas, nil
}

func (mysql *MySQL) GetByMacAddress(macAddress string) ([]entities.Alerta, error) {
	query := "SELECT id, macaddress, mensaje, accion, nivel_peligro, fecha FROM alertas WHERE macaddress = ? ORDER BY fecha DESC"
	rows, err := mysql.conn.DB.Query(query, macAddress)
	if err != nil {
		return nil, fmt.Errorf("error al obtener alertas por macaddress: %w", err)
	}
	defer rows.Close()

	var alertas []entities.Alerta
	for rows.Next() {
		var a entities.Alerta
		err := rows.Scan(&a.ID, &a.MacAddress, &a.Mensaje, &a.Accion, &a.Nivel_peligro, &a.Fecha)
		if err != nil {
			return nil, fmt.Errorf("error al escanear alerta: %w", err)
		}
		alertas = append(alertas, a)
	}

	return alertas, nil
}