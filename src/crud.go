package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// -------------------- Métodos --------------------

// Crear Vehículo para cliente
func (c *Cliente) CrearVehiculo() {
	var v Vehiculo
	fmt.Print("Ingrese ID del vehículo: ")
	fmt.Scan(&v.ID)
	fmt.Print("Ingrese matrícula: ")
	fmt.Scan(&v.Matricula)
	fmt.Print("Ingrese marca: ")
	fmt.Scan(&v.Marca)
	fmt.Print("Ingrese modelo: ")
	fmt.Scan(&v.Modelo)
	v.Estado = "Disponible"
	v.FechaEntrada = time.Now()
	fmt.Printf("Fecha de entrada : %v", v.FechaEntrada)
	v.Estado = "Disponible"

	ptr := &v
	c.Vehiculos = append(c.Vehiculos, ptr)
	vehiculos = append(vehiculos, ptr)
	fmt.Println("Vehículo creado con éxito.")
}

// Ver vehículos de cliente
func (c *Cliente) VerVehiculos() {
	fmt.Println("Vehículos de", c.Nombre)
	for _, v := range c.Vehiculos {
		fmt.Printf("ID: %d, Matrícula: %s, Marca: %s, Modelo: %s, Estado: %s \n",
			v.ID, v.Matricula, v.Marca, v.Modelo, v.Estado)
	}
}

// Asignar vehículo a taller
func AsignarVehiculoTaller() {
	var idVehiculo int
	fmt.Print("Ingrese ID del vehículo a asignar al taller: ")
	fmt.Scan(&idVehiculo)

	var v *Vehiculo
	for _, veh := range vehiculos {
		if veh.ID == idVehiculo {
			v = veh
			break
		}
	}
	if v == nil {
		fmt.Println("Vehículo no encontrado.")
		return
	}

	for _, p := range plazasTaller {
		if p.estado == "" {
			p.vehiculo = v
			p.estado = "Ocupada"
			v.Estado = "En taller"
			fmt.Printf("Vehículo %s asignado a plaza con mecánico %s\n", v.Matricula, p.mecanico.Nombre)

			return
		}
	}
	fmt.Println("No hay plazas disponibles en el taller.")
}

// Listar clientes con vehículos en taller
func ListarClientesEnTaller() {
	fmt.Println("Clientes con vehículos en taller:")
	for _, c := range clientes {
		for _, v := range c.Vehiculos {
			if v.Estado == "En taller" {
				fmt.Printf("Cliente: %s, Vehículo: %s (%s)\n", c.Nombre, v.Matricula, v.Modelo)
			}
		}
	}
}
func ListarPlazasEnTaller() {
	fmt.Println("Plazas del taller:")
	for i, p := range plazasTaller {
		// Verificar si la plaza tiene un vehículo asignado
		if p.vehiculo == nil {
			// Si no hay vehículo, imprimir que la plaza está vacía
			fmt.Printf("Plaza %d Estado: %s   (Vacía)\n", i, p.estado)
		} else {
			// Si la plaza tiene un vehículo, mostrar la información
			fmt.Printf("Plaza %d Estado: %s   Mecánico: %s   Matrícula Vehículo: %s\n", i, p.estado, p.mecanico.Nombre, p.vehiculo.Matricula)
		}
	}
}

// Listar incidencias de un vehículo
func ListarIncidenciasVehiculo() {
	var id int
	fmt.Print("Ingrese ID del vehículo: ")
	fmt.Scan(&id)
	for _, v := range vehiculos {
		if v.ID == id {
			if len(v.IncidenciasDetectadas) == 0 {
				fmt.Println("No hay incidencias para este vehículo.")
				return
			}
			fmt.Println("Incidencias del vehículo:")
			for _, inc := range v.IncidenciasDetectadas {
				fmt.Printf("ID: %d, Tipo: %s, Prioridad: %s, Estado: %s, Descripción: %s\n",
					inc.ID, inc.Tipo, inc.Prioridad, inc.Estado, inc.Descripcion)
			}
			return
		}
	}
	fmt.Println("Vehículo no encontrado.")
}

