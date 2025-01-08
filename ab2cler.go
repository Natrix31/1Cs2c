package ab2cler

import (
	"html/template"
	"log/slog"
	"net/http"
	"time"
	config "github.com/Natrix31/internal/config"
)

type CustomerDB struct {
	DBName     string
	DBType     string
	Status     string
	LastBackup time.Time
}

type indexData struct {
	Customer  string
	Databases []CustomerDB
}

//Это пока просто тестовые данные. Тут должен быть запрос в БД и получение реальных данных.
var idxdata = &indexData{
	Customer: "Зори белогорья",
	Databases: []CustomerDB{
		{DBName: "Палесика_ТП7", DBType: "file", Status: "Online", LastBackup: time.Now()},
		{DBName: "Новикевич_ТП7", DBType: "sql", Status: "Offline", LastBackup: time.Now().Add(-time.Hour * 2)},
		{DBName: "Новикевич_Бух", DBType: "sql", Status: "Online", LastBackup: time.Now().Add(-time.Minute * 30)},
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

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	slog.Info("Starting server on :8181...")
	conf, err := config.
	http.ListenAndServe(":8181", nil)
}
