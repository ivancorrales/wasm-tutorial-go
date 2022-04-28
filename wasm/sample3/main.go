package main

import (
	"syscall/js"
)

func main() {
	document := js.Global().Get("document")
	h2 := document.Call("getElementById", "description")
	h2.Set("innerHTML", "Sample III | Appending HTML elements")

	h3 := document.Call("createElement", "h3")
	h3.Set("innerHTML", "It looks great!")

	div := document.Call("getElementById", "main")
	div.Call("appendChild", h3)

}
