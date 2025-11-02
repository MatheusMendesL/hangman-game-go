package ConfigGame

import (
	ConfigErrors "_027_exercicio/ErrosCfg"
	"bufio"
	"errors"
	"fmt"
	"io"
	"math/rand/v2"
	"os"
	"strings"
)

type Forca struct {
	Word        string
	SliceSpaces []string
	Used        map[string]bool
	Lifes       int
}

var ForcaStruct = &Forca{}

func (f *Forca) Init() {
	words := []string{"cachorro", "gato", "elefante", "leao", "tigre", "girafa", "coelho", "urso", "macaco", "cavalo"}
	x := rand.IntN(len(words))
	f.Word = words[x]
	reader := strings.NewReader(f.Word)
	buffer := make([]byte, 1)
	f.Lifes = 7
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
		f.SliceSpaces = append(f.SliceSpaces, "_")
	}

	fmt.Println(f.SliceSpaces)

	f.Chute(letters)
}

func (f *Forca) Chutar() (string, string) {
	fmt.Println("Digite a letra")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	chute := scanner.Text()
	chute = strings.TrimSpace(chute)
	if len(chute) != 1 {
		fmt.Println()
		return "", ConfigErrors.ErrLetters
	}

	if f.Used[chute] {
		return "", ConfigErrors.ErrExists
	}
	return chute, ""
}

func (f *Forca) Chute(letters []string) {

	f.Used = make(map[string]bool)
	for !f.FimDeJogo() {

		chute, err := f.Chutar()

		if err != "" {
			fmt.Println("Erro: ", err)
			continue
		}

		f.Used[chute] = true

		for i, v := range letters {
			if chute == v {
				f.SliceSpaces[i] = chute
				fmt.Println(f.SliceSpaces)
			}
		}

		if !Contains(letters, chute) {
			f.Lifes--
			fmt.Printf("Você tem apenas %d vidas! \n", f.Lifes)
		}
	}
}

func (f *Forca) FimDeJogo() bool {
	if f.Lifes <= 0 {
		fmt.Printf("Você perdeu, a palavra era %s", f.Word)
		return true
	}
	if !Contains(f.SliceSpaces, "_") {
		fmt.Println("Fim de jogo, parabéns!")
		return true
	}
	return false
}

func Contains[T comparable](s []T, cmp T) bool {
	for _, str := range s {
		if str == cmp {
			return true
		}
	}

	return false
}
