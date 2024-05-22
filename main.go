package main

import (
	"fmt"

	"github.com/saulgil1/SegundoParcialGo/TSP"
)

func main() {

	// Funcion para leer el documento
	puntos := TSP.ReadDocument("dj38.tsp")

	// Funcion para vecino mas cercano
	RutaVecinoMasCercano := TSP.RecorridoVecinoMasCercano(puntos)
	DistanciaVecinoMasCercano := TSP.CalcularDistancia(RutaVecinoMasCercano, puntos)
	fmt.Println(DistanciaVecinoMasCercano)

	// Funcion para insercion

	// Funcion para vecindario
}
