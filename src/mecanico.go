package main

import (
	"fmt"
	"time"
)

func tiempoAtencion(tipo string) time.Duration {
	switch tipo {
	case "mecanica":
		return 5 * time.Second
	case "electrica":
		return 7 * time.Second
	case "carroceria":
		return 11 * time.Second
	}
	return 5 * time.Second
}

func atenderCoches(m *Mecanico, cola chan *Coche, results chan Resultado) {
	for coche := range cola {
		m.Libre = false
		tipo := coche.TipoIncidencia()
		duracion := tiempoAtencion(tipo)
		fmt.Printf("Mecánico %s (%s) atendiendo coche %d - %s\n",
			m.Nombre, m.Especialidad, coche.ID, coche.Vehiculo.Matricula)
		time.Sleep(duracion)
		coche.TiempoAcumulado += duracion
		coche.IncidenciaActual.Estado = "en proceso"
		m.Libre = true

		results <- Resultado{
			CocheID:        coche.ID,
			MecanicoID:     m.ID,
			TiempoAtencion: duracion,
			TotalAcumulado: coche.TiempoAcumulado,
			Incidencia:     tipo,
			Reencolar:      coche.TiempoAcumulado > 15*time.Second && !coche.Prioritario,
			CocheOriginal:  coche,
		}
	}
}
