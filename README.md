# Práctica 2 — Taller de Coches Concurrente
**Sistemas Distribuidos — URJC**

## Descripción
Simulación de un taller de coches implementada en Go usando goroutines y channels.
Cada coche que llega es atendido por un mecánico. Si no hay mecánicos libres, 
el coche espera en cola. Si un coche acumula más de 15 segundos de atención, 
se le da prioridad y se contrata un nuevo mecánico si es necesario.

## Funcionalidades
- Cola de espera ilimitada para coches
- 3 especialidades: mecánica (5s), eléctrica (7s), carrocería (11s)
- Prioridad automática si un coche supera 15s acumulados
- Contratación dinámica de mecánicos en tiempo de ejecución
- Tests comparativos de rendimiento

## Estructura
| Archivo | Contenido |
|---|---|
| `types.go` | Estructuras: Mecanico, Coche, Taller |
| `mecanico.go` | Lógica de atención y contratación |
| `dispatcher.go` | Asignación de coches a mecánicos |
| `crud.go` | Gestión de clientes, vehículos e incidencias |
| `main.go` | Punto de entrada y simulación |
| `taller_test.go` | Tests de rendimiento |

## Cómo ejecutar
```bash
go run .
```

## Cómo ejecutar tests
```bash
go test -v -timeout 120s
```

## Fragmentos clave

### Channel como cola de espera
```go
// ColaCh es un channel con buffer grande que actúa como cola
colaCh := make(chan *Coche, 100)
// Enviar coche a la cola
colaCh <- coche
```
Un channel en Go permite comunicación segura entre goroutines.
El buffer define cuántos coches pueden esperar sin bloquear.

### Goroutine por mecánico
```go
// Cada mecánico corre en su propia goroutine
go atenderCoche(mecanico, coche, colaCh)
```
`go` lanza la función de forma concurrente. Varios mecánicos
trabajan simultáneamente sin bloquearse entre sí.

### Prioridad por tiempo acumulado
```go
if coche.TiempoAcumulado > 15 {
    coche.Prioritario = true
    colaCh <- coche  // reencola con prioridad
}
```
Si el tiempo supera 15s, el coche vuelve a la cola marcado
como prioritario para ser atendido antes.

## Resultados tests
| Test | Base | Modificado |
|---|---|---|
| Doble de coches | ~5s | ~10s |
| Doble de mecánicos | ~21s | ~11s |
| Distribución especialidades | ~11s | ~11s |

## Autor
molilobo