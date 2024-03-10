package core

import "strings"

type Permutation struct {
	Key          []int
	EncodingText string
	Text         string
}

func NewPermutation(text string) Permutation {
	return Permutation{
		Key:          nil,
		Text:         text,
		EncodingText: "",
	}
}

func (e *Permutation) Encode() {
	textRunes := []rune(e.Text)
	result := make([]rune, 0)

	for ind, val := range textRunes {
		if val == ' ' {
			e.Key = append(e.Key, ind)
		}
	}

	e.Text = strings.ToUpper(e.Text)

	e.Text = strings.ReplaceAll(e.Text, " ", "")

	textRunes = []rune(e.Text)

	for len(textRunes)%5 != 0 {
		textRunes = append(textRunes, 'O')
	}

	for i, j := 0, len(textRunes)-1; i < j; i, j = i+1, j-1 {
		textRunes[i], textRunes[j] = textRunes[j], textRunes[i]
	}

	for i, r := range textRunes {
		// Добавляем символ в результат
		result = append(result, r)

		// Если номер символа делится на 5 без остатка и это не последний символ
		if (i+1)%5 == 0 && i != len(textRunes)-1 {
			// Вставляем пробел
			result = append(result, ' ')
		}
	}

	e.EncodingText = string(result)
}

func (e *Permutation) Decode() string {
	e.EncodingText = strings.ReplaceAll(e.EncodingText, " ", "")
	textRunes := []rune(e.EncodingText)

	for i, j := 0, len(textRunes)-1; i < j; i, j = i+1, j-1 {
		textRunes[i], textRunes[j] = textRunes[j], textRunes[i]
	}

	result := make([]rune, 0)

	j := 0

	for _, val := range textRunes {
		result = append(result, val)

		if len(e.Key) == 0 {
			return "Нет ключа для этого слова"
		}

		if len(result) == e.Key[j] {
			result = append(result, ' ')
			if j != len(e.Key)-1 {
				j += 1
			}
		}
	}

	for len(result) != len([]rune(e.Text))+len(e.Key) {
		result = remove(result, len(result)-1)
	}

	return string(result)
}

func remove(slice []rune, s int) []rune {
	return append(slice[:s], slice[s+1:]...)
}
