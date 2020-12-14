package services

import (
	"context"
	"library_manager/models"
	"library_manager/protos"
	"library_manager/mappers"
	"fmt"
)

type BookService struct {}

func (this *BookService) ListAll(ctx context.Context, req *protos.Empty, rsp *protos.BookList) error {
	var list []*models.Book
	query := new(models.Book).Query()
	count, _ := query.Count()
	if count > 0 {
		query.All(&list)
	}

	for _, book := range list {
		fmt.Println("book id = ", book.Id)
		rsp.Book = append(rsp.Book, mappers.BookToProto(book))
	}

	return nil
}
