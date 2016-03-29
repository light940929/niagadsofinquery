package models

import "gopkg.in/mgo.v2/bson"

type Format string

const (
	FASTQ               = "FASTQ"
	SAM                 = "SAM"
	BAM                 = "BAM"
	VCF                 = "VCF"
	BED                 = "BED"
	TXT                 = "TXT"
	PED                 = "PED"
	MAP                 = "MAP"
	FAM                 = "FAM"
	OTHER               = "OTHER"
	CollectionReference = "references"
)

// Reference model
type Reference struct {
	Id          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string        `json:"name" form:"name" binding:"required" bson:"name"`
	Description string        `json:"description" form:"description" binding:"required" bson:"description"`
	Plateform   string        `json:"plateform"  bson:"plateform"`
	Length      string        `json:"length" bson:"length"`
	SourceUri   string        `json:"source_uri" bson:"sourceuri"`
	Format      Format        `json:"format" bson:"format"`
	// User      bson.ObjectId `json:"user"`
}