// modificar vehiculos
func ModificarVehiculo() {
	var id int
	fmt.Print("Ingrese el ID del vehiculo: ")
	fmt.Scan(&id)

	var op int

	// Iteramos a través de los vehículos usando el índice `i`
	for i := 0; i < len(vehiculos); i++ {
		v := vehiculos[i]
		if v.ID == id {
			for {
				fmt.Println("Vehículo encontrado")
				fmt.Println("1. Modificar datos")
				fmt.Println("2. Eliminar")
				fmt.Scan(&op)

				if op == 1 {
					// Modificar los datos del vehículo
					fmt.Print("Ingrese nueva matrícula: ")
					fmt.Scan(&v.Matricula)
					fmt.Print("Ingrese nueva marca: ")
					fmt.Scan(&v.Marca)
					fmt.Print("Ingrese nuevo modelo: ")
					fmt.Scan(&v.Modelo)
					fmt.Println("Vehículo modificado con éxito.")
					return // Salimos de la función una vez modificamos
				} else if op == 2 {
					// Eliminar el vehículo del arreglo global
					vehiculos = append(vehiculos[:i], vehiculos[i+1:]...) // Elimina el vehículo en el índice `i`
					fmt.Println("Vehículo eliminado con éxito.")
					for _, c := range clientes {
						for j := 0; j < len(c.Vehiculos); j++ {
							if c.Vehiculos[j].ID == id {
								// Elimina el vehículo del cliente
								c.Vehiculos = append(c.Vehiculos[:j], c.Vehiculos[j+1:]...)
								break // Terminamos después de encontrar el vehículo
							}
						}
					}

					// Eliminar el vehículo de las plazas del taller
					for _, p := range plazasTaller {
						if p.vehiculo != nil && p.vehiculo.ID == id {
							p.vehiculo = nil // Liberar la plaza
							p.estado = ""    // Cambiar el estado de la plaza
							break
						}
					}
					return // Salimos de la función después de eliminar
				}
			}
		}
	}

	// Si no encontramos el vehículo, mostramos un mensaje
	fmt.Println("Vehículo no encontrado.")
}

