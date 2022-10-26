package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
  "github.com/ramialkaro/books/models"
)

var books []models.Book


func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get a Book
func getBook(w http.ResponseWriter, r *http.Request) {

	// GET params
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Book{})

}

// Create a books
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000)) // Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update a Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// GET params
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book models.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = strconv.Itoa(rand.Intn(1000000)) // Mock ID - not safe
			books = append(books, book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Delete a Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// GET params
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Init Router
	r := mux.NewRouter().StrictSlash(true)

	// Mock Data TODO implement DB
	books = append(books, models.Book{ID: "1", Isbn: "7933", Title: "Book 1", Author: &models.Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, models.Book{ID: "2", Isbn: "7934", Title: "Book 2", Author: &models.Author{Firstname: "Steve", Lastname: "Harvey"}})
	books = append(books, models.Book{ID: "3", Isbn: "7936", Title: "Book 3", Author: &models.Author{Firstname: "Micheal", Lastname: "Smith"}})

	// Group routes
	sub := r.PathPrefix("/api/v1/books").Subrouter()
	// Route Handlers / Endpoints
	sub.HandleFunc("/", getBooks).Methods("GET")
	sub.HandleFunc("/{id}", getBook).Methods("GET")
	sub.HandleFunc("/", createBook).Methods("POST")
	sub.HandleFunc("/{id}", updateBook).Methods("PUT")
	sub.HandleFunc("/{id}", deleteBook).Methods("DELETE")
  log.Println("listen on port 8080")
  log.Fatal(http.ListenAndServe(":8000", r))
}
