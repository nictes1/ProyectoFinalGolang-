package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/nictes1/FinalProjectGO/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ConsultoRelacion - consulta relacion entre dos usuarios*/
func ConsultoRelacion(rel models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblog")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         rel.UsuarioID,
		"usuariorelacionid": rel.UsuarioRelacionID,
	}
	var resultado models.Relacion
	fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
