package main

import (
	"fmt"
)

func main(){

	funcTest := func(param1 string, param2 int)(string){
		return param1
	}

	anExpression := false
	test := []string{"test1","test2","test3"}

	//É um do while
	for ok := true;ok; ok = anExpression {
		fmt.Println("Teste 1")
	}

	//Loop em range
	for _, value := range test {
		fmt.Println(value)
	}

	value, _ := test_func2()

	fmt.Println(value)

	fmt.Println(test_func3(funcTest))

	funcao_retorno := test_func4()

	funcao_retorno("Samuel")

	test_func5()

	test_func6("s","sa","sam","samu","samue","samuel")
}

func test_func()(string, error){
	return "", nil
}

func test_func2()(retornoString string,retornoErro error){
	retornoString = "Samuel"
	retornoErro = nil
	return
}

func test_func3(value func(string, int)(string))(retornoString string,retornoErro error){
	retornoString = value("Otavio", 20)
	retornoErro = nil
	return
}


func test_func4() func(string){
    funcao_retorno := func(valorString string){
		fmt.Println("É uma string", valorString)
	}
	return funcao_retorno
}

func test_func5(){
    func(valorString string){
		fmt.Println("Não tem nome a função é anonima", valorString)
	}("Samuel")
}

func test_func6(valoresString ...string){
    for _, x := range valoresString{
		fmt.Println(x)
	}
}