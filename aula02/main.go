package main

import "fmt"

func main() {

	//CONDICIONAIS ~

	// No Go as condicionais são iguais as outras linguagens
	// tipo javascript porém sem parenteses,
	// também não é preciso colocar ;
	if 1 > 1 {
		fmt.Println("True")
	}else if 3 > 1{
		fmt.Println("True 2")
	}else{
		fmt.Println("False")
	}


	//IF com init
	// retorno fica apenas dentro do escopo do if
	if retorno := test(); retorno == "teste" {
		fmt.Println("True")
	}

	test := "teste"
	//SWITCH
	// ele por padrão não continua para o próximo case, 
	// apenas se colocar fallthrough, mas ele executará o próximo case, 
	// independente da condição
	switch test {
	case "teste","teste3":
		fmt.Println("Caiu primeira opção")
		fallthrough
	case "teste2":
		fmt.Println("Caiu segunda opção")
	case "teste4":
		fmt.Println("Caiu terceira opção")
	case "teste5":
		fmt.Println("Caiu quarta opção")
	default:
		fmt.Println("Caiu no default")
	}
}

func test() string {
	return "teste"
}