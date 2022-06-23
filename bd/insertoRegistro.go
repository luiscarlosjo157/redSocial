package bd

import (
	"context"
	"time"

	"github.com/luiscarlosjo157/redSocial/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRegistro es la parada final con la B para insetar los datos del usuario*/
func InsertoRegistro(u models.Usuario) (string, bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	OdjID, _ := result.InsertedID.(primitive.ObjectID)
	return OdjID.String(), true, nil
}