// Listar vehículos de un cliente
func ListarVehiculosCliente() {
	var id int
	fmt.Print("Ingrese ID del cliente: ")
	fmt.Scan(&id)
	for _, c := range clientes {
		if c.ID == id {
			c.VerVehiculos()
			return
		}
	}
	fmt.Println("Cliente no encontrado.")
}
func ModificarCliente() {
	var id int
	fmt.Print("Ingrese el ID del cliente: ")
	fmt.Scan(&id)

	// Buscar al cliente por su ID
	for i := 0; i < len(clientes); i++ {
		c := clientes[i]
		if c.ID == id {
			var op int
			for {
				// Mostrar el menú de opciones
				fmt.Printf("Cliente encontrado: %s\n", c.Nombre)
				fmt.Println("1. Modificar datos")
				fmt.Println("2. Eliminar cliente")
				fmt.Scan(&op)

				// Opción para modificar los datos del cliente
				if op == 1 {
					fmt.Print("Ingrese nuevo nombre: ")
					fmt.Scan(&c.Nombre)
					fmt.Print("Ingrese nuevo teléfono: ")
					fmt.Scan(&c.Telefono)
					fmt.Print("Ingrese nuevo email: ")
					fmt.Scan(&c.Email)
					fmt.Println("Cliente modificado con éxito.")
				} else if op == 2 {
					// Eliminar los vehículos del cliente
					for _, v := range c.Vehiculos {
						// Eliminar cada vehículo de la lista global de vehículos
						for j := 0; j < len(vehiculos); j++ {
							if vehiculos[j].ID == v.ID {
								// Eliminar el vehículo de la lista de vehículos global
								vehiculos = append(vehiculos[:j], vehiculos[j+1:]...)
								break
							}
						}

						// Liberar las plazas ocupadas por los vehículos en el taller
						for _, p := range plazasTaller {
							if p.vehiculo != nil && p.vehiculo.ID == v.ID {
								p.vehiculo = nil
								p.estado = "" // Liberamos la plaza
							}
						}
					}

					// Eliminar el cliente de la lista de clientes
					clientes = append(clientes[:i], clientes[i+1:]...)
					fmt.Println("Cliente y sus vehículos eliminados con éxito.")
					return
				} else {
					fmt.Println("Opción no válida. Intente de nuevo.")
				}
			}
		}
	}

	// Si no encontramos el cliente, mostramos un mensaje
	fmt.Println("Cliente no encontrado.")
}
func ModificarIncidencia() {
	var id int
	fmt.Print("Ingrese el ID de la incidencia: ")
	fmt.Scan(&id)

	// Buscar la incidencia por su ID
	for i := 0; i < len(incidencias); i++ {
		inc := incidencias[i]
		if inc.ID == id {
			var op int
			for {
				// Mostrar el menú de opciones
				fmt.Printf("Incidencia encontrada: %s\n", inc.Descripcion)
				fmt.Println("1. Modificar datos")
				fmt.Println("2. Eliminar incidencia")
				fmt.Scan(&op)

				// Opción para modificar los detalles de la incidencia
				if op == 1 {
					fmt.Print("Ingrese nueva descripción: ")
					fmt.Scan(&inc.Descripcion)
					fmt.Print("Ingrese nuevo tipo (mecánica/eléctrica/carrocería): ")
					fmt.Scan(&inc.Tipo)
					fmt.Print("Ingrese nueva prioridad (baja/media/alta): ")
					fmt.Scan(&inc.Prioridad)
					fmt.Print("Ingrese nuevo estado (abierta/en proceso/cerrada): ")
					fmt.Scan(&inc.Estado)
					fmt.Println("Incidencia modificada con éxito.")
					return
				} else if op == 2 {
					// Eliminar la incidencia de los vehículos asociados
					for _, v := range vehiculos {
						for j := 0; j < len(v.IncidenciasDetectadas); j++ {
							if v.IncidenciasDetectadas[j].ID == inc.ID {
								// Eliminar la incidencia de la lista de incidencias del vehículo
								v.IncidenciasDetectadas = append(v.IncidenciasDetectadas[:j], v.IncidenciasDetectadas[j+1:]...)
								break
							}
						}
					}

					// Eliminar la incidencia de los mecánicos asignados
					for _, m := range mecanicos {
						for j := 0; j < len(inc.MecanicosAsignados); j++ {
							if inc.MecanicosAsignados[j].ID == m.ID {
								// Eliminar el mecánico de la lista de mecánicos asignados
								inc.MecanicosAsignados = append(inc.MecanicosAsignados[:j], inc.MecanicosAsignados[j+1:]...)
								break
							}
						}
					}

					// Finalmente, eliminar la incidencia de la lista global de incidencias
					incidencias = append(incidencias[:i], incidencias[i+1:]...)
					fmt.Println("Incidencia eliminada con éxito.")
					return
				} else {
					fmt.Println("Opción no válida. Intente de nuevo.")
				}
			}
		}
	}

	// Si no encontramos la incidencia, mostramos un mensaje
	fmt.Println("Incidencia no encontrada.")
}

// Cambiar estado de incidencia
func CambiarEstadoIncidencia() {
	var id int
	fmt.Print("Ingrese ID de la incidencia: ")
	fmt.Scan(&id)
	for _, inc := range incidencias {
		if inc.ID == id {
			var nuevo string
			fmt.Print("Ingrese nuevo estado (abierta/en proceso/cerrada): ")
			fmt.Scan(&nuevo)
			inc.Estado = nuevo
			fmt.Println("Estado actualizado.")
			return
		}
	}
	fmt.Println("Incidencia no encontrada.")
}

