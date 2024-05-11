package main

import (
	"net/http"
)

func (app *application) allUsers(w http.ResponseWriter, r *http.Request) {
	users, err := app.db.AllUsers()
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, users)
}

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {
	//userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	//if err != nil {
	//	app.errorJSON(w, err, http.StatusBadRequest)
	//	return
	//}

	//user, err := app.db.GetUser(userID)
	//if err != nil {
	//	app.errorJSON(w, err, http.StatusBadRequest)
	//	return
	//}

	_ = app.writeJSON(w, http.StatusOK, "user")
}
