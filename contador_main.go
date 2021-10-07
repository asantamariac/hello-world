package main

import "fmt"

func main() {
	encabezado := `
	##############################
	 Contador de numeros impares
	##############################
	`
	fmt.Println(encabezado)

	var (
		numero1  int
		numero2  int
		contador int
	)

	fmt.Println("Digita el primer numero: ")
	fmt.Scanln(&numero1)

	fmt.Println("Digita el segundo numero: ")
	fmt.Scanln(&numero2)

	for i := numero1; i < numero2; i++ {
		if i%2 != 0 {
			fmt.Printf("%d es impar \n", i)
			contador = contador + 1
		}
	}
	pie := `
	##############################
	 Resultado del conteo
	##############################
	`
	fmt.Println(pie)
	fmt.Printf("Entre %d y el %d hay %d numeros impares\n\n\n\n", numero1, numero2, contador)
}
