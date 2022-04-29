package main

import (
	"syscall/js"
)

func main() {
	document := js.Global().Get("document")

	h2 := document.Call("createElement", "h2")
	h2.Set("innerHTML", "Sample 3 | Appending HTML elements")

	div := document.Call("getElementById", "main")
	div.Call("appendChild", h2)

}
