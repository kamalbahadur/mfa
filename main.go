package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"mfa/handler"
	"mfa/persistence"
	"net/http"
)

var db *sql.DB

func main() {
	db = persistence.NewDB()
	router := mux.NewRouter()

	// Repo
	repo := persistence.NewRepository(db)

	// Handlers
	signupHandler := handler.NewSignupHandler(repo)
	verifyHandler := handler.NewVerifyHandler(repo)

	// Routes
	mfaRouter := router.PathPrefix("/mfa").Subrouter()
	mfaRouter.Handle("/signup", signupHandler).Methods(http.MethodGet)
	mfaRouter.Handle("/verify", verifyHandler).Methods(http.MethodGet)

	// Start server
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("There's an error with the server")
	}

}
