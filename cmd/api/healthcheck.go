package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status":      "available",
		"environment": app.cfg.env,
		"version":     version,
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
