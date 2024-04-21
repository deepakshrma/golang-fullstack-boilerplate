package config

type Database struct {
}

func (db *Database) Query(query string) string {
	return `{"name": "MemDatabase"}`
}

// MemDatabase is some mock database
func MemDatabase(config *Config) *Database {
	Logger.Info(config.DBHost)
	return &Database{}
}
