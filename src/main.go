package main

import (
	"fmt"
	"time"
)

type Metricas struct {
	TotalCoches      int
	TiempoTotal      time.Duration
	TiempoPromedio   time.Duration
	MecanicosFinales int
}

func ejecutarTaller(mecInicial []*Mecanico, cochesEntrada []*Coche) Metricas {
	cola := make(chan *Coche, 100)
	results := make(chan Resultado, 100)

	for _, m := range mecInicial {
		go atenderCoches(m, cola, results)
	}

	done := make(chan Metricas)

	go func() {
		nextID := len(mecInicial) + 1
		atendidos := 0
		totalTiempo := time.Duration(0)
		mecActuales := mecInicial

		for resultado := range results {
			atendidos++
			if resultado.Reencolar {
				atendidos--
				hayLibre := false
				for _, m := range mecActuales {
					if m.Libre {
						hayLibre = true
						break
					}
				}
				if !hayLibre {
					nuevo := &Mecanico{
						ID:           nextID,
						Nombre:       fmt.Sprintf("Mecanico%d", nextID),
						Especialidad: resultado.Incidencia,
						Activo:       true,
						Libre:        true,
					}
					nextID++
					mecActuales = append(mecActuales, nuevo)
					go atenderCoches(nuevo, cola, results)
				}
				resultado.CocheOriginal.Prioritario = true
				cola <- resultado.CocheOriginal
			} else {
				totalTiempo += resultado.TotalAcumulado
			}

			if atendidos == len(cochesEntrada) {
				close(cola)
				done <- Metricas{
					TotalCoches:      len(cochesEntrada),
					TiempoTotal:      totalTiempo,
					TiempoPromedio:   totalTiempo / time.Duration(len(cochesEntrada)),
					MecanicosFinales: len(mecActuales),
				}
				return
			}
		}
	}()

	for _, c := range cochesEntrada {
		cola <- c
	}

	return <-done
}
func simularTaller() {
	cola := make(chan *Coche, 100)
	results := make(chan Resultado, 100)

	if len(mecanicos) == 0 {
		fmt.Println("No hay mecánicos. Crea mecánicos primero desde el menú.")
		return
	}

	mecActivos := []*Mecanico{}
	for _, m := range mecanicos {
		if m.Activo {
			mecActivos = append(mecActivos, m)
		}
	}

	if len(mecActivos) == 0 {
		fmt.Println("No hay mecánicos activos.")
		return
	}

	for _, m := range mecActivos {
		go atenderCoches(m, cola, results)
	}

	cochesSimulacion := []*Coche{}
	idCoche := 1
	for _, v := range vehiculos {
		for _, inc := range v.IncidenciasDetectadas {
			if inc.Estado == "abierta" {
				cochesSimulacion = append(cochesSimulacion, &Coche{
					ID:               idCoche,
					Vehiculo:         v,
					IncidenciaActual: inc,
				})
				idCoche++
			}
		}
	}

	if len(cochesSimulacion) == 0 {
		fmt.Println("No hay vehículos con incidencias abiertas.")
		return
	}

	go dispatcher(mecActivos, cola, results, len(cochesSimulacion))

	for _, c := range cochesSimulacion {
		fmt.Printf("Coche %d entra al taller (%s) - %s\n",
			c.ID, c.TipoIncidencia(), c.Vehiculo.Matricula)
		cola <- c
	}

	time.Sleep(60 * time.Second)
	fmt.Println("Simulación terminada.")
}

func main() {
	Demo()
	menu()
}
