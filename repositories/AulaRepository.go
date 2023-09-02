package repositories

import (
	"context"
	"fmt"

	"github.com/maxilovera/go-crud-example/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AulaRepositoryInterface interface {
	ObtenerAulas() ([]model.Aula, error)
	InsertarAula(aula model.Aula) (*mongo.InsertOneResult, error)
}

type AulaRepository struct {
	db DB
}

func NewAulaRepository(db DB) *AulaRepository {
	return &AulaRepository{
		db: db,
	}
}

func (repository AulaRepository) ObtenerAulas() ([]model.Aula, error) {
	collection := repository.db.GetClient().Database("ejemplo").Collection("aulas")
	filtro := bson.M{}

	cursor, err := collection.Find(context.TODO(), filtro)

	defer cursor.Close(context.Background())

	// Itera a través de los resultados
	var aulas []model.Aula
	for cursor.Next(context.Background()) {
		var aula model.Aula
		err := cursor.Decode(&aula)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		aulas = append(aulas, aula)
	}

	return aulas, err
}

func (repository AulaRepository) InsertarAula(aula model.Aula) (*mongo.InsertOneResult, error) {
	collection := repository.db.GetClient().Database("ejemplo").Collection("aulas")
	resultado, err := collection.InsertOne(context.TODO(), aula)
	return resultado, err
}
