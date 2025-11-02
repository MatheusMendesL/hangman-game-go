// meu obj, fazer com que o reader leia e o contains verifica se o que eu enviar vai estar, eu entendi a ideia
// então vou fazer

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math/rand/v2"
	"os"
	"strings"
)

var errLetters = "Deve ser digitado uma letra apenas!"
var errExists = "Você já digitou essa letra"

func main() {

	fmt.Println("Jogo da forca!")
	fmt.Println("Tema: Animais")
	fmt.Println("Você tem 7 vidas apenas, caso errar, você perde uma")

	words := []string{"cachorro", "gato", "elefante", "leao", "tigre", "girafa", "coelho", "urso", "macaco", "cavalo"}
	x := rand.IntN(len(words))
	animal := words[x]
	reader := strings.NewReader(animal)
	buffer := make([]byte, 1)
	lifes := 7
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

	sliceSpaces := []string{}
	for i := 0; i < len(letters); i++ {
		sliceSpaces = append(sliceSpaces, "_")
	}

	fmt.Println(sliceSpaces)

	used := map[string]bool{}
	for {

		fmt.Println("Digite a letra")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)
		if len(chute) != 1 {
			fmt.Println("Erro: ", errLetters)
			continue
		}

		if used[chute] {
			fmt.Println("Erro: ", errExists)
			continue
		}

		used[chute] = true
		fmt.Println(contains(letters, chute))

		for i, v := range letters {
			if chute == v {
				sliceSpaces[i] = chute
				fmt.Println(sliceSpaces)
			}
		}

		if !contains(letters, chute) {
			lifes--
			fmt.Printf("Você tem apenas %d vidas! \n", lifes)
		}

		if lifes <= 0 {
			fmt.Printf("Você perdeu, a palavra era %s", animal)
			break
		}

		if !contains(sliceSpaces, "_") {
			fmt.Println("Fim de jogo, parabéns!")
			break
		}
	}
}

func contains[T comparable](s []T, cmp T) bool {
	for _, str := range s {
		if str == cmp {
			return true
		}
	}

	return false
}
