package main

import "net/http"

func (app *application) HealthHandler(w http.ResponseWriter, r *http.Request) {
	app.logger.Info("request received", "uri", r.URL.RequestURI())

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}
