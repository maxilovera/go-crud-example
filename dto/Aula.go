package dto

import (
	"time"

	"github.com/maxilovera/go-crud-example/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Aula struct {
	ID     int
	Nombre string
}

func (aula Aula) getObjectID() primitive.ObjectID {
	unixTime := time.Unix(int64(aula.ID), 0)

	return primitive.NewObjectIDFromTimestamp(unixTime)
}

func NewAula(aula model.Aula) *Aula {
	return &Aula{
		ID:     0,
		Nombre: aula.Nombre,
	}
}

func (aula Aula) GetModel() model.Aula {
	return model.Aula{
		ID:     aula.getObjectID(),
		Nombre: aula.Nombre,
	}
}
