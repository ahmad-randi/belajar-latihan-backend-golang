package book

/*
Repository bertugas MENYIMPAN dan MENGAMBIL data

Repository:
- tidak validasi
- tidak mikir aturan bisnis
- hanya urusan DATA
*/
type Repository struct {
	books []Book // database sementara (in-memory)
}

/*
Save menyimpan buku baru
*/
func (r *Repository) Save(book Book) {
	r.books = append(r.books, book)
}

/*
FindAll mengembalikan semua buku
*/
func (r *Repository) FindAll() []Book {
	return r.books
}

/*
FindByID mencari buku berdasarkan ID
*/
func (r *Repository) FindByID(id int) (*Book, error) {
	for i := range r.books {
		if r.books[i].ID == id {
			return &r.books[i], nil
		}
	}
	return nil, ErrBookNotFound
}

/*
Update mengganti data buku
*/
func (r *Repository) Update(book Book) error {
	for i := range r.books {
		if r.books[i].ID == book.ID {
			r.books[i] = book
			return nil
		}
	}
	return ErrBookNotFound
}

/*
Delete menghapus buku berdasarkan ID
*/
func (r *Repository) Delete(id int) error {
	for i := range r.books {
		if r.books[i].ID == id {
			r.books = append(r.books[:i], r.books[i+1:]...)
			return nil
		}
	}
	return ErrBookNotFound
}
