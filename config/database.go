package config

import (
	"encoding/json"
	"strings"
)

type Database struct {
}

type User struct {
	Age     int    `json:"age"`
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Address struct {
		Street string `json:"street"`
		City   string `json:"city"`
		State  string `json:"state"`
		Zip    int    `json:"zip"`
	} `json:"address"`
}

var users []User

func init() {
	data := `[
  {
    "age": 25,
    "name": "John Ward",
    "company": "PEARLESSA",
    "email": "stellaward@pearlessa.com",
    "address": {
      "street": "Bath Avenue, 705",
      "city": "Harrodsburg",
      "state": "Hawaii",
      "zip": 3276
    }
  },
  {
    "age": 40,
    "name": "Francis Molina",
    "company": "ECOLIGHT",
    "email": "francismolina@ecolight.com",
    "address": {
      "street": "Woodhull Street, 874",
      "city": "Brazos",
      "state": "Pennsylvania",
      "zip": 1567
    }
  },
  {
    "age": 30,
    "name": "Esther Hess",
    "company": "HYDROCOM",
    "email": "estherhess@hydrocom.com",
    "address": {
      "street": "Howard Avenue, 143",
      "city": "Germanton",
      "state": "Federated States Of Micronesia",
      "zip": 4233
    }
  }
]`
	json.Unmarshal([]byte(data), &users)
}

func (db *Database) Query(name string) []User {
	var filteredUser []User
	for _, user := range users {
		if strings.Contains(strings.ToLower(user.Name), strings.ToLower(name)) {
			filteredUser = append(filteredUser, user)
		}
	}
	return filteredUser
}

// MemDatabase is some mock database
func MemDatabase(config *Config) *Database {
	Logger.Info(config.DBHost)
	return &Database{}
}
