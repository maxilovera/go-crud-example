package services

import "github.com/maxilovera/go-crud-example/dto"

type AulaService struct {
}

func NewAulaService() *AulaService {
	return &AulaService{}
}

func (service AulaService) ObtenerAulas() []*dto.Aula {
	return []*dto.Aula{
		{ID: 1, Nombre: "Maxi"},
		{ID: 2, Nombre: "Mauro"},
		{ID: 3, Nombre: "Gonza"},
	}
}
