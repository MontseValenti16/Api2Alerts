/*
    
      "condicion": "oxigenacion < 95",
      "mensaje": "El oxímetro ha detectado una baja en la oxigenación.",
      "accion": "enviar alerta",
      "nivel_peligro": "moderado"
    
    
      "tipo": "oxigenacion",
      "sensor": "max30102",
      "condicion": "oxigenacion < 90",
      "mensaje": "La oxigenación ha bajado a un nivel peligroso.",
      "accion": "enviar aviso a 911 o institución de emergencia",
      "nivel_peligro": "peligroso"
    
    
      "condicion": "frecuencia_cardiaca < 60",
      "mensaje": "Frecuencia cardíaca baja detectada.",
      "accion": "alerta si se acompaña de síntomas como mareos, fatiga o desmayos",
      "nivel_peligro": "peligroso"
    
    
      "condicion": "frecuencia_cardiaca > 100",
      "mensaje": "Frecuencia cardíaca alta detectada.",
      "accion": "alerta si es persistente sin ejercicio o estrés",
      "nivel_peligro": "peligroso"
    

      "condicion": "aceleracion > umbral_caida || giroscopio > umbral_caida",
      "mensaje": "Posible caída detectada. Se verificará el estado del usuario.",
      "accion": "preguntar si está bien, si no responde en 10 segundos, enviar aviso de emergencia",
      "nivel_peligro": "peligroso"
    
    
      "condicion": "temperatura < 35",
      "mensaje": "Temperatura corporal muy baja detectada.",
      "accion": "necesita atención médica inmediata",
      "nivel_peligro": "peligroso"
    
      "condicion": "temperatura > 40",
      "mensaje": "Temperatura corporal extremadamente alta detectada.",
      "accion": "es una emergencia médica, puede ser causada por golpes de calor o enfermedades severas",
      "nivel_peligro": "muy peligroso"
    
*/


package application

import (
	"LifeGuardAlertas/ds18b20/domain/repository"
)
type GetAllHipertermiaUseCase struct {
	hipertermiaRepo *repository.Hipertermia
}

func NewGetAllHipertermiaUseCase(hipertermia repository.Hipertermia) *GetAllHipertermiaUseCase {
	return &GetAllHipertermiaUseCase{
		hipertermiaRepo: hipertermia,
	}
}

func (u *GetAllHipertermiaUseCase) Run() ([]*entities, error) {
	return u.hipertermiaRepo.FindAll()
}