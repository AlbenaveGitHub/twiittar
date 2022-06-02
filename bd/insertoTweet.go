package bd

import (
	"context"
	"time"

	"github.com/AlbenaveGitHub/twiittar/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoTweet(t models.GraboTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twiittar")
	col := db.Collection("tweet")

	registro := bson.M{
		"userid":  t.UserID,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}
	if result, err := col.InsertOne(ctx, registro); err != nil {
		return "", false, err
	} else {
		objID, _ := result.InsertedID.(primitive.ObjectID)
		return objID.String(), true, nil
	}

}
