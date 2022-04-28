package main

import (
	"strings"
	"syscall/js"
)

func main() {
	document := js.Global().Get("document")
	input := document.Call("getElementById", "inputText").Get("innerHTML")
	println(input.String())
	outputMessage := strings.ToUpper(input.String())
	println(outputMessage)
	document.Call("getElementById", "outputText").Set("innerHTML", outputMessage)

}
