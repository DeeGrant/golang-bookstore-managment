package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/DeeGrant/golang-bookstore-management/pkg/models"
	"github.com/DeeGrant/golang-bookstore-management/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetBooks()

	FinalResponse(w, books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book, _ := models.GetBook(Id)

	FinalResponse(w, book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	newBook := book.CreateBook()

	FinalResponse(w, newBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	deletedBook := models.DeleteBook(Id)

	FinalResponse(w, deletedBook)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book, db := models.GetBook(Id)

	updateBookData := &models.Book{}
	utils.ParseBody(r, updateBookData)

	if updateBookData.Name != "" {
		book.Name = updateBookData.Name
	}
	if updateBookData.Author != "" {
		book.Author = updateBookData.Author
	}
	if updateBookData.Publication != "" {
		book.Publication = updateBookData.Publication
	}
	db.Save(&book)

	FinalResponse(w, book)
}

func FinalResponse(w http.ResponseWriter, v any) {
	res, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
