package main

import (
	"fmt"
)

func dispatcher(
	mecanicos []*Mecanico,
	cola chan *Coche,
	results chan Resultado,
	totalCoches int,
) {
	nextID := len(mecanicos) + 1
	atendidos := 0

	for resultado := range results {
		atendidos++

		if resultado.Reencolar {
			atendidos--
			hayLibre := false
			for _, m := range mecanicos {
				if m.Libre {
					hayLibre = true
					break
				}
			}
			if !hayLibre {
				nuevo := &Mecanico{
					ID:           nextID,
					Nombre:       fmt.Sprintf("Mecánico%d", nextID),
					Especialidad: resultado.Incidencia,
					Activo:       true,
					Libre:        true,
				}
				nextID++
				mecanicos = append(mecanicos, nuevo)
				go atenderCoches(nuevo, cola, results)
				fmt.Printf("Nuevo mecánico %d contratado (%s)\n", nuevo.ID, nuevo.Especialidad)
			}
			resultado.CocheOriginal.Prioritario = true
			cola <- resultado.CocheOriginal
		} else {
			resultado.CocheOriginal.IncidenciaActual.Estado = "cerrada"
			nombreCliente := "desconocido"
			for _, c := range clientes {
				for _, v := range c.Vehiculos {
					if v.ID == resultado.CocheOriginal.Vehiculo.ID {
						nombreCliente = c.Nombre
						break
					}
				}
			}
			fmt.Printf("Coche %d terminado. Matricula: %s, Cliente: %s, Acumulado: %v, Incidencia: %s\n",
				resultado.CocheID,
				resultado.CocheOriginal.Vehiculo.Matricula,
				nombreCliente,
				resultado.TotalAcumulado,
				resultado.CocheOriginal.IncidenciaActual.Estado)
		}

		if atendidos == totalCoches {
			close(cola)
			return
		}
	}
}
