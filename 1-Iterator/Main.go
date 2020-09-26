package main

import (
	"./iterator"
	"fmt"
)

func main() {
	var bs iterator.BookShelf = iterator.BookShelf{}
	bs.AppendBook(iterator.Book{"A"})
	bs.AppendBook(iterator.Book{"B"})
	bs.AppendBook(iterator.Book{"C"})
	bs.AppendBook(iterator.Book{"D"})

	var it iterator.Iterator = bs.Iterator()
	for it.HasNext() {
		var book interface{} = it.Next()
		switch book.(type) {
		case iterator.Book:
			fmt.Println(book)
		default:
			fmt.Printf("book is unknown type. [valueType: %T]\n", it)
		}
	}
}
