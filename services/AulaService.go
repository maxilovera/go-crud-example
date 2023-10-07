package services

import (
	"log"

	"github.com/maxilovera/go-crud-example/dto"
	"github.com/maxilovera/go-crud-example/repositories"
	"github.com/maxilovera/go-crud-example/utils"
)

type AulaInterface interface {
	ObtenerAulas(filtroNombre string) []*dto.Aula
	ObtenerAulaPorID(id string) *dto.Aula
	EliminarAula(id string) bool
	InsertarAula(aula *dto.Aula) bool
	ModificarAula(aula *dto.Aula) bool
}

type AulaService struct {
	aulaRepository repositories.AulaRepositoryInterface
}

func NewAulaService(aulaRepository repositories.AulaRepositoryInterface) *AulaService {
	return &AulaService{
		aulaRepository: aulaRepository,
	}
}

func (service *AulaService) ObtenerAulas(filtroNombre string) []*dto.Aula {
	aulasDB, _ := service.aulaRepository.ObtenerAulas()

	var aulas []*dto.Aula
	for _, aulaDB := range aulasDB {
		aula := dto.NewAula(aulaDB)
		aulas = append(aulas, aula)
	}

	//Filtrar por filtroNombre

	return aulas
}

func (service *AulaService) ObtenerAulaPorID(id string) *dto.Aula {
	aulaDB, err := service.aulaRepository.ObtenerAulaPorID(id)

	var aula *dto.Aula
	if err == nil {
		aula = dto.NewAula(aulaDB)
	} else {
		log.Printf("[service:AulaService][method:ObtenerAulaPorId][reason:NOT_FOUND][id:%d]", id)
	}

	return aula
}

func (service *AulaService) InsertarAula(aula *dto.Aula) bool {
	service.aulaRepository.InsertarAula(aula.GetModel())

	return true
}

func (service *AulaService) ModificarAula(aula *dto.Aula) bool {
	service.aulaRepository.ModificarAula(aula.GetModel())

	return true
}

func (service *AulaService) EliminarAula(id string) bool {
	service.aulaRepository.EliminarAula(utils.GetObjectIDFromStringID(id))

	return true
}
