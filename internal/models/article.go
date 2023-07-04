package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Article struct {
	Id            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Aid           int                `bson:"aid" json:"aid"`
	Link          string             `bson:"link" json:"link"`
	Title         string             `bson:"title" json:"title"`
	SubscribedIDs []string           `bson:"subscribed_ids" json:"subscribed_ids"`
	CreatedAt     string             `json:"created_at"`
	UpdatedAt     string             `json:"updated_at"`
}
