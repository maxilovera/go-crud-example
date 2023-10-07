package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Aula struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"` //omitempty agrega el campo solo si no es nil
	Nombre string             `bson:"nombre"`
}
