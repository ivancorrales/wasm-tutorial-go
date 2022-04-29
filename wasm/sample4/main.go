package main

import (
	"syscall/js"
	"time"
)

func addGreenClass(classes js.Value) {
	if classes.Call("contains", "pink").Bool() {
		classes.Call("remove", "pink")
	}
	classes.Call("add", "green")
}

func addPinkClass(classes js.Value) {
	if classes.Call("contains", "green").Bool() {
		classes.Call("remove", "green")
	}
	classes.Call("add", "pink")
}

func main() {
	document := js.Global().Get("document")
	message := document.Call("getElementById", "simpleText")
	classes := message.Get("classList")
	addPinkClass(classes)
	time.Sleep(2 * time.Second)
	addGreenClass(classes)
	time.Sleep(2 * time.Second)
	addPinkClass(classes)
	time.Sleep(2 * time.Second)
	addGreenClass(classes)
	time.Sleep(2 * time.Second)
	addPinkClass(classes)
}
