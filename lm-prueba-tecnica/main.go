package main

import (
	"fmt"
	"lm-prueba-tecnica/translate"
)

func main() {

	texto, _ := translate.MultiporpouseTranslate("LETTY HOLA", "TEXT", "MORSE")
	fmt.Println(texto)

}
