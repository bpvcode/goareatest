package goareatest

import "math"

//Pi é um valor constante - constante PUBLIC
const Pi = 3.1416

//AreaCirculo - Calcula a área da circunferência - função PUBLIC
func AreaCirculo(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//AreaRetangulo - Calcula a área de um rectangulo - função PUBLIC
func AreaRetangulo(base, altura float64) float64 {
	return base * altura
}

//areaTriangulo - Calcula a área de um triangulo - função PRIVATE
func areaTriangulo(base, altura float64) float64 {
	return (base * altura) / 2
}
