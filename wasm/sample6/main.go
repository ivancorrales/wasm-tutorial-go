package main

import (
	"strings"
	"syscall/js"
)

func main() {
	document := js.Global().Get("document")
	input := document.Call("getElementById", "description").
		Get("innerHTML")
	outputMessage := strings.ToUpper(input.String())
	document.Call("getElementById", "descriptionUppercase").
		Set("innerHTML", outputMessage)
}
