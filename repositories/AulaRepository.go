package repositories

import (
	"context"
	"fmt"

	"github.com/maxilovera/go-crud-example/model"
	"github.com/maxilovera/go-crud-example/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AulaRepositoryInterface interface {
	ObtenerAulas() ([]model.Aula, error)
	ObtenerAulaPorID(id string) (model.Aula, error)
	EliminarAula(id primitive.ObjectID) (*mongo.DeleteResult, error)
	InsertarAula(aula model.Aula) (*mongo.InsertOneResult, error)
	ModificarAula(aula model.Aula) (*mongo.UpdateResult, error)
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

func (repository AulaRepository) ObtenerAulaPorID(id string) (model.Aula, error) {
	collection := repository.db.GetClient().Database("ejemplo").Collection("aulas")
	objectID := utils.GetObjectIDFromStringID(id)
	filtro := bson.M{"_id": objectID}

	cursor, err := collection.Find(context.TODO(), filtro)

	defer cursor.Close(context.Background())

	// Itera a través de los resultados
	var aula model.Aula

	for cursor.Next(context.Background()) {
		err := cursor.Decode(&aula)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}

	return aula, err
}

func (repository AulaRepository) InsertarAula(aula model.Aula) (*mongo.InsertOneResult, error) {
	collection := repository.db.GetClient().Database("ejemplo").Collection("aulas")
	resultado, err := collection.InsertOne(context.TODO(), aula)
	return resultado, err
}

func (repository AulaRepository) ModificarAula(aula model.Aula) (*mongo.UpdateResult, error) {
	collection := repository.db.GetClient().Database("ejemplo").Collection("aulas")

	filtro := bson.M{"_id": aula.ID}
	entidad := bson.M{"$set": bson.M{"nombre": aula.Nombre}}

	resultado, err := collection.UpdateOne(context.TODO(), filtro, entidad)

	return resultado, err
}

func (repository AulaRepository) EliminarAula(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	collection := repository.db.GetClient().Database("ejemplo").Collection("aulas")

	filtro := bson.M{"_id": id}

	resultado, err := collection.DeleteOne(context.TODO(), filtro)

	return resultado, err
}
