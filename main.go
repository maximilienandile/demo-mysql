package main

import (
	"context"
	"github.com/maximilienandile/demo-mysql/internal/book"
	"github.com/maximilienandile/demo-mysql/internal/storage"
	"log"
	"time"
)

func main() {
	store, err := storage.NewMysqlStorage(storage.MysqlConfig{
		Username: "user",
		Password: "password",
		DbName:   "db",
		Port:     3306,
		Host:     "localhost",
	})
	if err != nil {
		log.Fatalf("impossible to create mysql storage: %s", err)
	}
	b, err := store.Create(context.Background(), book.Book{
		Name:       "Practical Go Lessons",
		AuthorName: "Maximilien Andile",
		CreateTime: time.Now(),
	})
	if err != nil {
		log.Fatalf("impossible to insert: %s", err)
	}
	log.Println(b)
}
