package config

import "fmt"

type Database struct {
}

func (db *Database) Query(query string) string {
	return `{"name": "MemDatabase"}`
}

// MemDatabase is some mock database
func MemDatabase(config *Config) *Database {
	fmt.Println(config.DBHost)
	return &Database{}
}