// Listar mecánicos disponibles
func ListarMecanicosDisponibles() {
	fmt.Println("Mecánicos disponibles:")
	for _, m := range mecanicos {
		if m.Activo {
			asignado := false
			for _, inc := range incidencias {
				for _, mec := range inc.MecanicosAsignados {
					if mec.ID == m.ID {
						asignado = true
						break
					}
				}
			}
			if !asignado {
				fmt.Printf("ID: %d, Nombre: %s, Especialidad: %s\n", m.ID, m.Nombre, m.Especialidad)
			}
		}
	}
}
func ModificarMecanico() {
	var id int
	fmt.Print("Ingrese el ID del mecánico: ")
	fmt.Scan(&id)

	// Buscar el mecánico por su ID
	for i := 0; i < len(mecanicos); i++ {
		m := mecanicos[i]
		if m.ID == id {
			var op int
			var act int
			for {
				// Mostrar el menú de opciones
				fmt.Printf("Mecánico encontrado: %s\n", m.Nombre)
				fmt.Println("1. Modificar datos")
				fmt.Println("2. Eliminar mecánico")
				fmt.Scan(&op)

				// Opción para modificar los detalles del mecánico
				if op == 1 {
					fmt.Print("Ingrese nuevo nombre: ")
					fmt.Scan(&m.Nombre)
					fmt.Print("Ingrese nueva especialidad (mecánica/eléctrica/carrocería): ")
					fmt.Scan(&m.Especialidad)
					fmt.Print("Ingrese los años de experiencia: ")
					fmt.Scan(&m.AniosExperiencia)
					for {
						fmt.Print("¿Está activo? (1 si / 2 no): ")

						fmt.Scan(&act)
						if act == 1 {
							m.Activo = true
							break
						} else if act == 2 {
							m.Activo = false
							break
						}
					}

					fmt.Println("Mecánico modificado con éxito.")
					return
				} else if op == 2 {
					// Eliminar las incidencias asignadas a este mecánico
					for _, inc := range incidencias {
						for j := 0; j < len(inc.MecanicosAsignados); j++ {
							if inc.MecanicosAsignados[j].ID == m.ID {
								// Eliminar el mecánico de la lista de mecánicos asignados
								inc.MecanicosAsignados = append(inc.MecanicosAsignados[:j], inc.MecanicosAsignados[j+1:]...)
								break
							}
						}
					}

					// Eliminar las plazas de taller asignadas a este mecánico
					for j := 0; j < len(plazasTaller); j++ {
						if plazasTaller[j].mecanico != nil && plazasTaller[j].mecanico.ID == m.ID {
							// Eliminar la plaza de taller
							plazasTaller = append(plazasTaller[:j], plazasTaller[j+1:]...)
							break
						}
					}

					// Eliminar el mecánico de la lista global de mecánicos
					mecanicos = append(mecanicos[:i], mecanicos[i+1:]...)
					fmt.Println("Mecánico eliminado con éxito.")
					return
				} else {
					fmt.Println("Opción no válida. Intente de nuevo.")
				}
			}
		}
	}

	// Si no encontramos el mecánico, mostramos un mensaje
	fmt.Println("Mecánico no encontrado.")
}

// Listar incidencias de un mecánico
func ListarIncidenciasMecanico() {
	var id int
	fmt.Print("Ingrese ID del mecánico: ")
	fmt.Scan(&id)
	for _, m := range mecanicos {
		if m.ID == id {
			fmt.Printf("Incidencias asignadas a %s:\n", m.Nombre)
			for _, inc := range incidencias {
				for _, mec := range inc.MecanicosAsignados {
					if mec.ID == m.ID {
						fmt.Printf("ID: %d, Descripción: %s, Estado: %s\n", inc.ID, inc.Descripcion, inc.Estado)
					}
				}
			}
			return
		}
	}
	fmt.Println("Mecánico no encontrado.")
}
func DarAltaBajaMecanico() {
	var id int
	var op int
	fmt.Print("Ingrese ID del mecánico: ")
	fmt.Scan(&id)
	for _, m := range mecanicos {
		if m.ID == id {
			fmt.Printf("Estado de %s actual : %t\n", m.Nombre, m.Activo)
			fmt.Print("1.Dar de alta \n")
			fmt.Print("2.Dar de baja \n")
			fmt.Scan(&op)
			if op == 1 {
				m.Activo = true

			} else if op == 2 {
				m.Activo = false
			}
			return
		}
	}
	fmt.Println("Mecánico no encontrado.")

}

