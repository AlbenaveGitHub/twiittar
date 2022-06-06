package bd

import (
	"context"
	"time"

	"github.com/AlbenaveGitHub/twiittar/models"
)

/*InsertoRelacion graba la relacion en la BD */
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twiittar")
	col := db.Collection("relacion")

	if _, err := col.InsertOne(ctx, t); err != nil {
		return false, err
	} else {
		return true, nil
	}
}
