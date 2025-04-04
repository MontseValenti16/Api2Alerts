package mysql

import (
	"LifeGuardAlertas/src/core/db"
	"LifeGuardAlertas/src/datos/domain/entities"
	"LifeGuardAlertas/src/datos/domain/repository"
	"fmt"
	"time"
)

// Repositorio para datos anuales
type DatosAnualesMySQL struct {
	conn *db.ConnMySQL
}

func NewDatosAnualesMySQL() repository.DatosAnualesRepository {
	conn := db.GetDBPool()
	return &DatosAnualesMySQL{conn: conn}
}

func (r *DatosAnualesMySQL) GetAll() ([]*entities.DatoAnual, error) {
	query := "SELECT id_dato_anual, id_persona, anio, promediopasos, promediobpm, promediospo2, promediotemperatura FROM datos_anuales"
	rows, err := r.conn.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al consultar datos anuales: %w", err)
	}
	defer rows.Close()

	var datos []*entities.DatoAnual
	for rows.Next() {
		var dato entities.DatoAnual
		err := rows.Scan(&dato.ID, &dato.IDPersona, &dato.Anio, &dato.Promediopasos, &dato.Promediobpm, &dato.Promediospo2, &dato.Promediotemperatura)
		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %w", err)
		}
		datos = append(datos, &dato)
	}

	return datos, nil
}

// Repositorio para datos mensuales
type DatosMensualesMySQL struct {
	conn *db.ConnMySQL
}

func NewDatosMensualesMySQL() repository.DatosMensualRepository {
	conn := db.GetDBPool()
	return &DatosMensualesMySQL{conn: conn}
}

func (r *DatosMensualesMySQL) GetAll() ([]*entities.DatoMensual, error) {
	query := "SELECT id_dato_mensual, id_persona, fecha_inicio, fecha_fin, promediopasos, promediobpm, promediospo2, promediotemperatura FROM datos_mensuales"
	rows, err := r.conn.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al consultar datos mensuales: %w", err)
	}
	defer rows.Close()

	var datos []*entities.DatoMensual
	for rows.Next() {
		var dato entities.DatoMensual
		var fechaInicio, fechaFin string
		err := rows.Scan(&dato.ID, &dato.IDPersona, &fechaInicio, &fechaFin, &dato.Promediopasos, &dato.Promediobpm, &dato.Promediospo2, &dato.Promediotemperatura)
		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %w", err)
		}

		// Parsear fechas
		dato.FechaInicio, err = time.Parse("2006-01-02", fechaInicio)
		if err != nil {
			return nil, fmt.Errorf("error al parsear fecha_inicio: %w", err)
		}
		dato.FechaFin, err = time.Parse("2006-01-02", fechaFin)
		if err != nil {
			return nil, fmt.Errorf("error al parsear fecha_fin: %w", err)
		}

		datos = append(datos, &dato)
	}

	return datos, nil
}

// Repositorio para datos semanales
type DatosSemanalesMySQL struct {
	conn *db.ConnMySQL
}

func NewDatosSemanalesMySQL() repository.DatosSemanalesRepository {
	conn := db.GetDBPool()
	return &DatosSemanalesMySQL{conn: conn}
}

func (r *DatosSemanalesMySQL) GetAll() ([]*entities.DatoSemanal, error) {
	query := "SELECT id_dato_semanal, id_persona, fecha_inicio, fecha_fin, promediopasos, promediobpm, promediospo2, promediotemperatura FROM datos_semanales"
	rows, err := r.conn.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al consultar datos semanales: %w", err)
	}
	defer rows.Close()

	var datos []*entities.DatoSemanal
	for rows.Next() {
		var dato entities.DatoSemanal
		var fechaInicio, fechaFin string
		err := rows.Scan(&dato.ID, &dato.IDPersona, &fechaInicio, &fechaFin, &dato.Promediopasos, &dato.Promediobpm, &dato.Promediospo2, &dato.Promediotemperatura)
		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %w", err)
		}

		// Parsear fechas
		dato.FechaInicio, err = time.Parse("2006-01-02", fechaInicio)
		if err != nil {
			return nil, fmt.Errorf("error al parsear fecha_inicio: %w", err)
		}
		dato.FechaFin, err = time.Parse("2006-01-02", fechaFin)
		if err != nil {
			return nil, fmt.Errorf("error al parsear fecha_fin: %w", err)
		}

		datos = append(datos, &dato)
	}

	return datos, nil
}

// Repositorio para datos diarios
type DatosDiariosMySQL struct {
	conn *db.ConnMySQL
}

func NewDatosDiariosMySQL() repository.DatosDiariosRepository {
	conn := db.GetDBPool()
	return &DatosDiariosMySQL{conn: conn}
}

func (r *DatosDiariosMySQL) GetAll() ([]*entities.DatoDiario, error) {
	query := "SELECT id_dato_diario, id_persona, fecha, hora_inicio, hora_fin, promediopasos, promediobpm, promediospo2, promediotemperatura FROM datos_diarios"
	rows, err := r.conn.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al consultar datos diarios: %w", err)
	}
	defer rows.Close()

	var datos []*entities.DatoDiario
	for rows.Next() {
		var dato entities.DatoDiario
		err := rows.Scan(&dato.ID, &dato.IDPersona, &dato.Fecha, &dato.HoraInicio, &dato.HoraFin, &dato.Promediopasos, &dato.Promediobpm, &dato.Promediospo2, &dato.Promediotemperatura)
		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %w", err)
		}
		datos = append(datos, &dato)
	}

	return datos, nil
}
