package main

import (
	"fmt"
	"testing"
	"time"
)

// helpers
func crearMecanicos(specs []string) []*Mecanico {
	mecs := []*Mecanico{}
	for i, esp := range specs {
		mecs = append(mecs, &Mecanico{
			ID:           i + 1,
			Nombre:       fmt.Sprintf("Mec%d", i+1),
			Especialidad: esp,
			Activo:       true,
			Libre:        true,
		})
	}
	return mecs
}

func crearCoches(tipos []string) []*Coche {
	coches := []*Coche{}
	for i, tipo := range tipos {
		inc := &Incidencia{ID: i + 1, Tipo: tipo, Estado: "abierta"}
		veh := &Vehiculo{ID: i + 1, Matricula: fmt.Sprintf("MAT-%d", i+1)}
		veh.IncidenciasDetectadas = append(veh.IncidenciasDetectadas, inc)
		coches = append(coches, &Coche{
			ID:               i + 1,
			Vehiculo:         veh,
			IncidenciaActual: inc,
		})
	}
	return coches
}

// test1
func TestDuplicarCoches(t *testing.T) {
	tipos1 := []string{"mecanica", "mecanica", "mecanica"}
	tipos2 := []string{"mecanica", "mecanica", "mecanica", "mecanica", "mecanica", "mecanica"}

	mecs1 := crearMecanicos([]string{"mecanica", "electrica", "carroceria"})
	mecs2 := crearMecanicos([]string{"mecanica", "electrica", "carroceria"})

	inicio1 := time.Now()
	m1 := ejecutarTaller(mecs1, crearCoches(tipos1))
	elapsed1 := time.Since(inicio1)

	inicio2 := time.Now()
	m2 := ejecutarTaller(mecs2, crearCoches(tipos2))
	elapsed2 := time.Since(inicio2)

	fmt.Println("\n=== Test 1: Duplicar coches (mecanica) ===")
	fmt.Printf("3 coches: total=%v, promedio=%v, tiempo_real=%v\n", m1.TiempoTotal, m1.TiempoPromedio, elapsed1)
	fmt.Printf("6 coches: total=%v, promedio=%v, tiempo_real=%v\n", m2.TiempoTotal, m2.TiempoPromedio, elapsed2)
}

// test2
func TestDuplicarMecanicos(t *testing.T) {
	coches := crearCoches([]string{"mecanica", "electrica", "carroceria", "mecanica", "electrica", "carroceria"})

	mecs3 := crearMecanicos([]string{"mecanica", "electrica", "carroceria"})
	mecs6 := crearMecanicos([]string{"mecanica", "electrica", "carroceria", "mecanica", "electrica", "carroceria"})

	inicio1 := time.Now()
	m1 := ejecutarTaller(mecs3, coches)
	elapsed1 := time.Since(inicio1)

	coches = crearCoches([]string{"mecanica", "electrica", "carroceria", "mecanica", "electrica", "carroceria"})

	inicio2 := time.Now()
	m2 := ejecutarTaller(mecs6, coches)
	elapsed2 := time.Since(inicio2)

	fmt.Println("\n=== Test 2: 3 mecánicos vs 6 mecánicos ===")
	fmt.Printf("3 mecs: total=%v, promedio=%v, tiempo_real=%v, mecs_finales=%d\n", m1.TiempoTotal, m1.TiempoPromedio, elapsed1, m1.MecanicosFinales)
	fmt.Printf("6 mecs: total=%v, promedio=%v, tiempo_real=%v, mecs_finales=%d\n", m2.TiempoTotal, m2.TiempoPromedio, elapsed2, m2.MecanicosFinales)
}

// test3
func TestDistribucionEspecialidades(t *testing.T) {
	coches := crearCoches([]string{"mecanica", "mecanica", "mecanica", "electrica", "carroceria"})

	// 3 mecánica, 1 eléctrica, 1 carrocería
	mecsA := crearMecanicos([]string{"mecanica", "mecanica", "mecanica", "electrica", "carroceria"})

	// resetear y crear config B
	coches2 := crearCoches([]string{"mecanica", "mecanica", "mecanica", "electrica", "carroceria"})
	// 1 mecánica, 3 eléctrica, 3 carrocería
	mecsB := crearMecanicos([]string{"mecanica", "electrica", "electrica", "electrica", "carroceria", "carroceria", "carroceria"})

	inicio1 := time.Now()
	m1 := ejecutarTaller(mecsA, coches)
	elapsed1 := time.Since(inicio1)

	inicio2 := time.Now()
	m2 := ejecutarTaller(mecsB, coches2)
	elapsed2 := time.Since(inicio2)

	fmt.Println("\n=== Test 3: 3mec/1elec/1carr vs 1mec/3elec/3carr ===")
	fmt.Printf("3/1/1: total=%v, promedio=%v, tiempo_real=%v\n", m1.TiempoTotal, m1.TiempoPromedio, elapsed1)
	fmt.Printf("1/3/3: total=%v, promedio=%v, tiempo_real=%v\n", m2.TiempoTotal, m2.TiempoPromedio, elapsed2)
}
