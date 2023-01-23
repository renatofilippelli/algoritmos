package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	entrada := []int{2, 3, 4, 55, 6, 7}
	esperado := 62
	//return false
	fmt.Println(calc(entrada, esperado))

	entrada = []int{2, 3, 4, 55, 6, 7}
	esperado = 12
	//return false
	fmt.Println(calc(entrada, esperado))

}

func calc(entrada []int, esperado int) bool {

	//array int para string pura
	entradaTexto := strings.Trim(strings.Replace(fmt.Sprint(entrada), " ", " ", -1), "[]")

	//percorrer array
	for a := 0; a < len(entrada); a++ {

		//pega valor do arragity na posição a
		valor1 := entrada[a]

		//calcula do resultado esperado atual do array posição a
		resultado := strconv.Itoa(esperado - valor1)

		//string sem o valor da posição a
		stringarray := entradaTexto[strings.Index(entradaTexto, strconv.Itoa(valor1))+1:]

		//procura na string pura se existe resultado esperado
		if strings.Contains(stringarray, resultado) {
			return true
		}
	}

	return false
}
