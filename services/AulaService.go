package services

import (
	"github.com/maxilovera/go-crud-example/dto"
	"github.com/maxilovera/go-crud-example/repositories"
)

type AulaInterface interface {
	ObtenerAulas() []*dto.Aula
	InsertarAula(aula *dto.Aula) bool
}

type AulaService struct {
	aulaRepository repositories.AulaRepositoryInterface
}

func NewAulaService(aulaRepository repositories.AulaRepositoryInterface) *AulaService {
	return &AulaService{
		aulaRepository: aulaRepository,
	}
}

func (service *AulaService) ObtenerAulas() []*dto.Aula {
	aulasDB, _ := service.aulaRepository.ObtenerAulas()

	var aulas []*dto.Aula
	for _, aulaDB := range aulasDB {
		aula := dto.NewAula(aulaDB)
		aulas = append(aulas, aula)
	}

	return aulas
}

func (service *AulaService) InsertarAula(aula *dto.Aula) bool {
	service.aulaRepository.InsertarAula(aula.GetModel())

	return true
}
