package models

import "gopkg.in/mgo.v2/bson"

type SexChoice string

const (
	Male                = "1"
	Female              = "2"
	Other               = "unknown"
	CollectionPhenotype = "phenotypes"
)

// Phenotype model
type Phenotype struct {
	Id              bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title           string        `json:"title" form:"title" binding:"required" bson:"title"`
	Sex             SexChoice     `json:"sex" form:"sex" bson:"sex"`
	Birth           int64         `json:"birth" form:"birth" bson:"birth"`
	Ageonset        int64         `json:"age_on_set"  form:"ageonset" bson:"age_on_set"`
	FamilyID        string        `json:"family_id" form:"familyid" binding:"required" bson:"family_id"`
	IndividualID    string        `json:"individual_id" form:"individualid" binding:"required" bson:"individual_id"`
	PaternalID      string        `json:"paternal_id" form:"paternalid" bson:"paternal_id"`
	MaternalID      string        `json:"maternal_id" form:"maternalid" bson:"maternal_id"`
	AffectionStatus string        `json:"affection_status" form:"affectionstatus" bson:"affection_status"`
	// User      bson.ObjectId `json:"user"`
}
