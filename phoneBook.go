package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/create/{key}", createContact).Methods("POST") //key?/{id:[0-9]+}"
	r.HandleFunc("/contacts", viewContacts).Methods("GET")
	r.HandleFunc("/contact/{key}", viewContact).Methods("GET")
	r.HandleFunc("/edit/{key}", editContact).Methods("POST")
	r.HandleFunc("/delete/{key}", deleteContact).Methods("POST")

	srv := http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

type Contact struct {
	name   string `json "name"`
	number string `json "number"`
	id     string `json "id"`
}

func createContact(w http.ResponseWriter, r *http.Request) {
	//url := r.GetBody
	contact := Contact{"Ashly", "+79063270910", "id0001"}
	db, err := sql.Open("postgres", "postgresql://phoneBook:123456@localhost:5432/phoneBook")
	if err != nil {
		return
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO contacts (name, number, id) VALUES ($1, $2, $3)", contact.name, contact.number, contact.id)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "Ok")
}
func viewContacts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ok")
}
func viewContact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ok")
}
func editContact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ok")
}
func deleteContact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ok")
}
