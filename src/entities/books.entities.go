package models

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	Genre  string `json:"genre"`
}


var Books = []Book{
	{ID: 1, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Year: 1925, Genre: "Fiction"},
	{ID: 2, Title: "The Da Vinci Code", Author: "Dan Brown", Year: 2003, Genre: "Mystery"},
}
