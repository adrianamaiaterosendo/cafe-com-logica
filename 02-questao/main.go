package main

import "fmt"

/*Considerando os fundamentos de lógica de programação e os operadores lógicos utilizados
nas linguagens de programação, atente-se para a expressão abaixo.
 C * (B – A) <= D – B / C
 Qual seria o resultado da execução dessa expressão,
 caso o valor das variáveis fossem: A=3; B=6; C=2 e D=9 ?*/

func main() {

	// var (
	// 	a = 3
	// 	b = 6
	// 	c = 2
	// 	d = 9
	// )

	//a := 2

	const(
		a = 3
		b = 6
		c = 2
		d = 9
	)

	//a = 2

	parte1 := c*(b-a)
	parte2 := d-b/c
	expressao := c*(b-a) <= d-b/c

	fmt.Println("O resultado da parte 1 é:", parte1)
	fmt.Println("O resultado da parte 2 é:", parte2)

	fmt.Println("O resultado da expressão é:", expressao)

}
