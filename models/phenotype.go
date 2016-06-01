package models

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type SexChoice string

const (
	// Male                = "1"
	// Female              = "2"
	// Other               = "unknown"
	CollectionPhenotype = "phenotypes"
)

// Phenotype model
type Phenotype struct {
	Id              string `json:"id" `
	Name            string `json:"name"`
	Sex             string `json:"sex, omitempty"`
	Birth           string `json:"birth, omitempty"`
	Ageonset        string `json:"age_on_set, omitempty"`
	FamilyID        string `json:"family_id"`
	IndividualID    string `json:"individual_id"`
	PaternalID      string `json:"paternal_id"`
	MaternalID      string `json:"maternal_id"`
	AffectionStatus string `json:"affection_status, omitempty"`
}

// Create a phenotype
func CreatePhenotype(c *gin.Context) (*Phenotype, error) {
	db := c.MustGet("db").(*sql.DB)
	phenotype := new(Phenotype)
	phenotype.Name = c.PostForm("name")
	phenotype.Sex = c.PostForm("sex")
	phenotype.Birth = c.PostForm("birth")
	phenotype.Ageonset = c.PostForm("age_on_set")
	phenotype.FamilyID = c.PostForm("family_id")
	phenotype.IndividualID = c.PostForm("individual_id")
	phenotype.PaternalID = c.PostForm("paternal_id")
	phenotype.MaternalID = c.PostForm("maternal_id")
	phenotype.AffectionStatus = c.PostForm("affection_status")

	stmt, err := db.Prepare("INSERT INTO phenotypes(name,sex,birth,age_on_set,family_id,individual_id,paternal_id,maternal_id,affection_status) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?);")
	defer stmt.Close()
	stmt.Exec(&phenotype.Name, &phenotype.Sex, &phenotype.Birth, &phenotype.Ageonset, &phenotype.FamilyID, &phenotype.IndividualID, &phenotype.PaternalID, &phenotype.MaternalID, &phenotype.AffectionStatus)
	if err != nil {
		log.Print("createphenotypeerr: ", err)
		log.Print("createphenotype: ", stmt)
		return nil, err
	}
	log.Print("createphenotype: ", phenotype)
	return phenotype, nil
}

// Get a phenotype
func GetPhenotype(c *gin.Context) (*Phenotype, error) {

	db := c.MustGet("db").(*sql.DB)
	name := c.PostForm("name")
	phenotype := new(Phenotype)
	err := db.QueryRow("SELECT * FROM phenotypes WHERE name=?;", name).Scan(&phenotype.Id, &phenotype.Name, &phenotype.Sex, &phenotype.Birth, &phenotype.Ageonset, &phenotype.FamilyID, &phenotype.IndividualID, &phenotype.PaternalID, &phenotype.MaternalID, &phenotype.AffectionStatus)
	log.Print("getphenotype: ", phenotype)
	if err != nil {
		log.Print("getphenotypeerr: ", err)
		log.Print("getphenotypename: ", name)
		return nil, err
	}
	log.Print("getphenotype: ", phenotype)
	return phenotype, nil
}

// List all phenotypes
func ListPhenotypes(c *gin.Context) ([]*Phenotype, error) {
	db := c.MustGet("db").(*sql.DB)
	rows, err := db.Query("SELECT * FROM phenotypes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var phenotypes []*Phenotype

	for rows.Next() {
		phenotype := new(Phenotype)
		err := rows.Scan(&phenotype.Id, &phenotype.Name, &phenotype.Sex, &phenotype.Birth, &phenotype.Ageonset, &phenotype.FamilyID, &phenotype.IndividualID, &phenotype.PaternalID, &phenotype.MaternalID, &phenotype.AffectionStatus)
		if err != nil {
			log.Print("phenotypelist: ", err)
			return nil, err
		}
		phenotypes = append(phenotypes, phenotype)
	}
	if err = rows.Err(); err != nil {
		log.Print("phenotypeslist: ", err)
		return nil, err
	}
	return phenotypes, nil
}

// Update a phenotype
func UpdatePhenotype(c *gin.Context) (*Phenotype, error) {
	db := c.MustGet("db").(*sql.DB)
	phenotype := new(Phenotype)
	phenotype.Name = c.PostForm("name")
	phenotype.Sex = c.PostForm("sex")
	phenotype.Birth = c.PostForm("birth")
	phenotype.Ageonset = c.PostForm("age_on_set")
	phenotype.FamilyID = c.PostForm("family_id")
	phenotype.IndividualID = c.PostForm("individual_id")
	phenotype.PaternalID = c.PostForm("paternal_id")
	phenotype.MaternalID = c.PostForm("maternal_id")
	phenotype.AffectionStatus = c.PostForm("affection_status")
	stmt, err := db.Prepare("UPDATE phenotypes set sex=?, birth=?, age_on_set=?, family_id=?, individual_id=?, paternal_id=?, maternal_id=?, affection_status=?  WHERE name=? ;")
	defer stmt.Close()
	stmt.Exec(&phenotype.Sex, &phenotype.Birth, &phenotype.Ageonset, &phenotype.FamilyID, &phenotype.IndividualID, &phenotype.PaternalID, &phenotype.MaternalID, &phenotype.AffectionStatus, &phenotype.Name)

	if err != nil {
		log.Print("updaterr: ", err)
		return nil, err
	}

	log.Print("update: ", stmt)
	return phenotype, nil
}

// Delete a phenotype
func DeletePhenotype(c *gin.Context) (bool, error) {
	db := c.MustGet("db").(*sql.DB)
	name := c.PostForm("name")
	stmt, err := db.Exec("DELETE FROM phenotypes WHERE name=?;", name)
	if err != nil {
		log.Print("delete: ", err)
		return false, err
	}
	log.Print("deletetrue:", stmt)
	return true, nil
}

/*
CREATE TABLE `phenotypes` (
	`id` BIGINT NOT NULL AUTO_INCREMENT,
	`name` char(250) NOT NULL,
	`sex` varchar(20),
	`birth` varchar(100),
	`age_on_set` varchar(20),
	`family_id` varchar(255)  NOT NULL,
	`individual_id` varchar(255)  NOT NULL,
	`paternal_id` varchar(255),
	`maternal_id` varchar(255),
	`affection_status` varchar(60),
	unique(`name`),
    PRIMARY KEY (`id`)
);

*/
