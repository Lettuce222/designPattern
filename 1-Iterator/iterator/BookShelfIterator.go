package iterator

type BookShelfIterator struct {
	bookshelf BookShelf
	index     int
}

func (bsi *BookShelfIterator) HasNext() bool {
	return bsi.index < bsi.bookshelf.Length()
}

func (bsi *BookShelfIterator) Next() interface{} {
	var book Book = bsi.bookshelf.BookAt(bsi.index)
	bsi.index++
	return book
}
