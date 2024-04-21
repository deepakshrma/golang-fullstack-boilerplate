package controller

import (
	"boilerplate/config"
	"boilerplate/route/middleware"
	"encoding/json"
	"net/http"
	"regexp"
)

type UserController struct {
	db *config.Database
}

var reg, _ = regexp.Compile(`/users/(.+)`)

func (c *UserController) getUserByName(w http.ResponseWriter, r *http.Request) {
	m := reg.FindStringSubmatch(r.URL.Path)
	users := c.db.Query(m[1])
	json.NewEncoder(w).Encode(users)
}

func (c *UserController) createUser(w http.ResponseWriter, r *http.Request) {

}

func (c *UserController) Users(w http.ResponseWriter, r *http.Request) {
	ctx := middleware.UseContext(r)
	w.Header().Add("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		c.getUserByName(w, r)
	case http.MethodPost:
		c.createUser(w, r)
	default:
		ctx.Log.Error("Invalid Operation Users", "method", r.Method)
	}
}

func NewUsersHandler(db *config.Database) *UserController {
	return &UserController{
		db: db,
	}
}