// -------------------- Menú principal --------------------
func menu() {
	for {
		fmt.Println("\n--- Menú Principal ---")
		fmt.Println("1. Gestión Taller")
		fmt.Println("2. Gestión Vehículos")
		fmt.Println("3. Gestión Clientes")
		fmt.Println("4. Gestión Incidencias")
		fmt.Println("5. Gestión Mecánicos")
		fmt.Println("6. Simular Taller (concurrente)")
		fmt.Println("7. Salir")
		var op int
		var sub int
		fmt.Scan(&op)

		switch op {
		case 1:
			fmt.Println("1. Asignar vehículo a plaza")
			fmt.Println("2. Listar clientes con vehículos en taller")
			fmt.Println("3. Listar plazas ")
			fmt.Scan(&sub)
			if sub == 1 {
				AsignarVehiculoTaller()
			} else if sub == 2 {
				ListarClientesEnTaller()
			} else if sub == 3 {
				ListarPlazasEnTaller()
			}
		case 2:
			fmt.Println("1. Crear vehículo")
			fmt.Println("2. Ver vehículos de cliente")
			fmt.Println("3. Listar incidencias de vehículo")
			fmt.Println("4. Modificar vehículo")
			fmt.Scan(&sub)
			if sub == 1 {
				CrearVehiculo()
			} else if sub == 2 {
				ListarVehiculosCliente()
			} else if sub == 3 {
				ListarIncidenciasVehiculo()
			} else if sub == 4 {
				ModificarVehiculo()
			}
		case 3:
			fmt.Println("1. Crear cliente")
			fmt.Println("2. Ver clientes")
			fmt.Println("3. Modificar Cliente")
			fmt.Scan(&sub)
			if sub == 1 {
				CrearCliente()
			} else if sub == 2 {
				VerClientes()
			} else if sub == 3 {
				ModificarCliente()
			}
		case 4:
			fmt.Println("1. Crear incidencia")
			fmt.Println("2. Ver incidencias")
			fmt.Println("3. Cambiar estado de incidencia")
			fmt.Println("4. Modificar incidencia")
			fmt.Scan(&sub)
			if sub == 1 {
				CrearIncidencia()
			} else if sub == 2 {
				VerIncidencias()
			} else if sub == 3 {
				CambiarEstadoIncidencia()
			} else if sub == 4 {
				ModificarIncidencia()
			}
		case 5:
			fmt.Println("1. Crear mecánico")
			fmt.Println("2. Ver mecánicos")
			fmt.Println("3. Listar mecánicos disponibles")
			fmt.Println("4. Listar incidencias de mecánico")
			fmt.Println("5. Modificar Mecanicos")
			fmt.Println("6. Dar Baja/Alta Mecanicos")
			fmt.Scan(&sub)
			if sub == 1 {
				CrearMecanico()
			} else if sub == 2 {
				VerMecanicos()
			} else if sub == 3 {
				ListarMecanicosDisponibles()
			} else if sub == 4 {
				ListarIncidenciasMecanico()
			} else if sub == 5 {
				ModificarMecanico()
			} else if sub == 6 {
				DarAltaBajaMecanico()
			}
		case 6:
			simularTaller()
		case 7:
			return
		}
	}
}

// -------------------- Funciones auxiliares --------------------
var reader = bufio.NewReader(os.Stdin)

func leerLinea() string {
	texto, _ := reader.ReadString('\n')
	return strings.TrimRight(texto, "\r\n")
}
func CrearCliente() {
	var c Cliente
	fmt.Print("ID cliente: ")
	fmt.Scan(&c.ID)
	reader.ReadString('\n')

	fmt.Print("Nombre: ")
	c.Nombre = leerLinea()

	fmt.Print("Teléfono: ")
	c.Telefono = leerLinea()

	fmt.Print("Email: ")
	c.Email = leerLinea()

	clientes = append(clientes, &c)
	fmt.Println("Cliente creado.")
}

