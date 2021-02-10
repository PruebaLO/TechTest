package translate

import "fmt"

var morseList = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

//BinaryToText: traduce de binario a texto
func BinaryToText(textoATraducir string) string {
	var binary = ""

	fmt.Println("BinaryToText")

	return binary

}

//TextToBinary: traduce de texto a binario
func TextToBinary(textoATraducir string) string {
	var text = ""

	for _, val := range textoATraducir {

		if val == 32 {
			text = fmt.Sprintf("%s   ", text)
		} else {
			text = fmt.Sprintf("%s%b ", text, val)
		}
	}

	return text

}

//BinaryToMorse: traduce de binario a Morse
func BinaryToMorse(textoATraducir string) string {
	var binary = ""

	fmt.Println("BinaryToMorse")

	return binary
}

//MorseToBinary: traduce de Morse a binario
func MorseToBinary(textoATraducir string) string {
	var morse = ""

	fmt.Println("MorseToBinary")

	return morse
}

//MorseToText: traduce de Morse a Texto
func MorseToText(textoATraducir string) string {
	var morse = ""

	fmt.Println("MorseToBinary")

	return morse
}

//TextoToMose: traduce de texto a Morse
func TextoToMorse(textoATraducir string) string {
	var text = ""

	for _, val := range textoATraducir {

		if val == 32 {
			text = fmt.Sprintf("%s   ", text)
		} else {
			text = fmt.Sprintf("%s%v ", text, morseList[val-65])
		}
	}

	return text
}
