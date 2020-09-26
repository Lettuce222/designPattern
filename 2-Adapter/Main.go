package main

import "./adapter"

func main() {
	var p adapter.Print = adapter.NewPrintBanner("Hello")
	p.PrintWeak()
	p.PrintStrong()
}
