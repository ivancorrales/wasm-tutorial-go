package main

import (
	"syscall/js"
	"time"
)

func main() {
	document := js.Global().Get("document")
	message := document.Call("getElementById", "description")

	message.Get("style").Call("setProperty", "display", "block")
	time.Sleep(2 * time.Second)
	message.Get("style").Call("setProperty", "display", "none")
	time.Sleep(2 * time.Second)
	message.Get("style").Call("setProperty", "display", "block")
	time.Sleep(2 * time.Second)
	message.Get("style").Call("setProperty", "display", "none")
	time.Sleep(2 * time.Second)
	message.Get("style").Call("setProperty", "display", "block")
	time.Sleep(2 * time.Second)
	message.Get("style").Call("setProperty", "display", "none")
}
