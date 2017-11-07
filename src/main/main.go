package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/bmizerany/pq"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "deb.deb"
	dbname   = "golang_demo"
)

var people []Person

func main() {
	//	http.HandleFunc("/api/hello", func(w http.ResponseWriter, req *http.Request) {
	//		w.Write([]byte("Hello World!!"))
	//	})
	//
	//	http.ListenAndServe(":9090", nil)

	//===============================================

	router := mux.NewRouter()

	router.HandleFunc("/api/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/api/people", CreatePerson).Methods("POST")

	log.Fatal(http.ListenAndServe(":9090", router))

}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)

	result := save(person)
	json.NewEncoder(w).Encode(result)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	inputId := params["id"]

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, dbOpeningError := sql.Open("postgres", psqlInfo)

	if dbOpeningError != nil {
		panic(dbOpeningError)
	}
	defer db.Close()

	sqlStatement := "SELECT id, first_name, last_name from users where id = $1"

	id := ""
	first_name := ""
	last_name := ""

	errInsert := db.QueryRow(sqlStatement, inputId).Scan(&id, &first_name, &last_name)

	if errInsert != nil {
		panic(errInsert)
	}

	json.NewEncoder(w).Encode(&Person{})
}

func save(person Person) (result Person) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, dbOpeningError := sql.Open("postgres", psqlInfo)

	if dbOpeningError != nil {
		panic(dbOpeningError)
	}
	defer db.Close()

	sqlStatement := "INSERT INTO users (id, first_name, last_name)  VALUES ($1, $2, $3) RETURNING id , first_name, last_name"

	id := time.Now().Format("20060102150405")
	first_name := ""
	last_name := ""

	errInsert := db.QueryRow(sqlStatement, id, person.Firstname, person.Lastname).Scan(&id, &first_name, &last_name)

	if errInsert != nil {
		panic(errInsert)
	}

	result = Person{id, first_name, last_name}

	return result
}

type Person struct {
	ID        string //`json:"id,omitempty"`
	Firstname string //`json:"firstname,omitempty"`
	Lastname  string //`json:"lastname,omitempty"`
}
