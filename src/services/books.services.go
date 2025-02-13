package services

import "go-library-api/src/entities"

type BookService struct{}

// GetAllBooks retrieves all books with optional filters
func (s *BookService) GetAllBooks(author, genre, title, order string) ([]models.Book, error) {
	// Method not implemented yet
	return nil, nil
}

// GetSortedBooks retrieves books sorted by the specified order
func (s *BookService) GetSortedBooks(order string) ([]models.Book, error) {
	// Method not implemented yet
	return nil, nil
}

// GetBookById retrieves a book by its ID
func (s *BookService) GetBookById(id int) (models.Book, error) {
	// Method not implemented yet
	return models.Book{}, nil
}

// CreateBook creates a new book
func (s *BookService) CreateBook(book models.Book) (models.Book, error) {
	// Method not implemented yet
	return models.Book{}, nil
}

// UpdateBook updates an existing book
func (s *BookService) UpdateBook(id int, book models.Book) (models.Book, error) {
	// Method not implemented yet
	return models.Book{}, nil
}

// DeleteBook deletes a book by its ID
func (s *BookService) DeleteBook(id int) error {
	// Method not implemented yet
	return nil
}
