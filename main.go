package main

// Import dependencies
import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Tutorial struct {
	Id      int
	Title   string
	Slug    string
	Content string
}

/* Slice tutorial */
var tutorials = []Tutorial{
	Tutorial{Id: 1, Title: "Operator Golang", Slug: "operator-golang", Content: "Belajar operator di Golang"},
	Tutorial{Id: 2, Title: "Kondisional Golang", Slug: "kondisional-golang", Content: "Belajar kondisional di Golang"},
	Tutorial{Id: 3, Title: "Array Golang", Slug: "array-golang", Content: "Belajar array di Golang"},
	Tutorial{Id: 4, Title: "Looping Golang", Slug: "looping-golang", Content: "Belajar looping di Golang"},
	Tutorial{Id: 5, Title: "Tipe Data Golang", Slug: "tipe-data-golang", Content: "Belajar tipe data di Golang"},
	Tutorial{Id: 6, Title: "Slice Golang", Slug: "slice-golang", Content: "Belajar slice di Golang"},
}

func main() {
	// Inisialisasi gorilla/mux router
	r := mux.NewRouter()

	// Halaman index
	r.Handle("/", http.FileServer(http.Dir("./views/")))

	// Menampilkan status API
	r.Handle("/status", StatusHandler).Methods("GET")

	// Menampilkan list tutorial
	r.Handle("/tutorials", TutorialsHandler).Methods("GET")

	// Deklarasi port yang digunakan
	http.ListenAndServe(":3000", r)
}

// Untuk menampilkan status API
var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})

// Untuk menampilkan semua tutorial
var TutorialsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// convert slice ke json
	payload, _ := json.Marshal(tutorials)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})
