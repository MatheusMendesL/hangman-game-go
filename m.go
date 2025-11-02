package main

import (
	"_027_exercicio/ConfigGame"
	"fmt"
)

var Forca = ConfigGame.ForcaStruct

func main() {

	fmt.Println("Jogo da forca!")
	fmt.Println("Tema: Animais")
	fmt.Println("Você tem 7 vidas apenas, caso errar, você perde uma")

	Forca.Init()
}
