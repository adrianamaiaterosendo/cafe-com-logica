package main

import "fmt"

/* Considerando a, b, c e d como variáveis com valores iniciais iguais a 5, 7, 3 e 9, respectivamente, assinale a opção correta.
Alternativas
A
O resultado da expressão (a != 3 || b < 10 || c == 5) é falso.
B
O resultado da expressão (d > 8 && c == 3 || a >= 10) é verdadeiro.
C
O resultado da expressão !(d == 12 && a != 10) é falso.
D
O resultado da expressão (c == 4 || d <=6) && (a >= 5 && b !=9) || ( ! (a < 5) ) é falso.
E
O resultado da expressão (a == 3 || b > 10 || d == 8 ) é verdadeiro.
*/

func main() {
	a := 5
	b := 7
	c := 3
	d := 9

	expressao1 := (a != 3 || b < 10 || c == 5)
	expressao2 := (d > 8 && c == 3 || a >= 10)
	expressao3 := !(d == 12 && a != 10)
	expressao4 := (c == 4 || d <= 6) && (a >= 5 && b != 9) || (!(a < 5))
	expressao5 := (a == 3 || b > 10 || d == 8)

	fmt.Println("expressão 1 ", expressao1)
	fmt.Println("expressão 2 ", expressao2)
	fmt.Println("expressão 3 ", expressao3)
	fmt.Println("expressão 4 ", expressao4)
	fmt.Println("expressão 5 ", expressao5)

}
