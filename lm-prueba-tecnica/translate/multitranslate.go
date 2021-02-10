package translate

import (
	"fmt"
)

//TEXT, BINARY, MORSE, etc

func MultiporpouseTranslate(textoATraducir string, formatoOrigen string, formatoDestino string) (string, error) {
	var textoTraducido string

	tipoTraduccion := fmt.Sprintf("%s%s", formatoOrigen, formatoDestino)

	switch tipoTraduccion {
	case "BINARYTEXT":
		textoTraducido = BinaryToText(textoATraducir)
	case "TEXTBINARY":
		textoTraducido = TextToBinary(textoATraducir)
	case "BINARYMORSE":
		textoTraducido = BinaryToMorse(textoATraducir)
	case "MORSEBINARY":
		textoTraducido = MorseToBinary(textoATraducir)
	case "MORSETEXT":
		textoTraducido = MorseToText(textoATraducir)
	case "TEXTMORSE":
		textoTraducido = TextoToMorse(textoATraducir)
	default:
		fmt.Println("No es un tipo de conversion valido")
	}

	return textoTraducido, nil
}
