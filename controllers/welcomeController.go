package controllers

import (
	"context"
	"library_manager/models"
	"library_manager/protos"
	"library_manager/mappers"
)

type WelcomeController struct {
	baseController
}

func (this *WelcomeController) ListAll() {
	rsp, err := bookServiceClent.ListAll(context.Background(), &protos.Empty{})
    if err != nil {
        panic(err)
	}

	bookList := make([]*models.Book, 0)
	for _, book := range rsp.Book {
		bookList = append(bookList, mappers.BookToModel(book))
	}
	
	this.Data["json"] = bookList
	this.ServeJSON()
}
