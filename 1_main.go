package main

import (
	"fmt"
	"../try/urid"
)

var i = urid.Id

func increment() int{
	i++
	return i
}

func main(){
	fmt.Println("Number encode : 8687")
	fmt.Println()
	fmt.Println("decimal \t binary \t hexa \t UTF-8")
	fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	fmt.Println()
	urid.PrintVar()
	fmt.Println()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}