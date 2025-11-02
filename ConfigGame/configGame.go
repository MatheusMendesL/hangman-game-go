package ConfigGame

import (
	ConfigErrors "_027_exercicio/ErrosCfg"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Forca struct {
	Word        string
	SliceSpaces []string
	Used        map[string]bool
	Lifes       int
}

type Jogo interface {
	Chute() []string
}

var ForcaStruct = &Forca{}

/* func (Forca) LerEntrada() string {

} */

func (f *Forca) Chute(letters []string) {
	if f.Used == nil {
		f.Used = make(map[string]bool)
	}
	for {

		fmt.Println("Digite a letra")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)
		if len(chute) != 1 {
			fmt.Println("Erro: ", ConfigErrors.ErrLetters)
			continue
		}

		if f.Used[chute] {
			fmt.Println("Erro: ", ConfigErrors.ErrExists)
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

		res := f.FimDeJogo()
		if res {
			break
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
