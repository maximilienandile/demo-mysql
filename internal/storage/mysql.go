package storage

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/maximilienandile/demo-mysql/internal/book"
)

type MysqlStorage struct {
	db *sql.DB
}

type MysqlConfig struct {
	Username string
	Password string
	DbName   string
	Port     uint
	Host     string
}

func NewMysqlStorage(conf MysqlConfig) (MysqlStorage, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.DbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return MysqlStorage{}, fmt.Errorf("impossible to open SQL connexion: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return MysqlStorage{}, fmt.Errorf("impossible to ping db: %w", err)
	}
	return MysqlStorage{
		db: db,
	}, nil
}

func (s MysqlStorage) Create(ctx context.Context, b book.Book) (book.Book, error) {
	// TODO: check the input
	query := "INSERT INTO `book` (`create_time`, `name`, `author_name`) VALUES (?, ?, ?)"
	insertResult, err := s.db.ExecContext(ctx, query, b.CreateTime, b.Name, b.AuthorName)
	if err != nil {
		return b, fmt.Errorf("error while insert: %w", err)
	}
	id, err := insertResult.LastInsertId()
	if err != nil {
		return b, fmt.Errorf("error while get last insert id: %w", err)
	}
	b.ID = id

	return b, nil
}
