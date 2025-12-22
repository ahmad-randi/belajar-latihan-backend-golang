package book

import (
	"buku-project/internal/book/dto"
	"time"
)

/*
Service adalah OTAK APLIKASI

Semua aturan bisnis ada di sini
*/
var repo = &Repository{}

/*
CreateBook menambah buku baru
*/
func CreateBook(req dto.CreateBookRequest) error {
	if err := ValidateCreateBook(req); err != nil {
		return err
	}

	now := time.Now()

	book := Book{
		ID:          GenerateBookID(),
		Title:       req.Title,
		Author:      req.Author,
		Description: req.Description,
		IsBorrowed:  false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	repo.Save(book)
	return nil
}

/*
GetAllBooks menampilkan semua buku
*/
func GetAllBooks() []Book {
	return repo.FindAll()
}

/*
UpdateBook mengubah buku
*/
func UpdateBook(req dto.UpdateBookRequest) error {
	if err := ValidateUpdateBook(req); err != nil {
		return err
	}

	book, err := repo.FindByID(req.ID)
	if err != nil {
		return err
	}

	if book.IsBorrowed {
		return ErrBookUpdateAlreadyBorrowed
	}

	book.Title = req.Title
	book.Author = req.Author
	book.Description = req.Description
	book.UpdatedAt = time.Now()

	return repo.Update(*book)
}

/*
BorrowBook meminjam buku
*/
func BorrowBook(id int) error {
	book, err := repo.FindByID(id)
	if err != nil {
		return err
	}
	if book.IsBorrowed {
		return ErrBookAlreadyBorrowed
	}

	book.IsBorrowed = true
	book.UpdatedAt = time.Now()
	return repo.Update(*book)
}

/*
ReturnBook mengembalikan buku
*/
func ReturnBook(id int) error {
	book, err := repo.FindByID(id)
	if err != nil {
		return err
	}
	if !book.IsBorrowed {
		return ErrBookNotBorrowed
	}

	book.IsBorrowed = false
	book.UpdatedAt = time.Now()
	return repo.Update(*book)
}

/*
DeleteBook menghapus buku
*/
func DeleteBook(id int) error {
	book, err := repo.FindByID(id)
	if err != nil {
		return err
	}
	if book.IsBorrowed {
		return ErrBookDeleteAlreadyBorrowed
	}
	return repo.Delete(id)
}
