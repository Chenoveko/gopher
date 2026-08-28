package library

type TextBook struct {
	Book
	publisher string
	level     string
}

// Constructor
func NewTextBook(title, author string, pages int, publisher, level string) *TextBook {
	return &TextBook{
		Book{title, author, pages},
		publisher,
		level,
	}
}

// Setter Methods
func (textBook *TextBook) SetPublisher(publisher string) {
	textBook.publisher = publisher
}

func (textBook *TextBook) SetLevel(level string) {
	textBook.level = level
}

// Getter Methods
func (textBook *TextBook) Publisher() string {
	return textBook.publisher
}

func (textBook *TextBook) Level() string {
	return textBook.level
}
