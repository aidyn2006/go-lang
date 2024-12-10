package services

import (
	"awesomeProject/database"
	"awesomeProject/models"
	_ "gorm.io/gorm"
)

// Создание нового сервиса для книг
type BookService struct{}

// Создание новой книги
func (s *BookService) CreateBook(book *models.Book) error {
	result := database.DB.Create(book)
	return result.Error
}

// Получение всех книг
func (s *BookService) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	result := database.DB.Find(&books)
	return books, result.Error
}

// Получение книги по ID
func (s *BookService) GetBookByID(id uint) (models.Book, error) {
	var book models.Book
	result := database.DB.First(&book, id)
	return book, result.Error
}

// Обновление книги
func (s *BookService) UpdateBook(id uint, book *models.Book) error {
	var existingBook models.Book
	if err := database.DB.First(&existingBook, id).Error; err != nil {
		return err
	}
	// Обновляем поля книги
	existingBook.Title = book.Title
	existingBook.Author = book.Author
	existingBook.Year = book.Year

	result := database.DB.Save(&existingBook)
	return result.Error
}

// Удаление книги
func (s *BookService) DeleteBook(id uint) error {
	result := database.DB.Delete(&models.Book{}, id)
	return result.Error
}
