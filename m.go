package main

import (
	"_027_exercicio/ConfigGame"
	"errors"
	"fmt"
	"io"
	"math/rand/v2"
	"strings"
)

var Forca = ConfigGame.ForcaStruct

func main() {

	fmt.Println("Jogo da forca!")
	fmt.Println("Tema: Animais")
	fmt.Println("Você tem 7 vidas apenas, caso errar, você perde uma")

	words := []string{"cachorro", "gato", "elefante", "leao", "tigre", "girafa", "coelho", "urso", "macaco", "cavalo"}
	x := rand.IntN(len(words))
	Forca.Word = words[x]
	reader := strings.NewReader(Forca.Word)
	buffer := make([]byte, 1)
	Forca.Lifes = 7
	letters := []string{}

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		letters = append(letters, string(buffer[:n]))
	}

	for i := 0; i < len(letters); i++ {
		Forca.SliceSpaces = append(Forca.SliceSpaces, "_")
	}

	fmt.Println(Forca.SliceSpaces)

	Forca.Chute(letters)
}
