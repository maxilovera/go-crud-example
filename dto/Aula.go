package dto

import (
	"github.com/maxilovera/go-crud-example/model"
	"github.com/maxilovera/go-crud-example/utils"
)

type Aula struct {
	ID     string
	Nombre string
}

func NewAula(aula model.Aula) *Aula {
	return &Aula{
		ID:     utils.GetStringIDFromObjectID(aula.ID),
		Nombre: aula.Nombre,
	}
}

func (aula Aula) GetModel() model.Aula {
	return model.Aula{
		ID:     utils.GetObjectIDFromStringID(aula.ID),
		Nombre: aula.Nombre,
	}
}
