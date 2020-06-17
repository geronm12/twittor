package models

import (
	jwt "github.com/drijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Claim ... es la etructura del jsonWebtoken que recibiremos*/
type Claim struct {
	Email string             `json:"email"`
	ID    primitive.ObjectID `bson:"_id" json:"_id, omitempty"`
	jwt.StandardClaims
}


