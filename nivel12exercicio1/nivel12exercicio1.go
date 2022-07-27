package main

import (
	"fmt"
)

// Imprimi Idade: + o valor da idade do animal
func main() {
	fmt.Println("Idade do animal:", Idade(10))
}

//Recebe um int e retorna um int
func Idade(i int) int {
	return i * 7
}

//  - Crie uma função Idade, que toma como parâmetro um número de anos e retorna a idade
// equivalente em anos caninos. (1 ano humano → 7 anos caninos)
// - Documente seu código com comentários, e utilize a função Idade na sua função main.
// - Rode seu programa para verificar se ele funciona.
// - Rode um local server com godoc e leia sua documentação.
