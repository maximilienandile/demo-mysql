package book

import "time"

type Book struct {
	// TODO : use a string, populate with an UUID
	ID         int64
	Name       string
	AuthorName string
	CreateTime time.Time
}
