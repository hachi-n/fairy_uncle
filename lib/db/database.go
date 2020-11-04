package db

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

const (
	databaseDir      = ".config/fairy_uncle"
	databaseFileName = "fairy_uncle.db"
)

var DB *Database

type Database struct {
	Data []byte
	mu   sync.Mutex
}

func init() {
	checkDatabase()
}

func databaseFileDir() string {
	return filepath.Join(os.Getenv("HOME"), databaseDir)
}

func databaseFilePath() string {
	return filepath.Join(databaseFileDir(), databaseFileName)
}

func checkDatabase() {
	dbfp := databaseFilePath()
	_, err := os.Stat(dbfp)
	if err != nil {
		os.MkdirAll(databaseFileDir(), 0755)
		_, err := os.OpenFile(dbfp, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Printf("no database error: Please check your permission %s\n", dbfp)
			os.Exit(1)
		}
	}
}

func LoadDatabase() (*Database, error) {
	b, err := ioutil.ReadFile(databaseFilePath())
	if err != nil {
		return nil, err
	}

	DB = &Database{Data: b}
	return DB, nil
}

func (db *Database) Fetch() ([]byte, error) {
	if db.Data != nil {
		return db.Data, nil
	}

	b, err := ioutil.ReadFile(databaseFilePath())
	if err != nil {
		return nil, err
	}

	db.Data = b
	return db.Data, nil
}

func (db *Database) Store(data []byte) error {
	db.mu.Lock()
	db.Data = data
	db.mu.Unlock()

	return ioutil.WriteFile(databaseFilePath(), db.Data, 0666)
}
