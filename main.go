package main

import (
	"PolybianSquare/core"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

var TypeEncryption = map[int]string{
	0: "PolybianSquare",
	1: "Permutation",
}

func main() {
	gtk.Init(nil)

	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	err = b.AddFromFile("main.glade")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	obj, err := b.GetObject("window1")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	win := obj.(*gtk.Window)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	obj, _ = b.GetObject("input_text")
	entryText := obj.(*gtk.Entry)

	obj, _ = b.GetObject("output_text")
	outputText := obj.(*gtk.TextView)

	obj, _ = b.GetObject("encode_btn")
	encodeBtn := obj.(*gtk.Button)

	obj, _ = b.GetObject("decode_btn")
	decodeBtn := obj.(*gtk.Button)

	obj, _ = b.GetObject("encryption_type")
	encType, _ := obj.(*gtk.ComboBox)

	IndTypeEncryption := -1

	encType.Connect("changed", func() {
		IndTypeEncryption = encType.GetActive()
	})

	var square core.PolybianSquare
	var permutation core.Permutation

	encodeBtn.Connect("clicked", func() {
		switch TypeEncryption[IndTypeEncryption] {
		case "PolybianSquare":
			text, err := entryText.GetText()
			if err != nil {
				log.Fatal("Ошибка при обработке кодирования", err)
			}

			enc := core.NewPolybian(text)
			enc.GenerateKey()
			enc.Encoding()
			square = enc

			output, _ := outputText.GetBuffer()

			output.SetText(enc.EncodingText)
		case "Permutation":
			text, err := entryText.GetText()
			if err != nil {
				log.Fatal("Ошибка при обработке кодирования", err)
			}
			enc := core.NewPermutation(text)
			enc.Encode()

			permutation = enc

			output, _ := outputText.GetBuffer()
			output.SetText(enc.EncodingText)
		case "":
			output, _ := outputText.GetBuffer()

			output.SetText("Выберите тип шифрования")
		}
	})

	decodeBtn.Connect("clicked", func() {
		switch TypeEncryption[IndTypeEncryption] {
		case "PolybianSquare":
			text := square.Decoding()
			output, _ := outputText.GetBuffer()

			output.SetText(text)
		case "Permutation":
			decodeText := permutation.Decode()

			output, _ := outputText.GetBuffer()

			output.SetText(decodeText)
		case "":
			output, _ := outputText.GetBuffer()

			output.SetText("Выберите тип шифрования")
		}
	})

	win.ShowAll()

	gtk.Main()
}
