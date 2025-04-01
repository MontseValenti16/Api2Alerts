package mysql

import (
	"LifeGuard/src/core/db"
	"LifeGuard/src/ds18b20/domain/entities"
	"fmt"
	"log"
)

type MySQL struct {
	db *db.MySQLDB
}

func NewMySQL() *MySQL {
	conn := db.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al conectar a la base de datos: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) Save(data *entities.MensajeMqtt) error {
	query := "INSERT INTO ds18b20 (temperatura) VALUES (?)"
	_, err := mysql.db.ExecutePreparedQuery(query, data.Temperatura)
	if err != nil {
		return fmt.Errorf("error al guardar el dato en la base de datos: %w", err)
	}
	return nil
}