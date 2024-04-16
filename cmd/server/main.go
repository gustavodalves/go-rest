package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gustavodalves/go-api/internal/application"
	"github.com/gustavodalves/go-api/internal/database"
	"github.com/gustavodalves/go-api/internal/web"
)

func main() {
	mux := mux.NewRouter()

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3309)/test")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	database := database.NewUserDb(db)

	service := application.NewUserService(database)

	handler := web.NewUserHandler(service)

	mux.HandleFunc("/user", handler.Get).Methods("GET")
	mux.HandleFunc("/user/{id}", handler.GetUnique).Methods("GET")
	mux.HandleFunc("/user", handler.Post).Methods("POST")

	http.ListenAndServe(":8080", mux)
}
