package library

import "fmt"

// Interface
type PrintTable interface {
	PrintInfo()
}

func Print(p PrintTable) {
	p.PrintInfo()
}

// Encapsulate fields (lowercase = private)
type Book struct {
	title  string
	author string
	pages  int
}

// Constructor
func NewBook(title, author string, pages int) *Book {
	return &Book{title, author, pages}
}

// Setter Methods
func (book *Book) SetTitle(title string) {
	book.title = title
}

func (book *Book) SetAuthor(author string) {
	book.author = author
}

func (book *Book) SetPages(pages int) {
	book.pages = pages
}

// Getter Methods
func (book *Book) Title() string {
	return book.title
}

func (book *Book) Author() string {
	return book.author
}

func (book *Book) Pages() int {
	return book.pages
}

func (book *Book) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n", book.Title(), book.Author(), book.Pages())
}