func VerClientes() {
	fmt.Println("Clientes:")
	for _, c := range clientes {
		fmt.Printf("ID: %d, Nombre: %s\n", c.ID, c.Nombre)
	}
}

func CrearVehiculo() {
	var idCliente int
	fmt.Print("ID cliente: ")
	fmt.Scan(&idCliente)
	for _, c := range clientes {
		if c.ID == idCliente {
			c.CrearVehiculo()
			return
		}
	}
	fmt.Println("Cliente no encontrado.")
}

func CrearIncidencia() {
	var inc Incidencia
	fmt.Print("ID incidencia: ")
	fmt.Scan(&inc.ID)
	fmt.Print("Descripción: ")
	fmt.Scan(&inc.Descripcion)
	fmt.Print("Tipo (mecánica/elétrica/carrocería): ")
	fmt.Scan(&inc.Tipo)
	fmt.Print("Prioridad (baja/media/alta): ")
	fmt.Scan(&inc.Prioridad)
	inc.Estado = "abierta"
	ptr := &inc
	incidencias = append(incidencias, ptr)

	// Asociar incidencia a vehículo
	var idVeh int
	fmt.Print("ID del vehículo para asignar la incidencia: ")
	fmt.Scan(&idVeh)
	for _, v := range vehiculos {
		if v.ID == idVeh {
			v.IncidenciasDetectadas = append(v.IncidenciasDetectadas, ptr)
			break
		}
	}
	// Asociar incidencia a vehículo
	var idMec int
	fmt.Print("ID del mecanico para asignar la incidencia: ")
	fmt.Scan(&idMec)
	for _, m := range mecanicos {
		if m.ID == idMec {
			inc.MecanicosAsignados = append(inc.MecanicosAsignados, m)
			break
		}
	}

	fmt.Println("Incidencia creada.")
}

func VerIncidencias() {
	fmt.Println("Incidencias:")
	for _, inc := range incidencias {
		fmt.Printf("ID: %d, Descripción: %s, Estado: %s\n", inc.ID, inc.Descripcion, inc.Estado)
	}
}

func CrearMecanico() {
	var m Mecanico
	fmt.Print("ID mecánico: ")
	fmt.Scan(&m.ID)
	reader.ReadString('\n')

	fmt.Print("Nombre: ")
	m.Nombre = leerLinea()

	fmt.Print("Especialidad (mecanica/electrica/carroceria): ")
	m.Especialidad = leerLinea()

	m.Activo = true
	m.Libre = true
	fmt.Print("Años de experiencia: ")
	fmt.Scan(&m.AniosExperiencia)

	mecanicos = append(mecanicos, &m)
	for i := 0; i < 2; i++ {
		plazasTaller = append(plazasTaller, &plaza{mecanico: &m})
	}
	fmt.Println("Mecánico creado y plazas asignadas.")
}

func VerMecanicos() {
	fmt.Println("Mecánicos:")
	for _, m := range mecanicos {
		fmt.Printf("ID: %d, Nombre: %s, Especialidad: %s, Activo: %t\n", m.ID, m.Nombre, m.Especialidad, m.Activo)
	}
}

