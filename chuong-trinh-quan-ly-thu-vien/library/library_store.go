package library

import "mamba.com/chuong-trinh-quan-ly-thu-vien/models"

type Library struct {
	book map[string]models.Book
}

func NewLibrary() *Library {
	return &Library{
		book: make(map[string]models.Book),
	}
}
