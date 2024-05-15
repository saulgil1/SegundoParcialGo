package TSP

import (
	"math"
)

func VecinoMasCercano(puntoRef Point, puntos []Point) Point {
	var puntoCercano Point
	minDistancia := math.Inf(1)

	for i, punto := range puntos {
		if puntoRef == puntos[i] {
			continue
		}
		dist := float64(EuclideanDistance(puntoRef, punto))
		if dist < minDistancia {
			minDistancia = dist
			puntoCercano = punto
		}
	}

	return puntoCercano
}

func RecorridoVecinoMasCercano(puntos []Point) []Point {
	var recorrido []Point
	puntosRestantes := make([]Point, len(puntos))
	copy(puntosRestantes, puntos)

	puntoActual := puntos[0] // Puedes elegir cualquier punto de inicio
	recorrido = append(recorrido, puntoActual)

	// Eliminar el punto de inicio de puntosRestantes
	puntosRestantes = puntosRestantes[1:]

	for len(puntosRestantes) > 0 {
		puntoCercano := VecinoMasCercano(puntoActual, puntosRestantes)
		recorrido = append(recorrido, puntoCercano)

		// Actualizar puntoActual y eliminar el puntoCercano de puntosRestantes
		puntoActual = puntoCercano
		for i, punto := range puntosRestantes {
			if punto == puntoCercano {
				puntosRestantes = append(puntosRestantes[:i], puntosRestantes[i+1:]...)
				break
			}
		}
	}

	return recorrido
}
