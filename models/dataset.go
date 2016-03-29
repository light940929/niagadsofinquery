package models

import "gopkg.in/mgo.v2/bson"

type Type string

const (
	VARIANTS          = "variants"
	ANNOTATIONS       = "annotations"
	CollectionDataset = "datasets"
)

// Dataset model
type Dataset struct {
	Id          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string        `json:"name" form:"name" binding:"required" bson:"name"`
	Description string        `json:"description" form:"description" binding:"required" bson:"description"`
	Type        Type          `json:"type" binding:"required" bson:"type"`
	Created     int64         `json:"created" bson:"created"`
	// User      bson.ObjectId `json:"user"`
}
