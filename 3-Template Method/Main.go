package main

import (
	"./template"
)

func main() {
	d1 := template.NewAbstractDisplay(*template.NewCharOPC('H'))
	d2 := template.NewAbstractDisplay(*template.NewStringOPC("Hello, World"))
	d3 := template.NewAbstractDisplay(*template.NewStringOPC("こんにちは。"))

	d1.Display()
	d2.Display()
	d3.Display()
}
