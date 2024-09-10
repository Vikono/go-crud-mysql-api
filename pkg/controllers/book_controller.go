package controllers

import (
	"encoding/json"
	"main/pkg/models"
	"main/pkg/utils"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	bookList := models.GetAllBooks()
	res, _ := json.Marshal(bookList)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	b := book.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// func GetBook(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	bookId := vars["bookId"]
// }
