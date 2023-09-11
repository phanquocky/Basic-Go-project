package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	// "math/rand"
	"net/http"
	// "strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Movie struct {
	ID    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
	// Director *Director `json:"director"`
}

// type Director struct {
// 	Firstname string `json:"firstname"`
// 	Lastname  string `json:"lastname"`
// }

func getmovies(w http.ResponseWriter, r *http.Request) {
	db := connectDB()

	sqlStattement := `SELECT * FROM Movies`
	rows, err := db.Query(sqlStattement)
	if err != nil {
		fmt.Println("this1")
		fmt.Println(err)
		return
	}

	var movies []Movie
	for rows.Next() {
		var movie Movie
		rows.Scan(&movie.ID, &movie.Isbn, &movie.Title)
		movies = append(movies, movie)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)

	defer rows.Close()
	defer db.Close()
}

// func deletemovie(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "application/json")
// 	params := mux.Vars(r)

// 	for index, item := range movies {
// 		if item.ID == params["id"] {
// 			movies = append(movies[:index], movies[index+1:]...)
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(movies)
// }

// func getmovie(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "application/json")
// 	params := mux.Vars(r)

// 	for _, item := range movies {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// }

// func createmovie(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "application/json")
// 	var movie Movie
// 	_ = json.NewDecoder(r.Body).Decode(&movie)
// 	movie.ID = strconv.Itoa(rand.Intn(10000000))
// 	movies = append(movies, movie)

// 	json.NewEncoder(w).Encode(movies)
// }

// func updatemovie(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "application/json")
// 	params := mux.Vars(r)

// 	for index, item := range movies {
// 		if item.ID == params["id"] {
// 			movies = append(movies[:index], movies[index+1:]...)
// 			var movie Movie
// 			_ = json.NewDecoder(r.Body).Decode(&movie)
// 			movie.ID = strconv.Itoa(rand.Intn(10000000))
// 			movies = append(movies, movie)
// 			json.NewEncoder(w).Encode(movies)
// 			return
// 		}
// 	}
// }

// NOTICE: this info is in .env file
const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "quocky"
	dbname   = "postgres"
)

func connectDB() *sql.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", sqlInfo)
	if err != nil {
		fmt.Println("cannot connect to db")
		fmt.Println(err)
		return nil
	}

	// err = db.Ping()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil
	// }

	return db
}

func seekData() {
	db := connectDB()
	sqlStatement := `create table Movies(id varchar(40), isbn varchar(40) title varchar(40))`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		fmt.Println("cannot create table")
		fmt.Println(err)
		return
	}

	sqlInsertStatement := `insert into Movies(id, isbn, title) values ($1, $2)`
	_, err = db.Exec(sqlInsertStatement, "1", "1234", "movie1")
	if err != nil {
		fmt.Println("cannot insert movie")
	}
	_, err = db.Exec(sqlInsertStatement, "2", "4321", "movie2")
	if err != nil {
		fmt.Println("cannot insert movie")
	}

	defer db.Close()
}

func main() {
	seekData()
	r := mux.NewRouter()

	r.HandleFunc("/movies", getmovies).Methods("GET")
	// r.HandleFunc("/movies/{id}", getmovie).Methods("GET")
	// r.HandleFunc("/movies", createmovie).Methods("POST")
	// r.HandleFunc("/movies/{id}", updatemovie).Methods("PUT")
	// r.HandleFunc("/movies/{id}", deletemovie).Methods("DELETE")

	fmt.Println("Starting server at 8080 port .... ")
	log.Fatal(http.ListenAndServe(":8080", r))
}
