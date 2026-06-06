// types.go
package main

import "time"

type plaza struct {
	estado   string
	mecanico *Mecanico
	vehiculo *Vehiculo
}

type Cliente struct {
	ID        int
	Nombre    string
	Telefono  string
	Email     string
	Vehiculos []*Vehiculo
}

type Vehiculo struct {
	ID                    int
	Matricula             string
	Marca                 string
	Modelo                string
	FechaEntrada          time.Time
	FechaEstimSalida      time.Time
	IncidenciasDetectadas []*Incidencia
	Estado                string
}

type Incidencia struct {
	ID                 int
	MecanicosAsignados []*Mecanico
	Tipo               string // mecanica, electrica, carroceria
	Prioridad          string // baja, media, alta
	Descripcion        string
	Estado             string // abierta, en proceso, cerrada
}

type Mecanico struct {
	ID               int
	Nombre           string
	Especialidad     string
	AniosExperiencia int
	Activo           bool
	Libre            bool
}

type Coche struct {
	ID               int
	Vehiculo         *Vehiculo
	IncidenciaActual *Incidencia
	TiempoAcumulado  time.Duration
	Prioritario      bool
}

func (c *Coche) TipoIncidencia() string {
	return c.IncidenciaActual.Tipo
}

type Resultado struct {
	CocheID        int
	MecanicoID     int
	TiempoAtencion time.Duration
	TotalAcumulado time.Duration
	Incidencia     string
	Reencolar      bool
	CocheOriginal  *Coche
}

var clientes []*Cliente
var vehiculos []*Vehiculo
var incidencias []*Incidencia
var mecanicos []*Mecanico
var plazasTaller []*plaza
