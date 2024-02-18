package main

import (
	"PolybianSquare/core"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

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

	var encryption core.Encryption

	encodeBtn.Connect("clicked", func() {
		text, err := entryText.GetText()
		if err != nil {
			log.Fatal("Ошибка при обработке кодирования", err)
		}

		enc := core.New(text)
		enc.GenerateKey()
		enc.Encoding()
		encryption = enc

		output, _ := outputText.GetBuffer()

		output.SetText(enc.EncodingText)
	})

	decodeBtn.Connect("clicked", func() {
		text := encryption.Decoding()
		output, _ := outputText.GetBuffer()

		output.SetText(text)
	})

	win.ShowAll()

	gtk.Main()
}
