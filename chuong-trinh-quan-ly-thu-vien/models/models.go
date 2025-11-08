package models

type Book struct {
	Id         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Borrower struct {
	Id    string
	Name  string
	Email string
}
