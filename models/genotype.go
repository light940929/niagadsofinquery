package models

import "gopkg.in/mgo.v2/bson"

const (
	CollectionGenotype = "genotypes"
)

// Genotype model
type Genotype struct {
	Id         bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title      string        `json:"title" form:"title" binding:"required" bson:"title"`
	Chr        string        `json:"chr" form:"chr" binding:"required" bson:"chr"`
	Coordinate string        `json:"coordinate" form:"coordinate" binding:"required" bson:"coordinate"`
	VariantID  string        `json:"variant_id" form:"variantid" bson:"variant_id"`
	Location   string        `json:"location" form:"location" bson:"location"`
	Call       string        `json:"call" form:"call" bson:"call"`
	// User      bson.ObjectId `json:"user"`
}
