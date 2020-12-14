package mappers

import (
	"library_manager/models"
	"library_manager/protos"
)

func BookToProto(model *models.Book) (proto *protos.Book) {
	proto = new(protos.Book)
	proto.Id = model.Id
	proto.Name = model.Name
	proto.Author = model.Author
	return
}

func BookToModel(proto *protos.Book) (model *models.Book) {
	model = new(models.Book)
	model.Id = proto.Id
	model.Name = proto.Name
	model.Author = proto.Author
	return
}
