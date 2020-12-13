package main

import "fmt"

func main() {
	var a int = 10

	fmt.Println(a)

	fmt.Println("&", &a) //apunta al valor
	valRef(&a)

	fmt.Println(a)
}

func valRef(p *int) {
	*p = 30
	fmt.Println("*", *p) //leer valor por el puntero

}
