package main

import "fmt"

func media(list[]int)(int,error){
	if len(list)==0{
		return 0, fmt.Errorf("vazio")
	}
	soma :=0
	for _,x:=range list{

	}

}

func main_2(){
	result, err:=media(]int{})
	if error!=nil{
	}
}

func divisao(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("A divisão não pode ser feita com 0")
	} else {
		return x / y, nil
	}
}

func main() {
	var x, y float64

	fmt.Println("valor x: ")
	fmt.Scan(&x)

	fmt.Println("valor y: ")
	fmt.Scan(&y)
	resultadoDivisao, err := divisao(x, y)

	if err != nil {
		fmt.Println("deu pal")
		return
	}
	fmt.Println("O resultado da divisão é: ", resultadoDivisao)
}
