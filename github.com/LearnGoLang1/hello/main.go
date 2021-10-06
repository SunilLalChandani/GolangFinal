package main

import (
	"fmt"

	"github.com/ddadumitrescu/hellomod"
)

func main() {
	MeFirst()
	hellomod.Salut()
	Callme()

}

func MeFirst() {
	fmt.Println("MeFirst")

}
