package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/saulgil1/SegundoParcialGo/TSP"
)

/*
type Point struct {
	X float64
	Y float64
}
*/

func main() {
	// Abrir el archivo .tsp
	file, err := os.Open("dj38.tsp")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Variable para verificar si ya se encontró la sección NODE_COORD_SECTION
	encontrado := false

	// Lista para almacenar los puntos
	var puntos []TSP.Point

	// Escanear el archivo línea por línea
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()

		if encontrado {
			// Dividir la línea en partes separadas por espacio
			parts := strings.Fields(linea)
			if len(parts) != 3 {
				continue // Ignorar líneas que no tienen el formato correcto
			}
			// Convertir las coordenadas a números flotantes
			x, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				log.Fatal(err)
			}
			y, err := strconv.ParseFloat(parts[2], 64)
			if err != nil {
				log.Fatal(err)
			}
			// Crear un punto y agregarlo a la lista de puntos
			puntos = append(puntos, TSP.Point{X: x, Y: y})
		}

		// Buscar la línea que indica el comienzo de la sección NODE_COORD_SECTION
		if strings.TrimSpace(linea) == "NODE_COORD_SECTION" {
			encontrado = true
			continue
		}
	}

	// Verificar errores de escaneo
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Calcular la distancia entre dos puntos específicos (por ejemplo, punto 7 y punto 4)
	vecino := TSP.RecorridoVecinoMasCercano(puntos)
	for i := 0; i < len(vecino)-1; i++ {
		fmt.Println(TSP.EuclideanDistance(vecino[i], vecino[i+1]))
	}
}
