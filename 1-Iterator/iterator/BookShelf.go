package iterator

type BookShelf struct {
	books []Book
}

func (bs *BookShelf) BookAt(index int) Book {
	return bs.books[index]
}

func (bs *BookShelf) AppendBook(book Book) {
	bs.books = append(bs.books, book)
}

func (bs *BookShelf) Length() int {
	return len(bs.books)
}

func (bs BookShelf) Iterator() Iterator {
	return &BookShelfIterator{bs, 0}
}
