package core

import (
	"PolybianSquare/constants"
	"math/rand"
)

type Encryption struct {
	Key                     [][]rune
	EncodingText            string
	MapPositionSymbol       map[rune][]int
	MapPositionEncodeSymbol map[rune][]int
	Text                    string
}

func New(text string) (e Encryption) {
	return Encryption{
		MapPositionSymbol:       make(map[rune][]int),
		MapPositionEncodeSymbol: make(map[rune][]int),
		Text:                    text,
	}
}

func (e *Encryption) GenerateKey() {
	helperMap := make(map[rune]int)
	textRunes := []rune(e.Text)
	var alphabetRunes []rune
	if textRunes[0] < 122 && textRunes[0] > 65 {
		alphabetRunes = []rune(constants.EnglishAlphabet)
	} else {
		alphabetRunes = []rune(constants.RussianAlphabet)
	}
	e.Key = make([][]rune, constants.LenY)

	for ind, _ := range e.Key {
		e.Key[ind] = make([]rune, constants.LenX)
	}

	for ind, val := range textRunes {
		if _, ok := e.MapPositionSymbol[val]; ok {
			continue
		}

		x := rand.Intn(constants.LenX)
		y := rand.Intn(constants.LenY)

		if e.Key[y][x] != 0 {
			for e.Key[y][x] != 0 {
				x = rand.Intn(constants.LenX)
				y = rand.Intn(constants.LenY)
			}
		}

		e.Key[y][x] = val
		e.MapPositionSymbol[val] = []int{y, x}
		helperMap[val] = ind
	}

	for i := 0; i < constants.LenY; i++ {
		for j := 0; j < constants.LenX; j++ {
			if e.Key[i][j] == 0 {
				randomRune := alphabetRunes[rand.Intn(len(alphabetRunes))]
				if _, ok := helperMap[randomRune]; ok {
					for ok {
						randomRune = alphabetRunes[rand.Intn(len(alphabetRunes))]
						_, ok = helperMap[randomRune]
					}
				}
				e.Key[i][j] = randomRune
				helperMap[randomRune] = i
			}
		}
	}
}

func (e *Encryption) Encoding() {
	textRunes := []rune(e.Text)
	result := make([]rune, len(textRunes))
	position := make([]int, 2)
	for ind, val := range textRunes {
		copy(position, e.MapPositionSymbol[val])
		position[0] += 1
		if position[0] == constants.LenY {
			position[0] = 0
		}

		e.MapPositionEncodeSymbol[e.Key[position[0]][position[1]]] = []int{position[0], position[1]}
		result[ind] = e.Key[position[0]][position[1]]
	}

	e.EncodingText = string(result)
}

func (e *Encryption) Decoding() string {
	textRunes := []rune(e.EncodingText)
	result := make([]rune, len(textRunes))
	position := make([]int, 2)
	for ind, val := range textRunes {
		copy(position, e.MapPositionEncodeSymbol[val])
		position[0] -= 1
		if position[0] < 0 {
			position[0] = constants.LenY - 1
		}

		result[ind] = e.Key[position[0]][position[1]]
	}

	return string(result)
}
