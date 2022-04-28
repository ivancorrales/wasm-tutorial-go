package main

import (
	"fmt"
	"syscall/js"
	"time"
)

func main() {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	h2Element := js.Global().Get("document").Call("getElementById", "description")
	h2Element.Set("innerHTML", fmt.Sprintf("Sample II | Modifying the HTML (%s)", currentTime))
}
