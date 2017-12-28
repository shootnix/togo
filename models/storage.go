package models

import (
	"database/sql"
	"errors"
	"os"
	"os/user"
	"path"

	_ "github.com/mattn/go-sqlite3" // sqlite3 interface
)

// Storage object
type Storage struct {
	HomeDir    string
	WorkDir    string
	DbFilePath string
	DbHandler  *sql.DB
}

var storage *Storage

// InitStorage : init storage object, working dir and database
func InitStorage() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	storage = &Storage{
		HomeDir:    usr.HomeDir,
		WorkDir:    path.Join(usr.HomeDir, ".goto"),
		DbFilePath: path.Join(usr.HomeDir, ".goto", "tasks.db"),
	}
	if _, err := os.Stat(storage.WorkDir); os.IsNotExist(err) {
		// create folder
		err := os.Mkdir(storage.WorkDir, 0755)
		raiseIfError(err)
	}

	storage.connectToDatabase()
	if _, err := os.Stat(storage.DbFilePath); os.IsNotExist(err) {
		storage.initDatabase()
	}
}

func (s *Storage) connectToDatabase() {
	db, err := sql.Open("sqlite3", s.DbFilePath)
	raiseIfError(err)
	if db == nil {
		raiseIfError(errors.New("db is nil"))
	}
	s.DbHandler = db
}

func (s *Storage) initDatabase() {
	sqlTable := `
	CREATE TABLE tasks (
		id        integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		task      text,
		created   timestamp NOT NULL DEFAULT CURRENT_DATE,
		resolved  timestamp NULL DEFAULT NULL,
		priority  integer NOT NULL DEFAULT 0
	);
	`
	_, err := s.DbHandler.Exec(sqlTable)
	raiseIfError(err)
}
