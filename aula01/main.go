package main

import "fmt"

func main() {

	//variaveis 


	var test2 interface{}

	test2 = "teste2"

	test := map[string]string{
		"teste": "teste",
	}

	var test3 int16 

	test3 = 8

	fmt.Println(test)
	fmt.Println(test2)
	fmt.Println(test3)


	//Arrays e Slices => Arrays tem tamanho definido; Slice tem tamanho indefinido
	//Array
	var test4[4]string = [4]string{"test","test","test"}
	fmt.Println(test4[0])
	//tamanho da quantidade de itens que tem
	fmt.Println(len(test4))
	//tamanho que é suportado
	fmt.Println(cap(test4))
	//Slice
	var test5[]string = []string{"test","test","test"}
	fmt.Println(test4[0])
	//tamanho da quantidade de itens que tem
	fmt.Println(len(test5))
	//tamanho que é suportado
	fmt.Println(cap(test5))
	test5 = append(test5,"3")
	//tamanho da quantidade de itens que tem
	fmt.Println(len(test5))
	//tamanho que é suportado
	fmt.Println(cap(test5))


	//STRUCTS
	//struct publica
	type User struct {
		name string
		age int
	}
	//struct privada
	type user struct {
		//variaveis publicas vistas por outros pacotes
		Name string
		Age int
		Nome_Mae string
		Email string
		//variaveis privadas, só vistas pelo pacote
		senha string
	}

	var user_object user = user{
		Name: "Samuel",
		Age: 25,
	}

	fmt.Println(user_object)
}


// função privada
func teste(){

}

// função publica 
func Teste(){

}

