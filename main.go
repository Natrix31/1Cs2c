package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"time"
)

type indexData struct {
	Customer string
	DBs      []struct {
		DBName     string
		DBType     string
		Status     string
		LastBackup time.Time
	}
}

var idxdata = &indexData{
	Customer: "John Doe",
	DBs: []struct {
		DBName     string
		DBType     string
		Status     string
		LastBackup time.Time
	}{
		{DBName: "db1", DBType: "MySQL", Status: "Online", LastBackup: time.Now()},
		{DBName: "db2", DBType: "PostgreSQL", Status: "Offline", LastBackup: time.Now().Add(-time.Hour * 2)},
		{DBName: "db3", DBType: "MongoDB", Status: "Online", LastBackup: time.Now().Add(-time.Minute * 30)},
	},
}

var tpl = template.Must(template.ParseFiles("index.html"))

func ViewIndexHandler(w http.ResponseWriter, _ *http.Request) {
	tpl.Execute(w, idxdata)
}

func ViewSettingsHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the Settings Page!"))
}

func main() {

	http.HandleFunc("/", ViewIndexHandler)
	http.HandleFunc("/settings", ViewSettingsHandler)

	slog.Info("Starting server on :8181...")
	http.ListenAndServe(":8181", nil)
}