// -------------------- Demo --------------------
func Demo() {
	// Mecánicos
	m1 := &Mecanico{ID: 1, Nombre: "Juan Pérez", Especialidad: "mecanica", AniosExperiencia: 5, Activo: true, Libre: true}
	m2 := &Mecanico{ID: 2, Nombre: "Ana López", Especialidad: "electrica", AniosExperiencia: 3, Activo: true, Libre: true}
	m3 := &Mecanico{ID: 3, Nombre: "Luis García", Especialidad: "carroceria", AniosExperiencia: 7, Activo: true, Libre: true}
	mecanicos = append(mecanicos, m1, m2, m3)

	// Plazas
	for i := 0; i < 2; i++ {
		plazasTaller = append(plazasTaller, &plaza{mecanico: m1})
		plazasTaller = append(plazasTaller, &plaza{mecanico: m2})
		plazasTaller = append(plazasTaller, &plaza{mecanico: m3})
	}

	// Clientes
	c1 := &Cliente{ID: 1, Nombre: "Carlos García", Telefono: "123456789", Email: "carlos@mail.com"}
	c2 := &Cliente{ID: 2, Nombre: "María López", Telefono: "987654321", Email: "maria@mail.com"}
	c3 := &Cliente{ID: 3, Nombre: "Pedro Martínez", Telefono: "555444333", Email: "pedro@mail.com"}
	c4 := &Cliente{ID: 4, Nombre: "Laura Sánchez", Telefono: "111222333", Email: "laura@mail.com"}
	clientes = append(clientes, c1, c2, c3, c4)

	// Vehículos
	v1 := &Vehiculo{ID: 1, Matricula: "ABC-123", Marca: "Toyota", Modelo: "Corolla", FechaEntrada: time.Now(), Estado: "Disponible"}
	v2 := &Vehiculo{ID: 2, Matricula: "XYZ-456", Marca: "Honda", Modelo: "Civic", FechaEntrada: time.Now(), Estado: "Disponible"}
	v3 := &Vehiculo{ID: 3, Matricula: "DEF-789", Marca: "Ford", Modelo: "Focus", FechaEntrada: time.Now(), Estado: "Disponible"}
	v4 := &Vehiculo{ID: 4, Matricula: "GHI-012", Marca: "BMW", Modelo: "Serie3", FechaEntrada: time.Now(), Estado: "Disponible"}
	v5 := &Vehiculo{ID: 5, Matricula: "JKL-345", Marca: "Audi", Modelo: "A4", FechaEntrada: time.Now(), Estado: "Disponible"}

	c1.Vehiculos = append(c1.Vehiculos, v1)
	c2.Vehiculos = append(c2.Vehiculos, v2)
	c3.Vehiculos = append(c3.Vehiculos, v3)
	c4.Vehiculos = append(c4.Vehiculos, v4, v5)
	vehiculos = append(vehiculos, v1, v2, v3, v4, v5)

	// Incidencias
	inc1 := &Incidencia{ID: 1, Tipo: "mecanica", Prioridad: "media", Descripcion: "Cambio de aceite", Estado: "abierta", MecanicosAsignados: []*Mecanico{m1}}
	inc2 := &Incidencia{ID: 2, Tipo: "electrica", Prioridad: "alta", Descripcion: "Fallo batería", Estado: "abierta", MecanicosAsignados: []*Mecanico{m2}}
	inc3 := &Incidencia{ID: 3, Tipo: "carroceria", Prioridad: "baja", Descripcion: "Golpe puerta", Estado: "abierta", MecanicosAsignados: []*Mecanico{m3}}
	inc4 := &Incidencia{ID: 4, Tipo: "mecanica", Prioridad: "alta", Descripcion: "Frenos desgastados", Estado: "abierta", MecanicosAsignados: []*Mecanico{m1}}
	inc5 := &Incidencia{ID: 5, Tipo: "electrica", Prioridad: "media", Descripcion: "Fallo alternador", Estado: "abierta", MecanicosAsignados: []*Mecanico{m2}}

	v1.IncidenciasDetectadas = append(v1.IncidenciasDetectadas, inc1)
	v2.IncidenciasDetectadas = append(v2.IncidenciasDetectadas, inc2)
	v3.IncidenciasDetectadas = append(v3.IncidenciasDetectadas, inc3)
	v4.IncidenciasDetectadas = append(v4.IncidenciasDetectadas, inc4)
	v5.IncidenciasDetectadas = append(v5.IncidenciasDetectadas, inc5)
	incidencias = append(incidencias, inc1, inc2, inc3, inc4, inc5)

	fmt.Println("=== Demo cargada ===")
}
