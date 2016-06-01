package models

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const (
	CollectionGenotype = "genotypes"
)

// Genotype model
type Genotype struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Chr        string `json:"chr"`
	Coordinate string `json:"coordinate, omitempty"`
	VariantID  string `json:"variant_id"`
	Location   string `json:"location, omitempty"`
	Calls      string `json:"calls, omitempty"`
}

// Create a genotype
func CreateGenotype(c *gin.Context) (*Genotype, error) {
	db := c.MustGet("db").(*sql.DB)
	genotype := new(Genotype)
	genotype.Name = c.PostForm("name")
	genotype.Chr = c.PostForm("chr")
	genotype.Coordinate = c.PostForm("coordinate")
	genotype.VariantID = c.PostForm("variant_id")
	genotype.Location = c.PostForm("location")
	genotype.Calls = c.PostForm("calls")

	stmt, err := db.Prepare("INSERT INTO genotypes(name,chr,coordinate,variant_id,location,calls) VALUES(?, ?, ?, ?, ?, ?);")
	defer stmt.Close()
	stmt.Exec(&genotype.Name, &genotype.Chr, &genotype.Coordinate, &genotype.VariantID, &genotype.Location, &genotype.Calls)
	if err != nil {
		log.Print("creategenotypeerr: ", err)
		log.Print("creategenotype: ", stmt)
		return nil, err
	}
	log.Print("creategenotypet: ", genotype)
	return genotype, nil
}

// Get a genotype
func GetGenotype(c *gin.Context) (*Genotype, error) {

	db := c.MustGet("db").(*sql.DB)
	name := c.PostForm("name")
	genotype := new(Genotype)
	err := db.QueryRow("SELECT * FROM genotypes WHERE name=?;", name).Scan(&genotype.Id, &genotype.Name, &genotype.Chr, &genotype.Coordinate, &genotype.VariantID, &genotype.Location, &genotype.Calls)
	log.Print("getgenotype: ", genotype)
	if err != nil {
		log.Print("getgenotyperr: ", err)
		log.Print("getgenotypename: ", name)
		return nil, err
	}
	return genotype, nil
}

// List all genotypes
func ListGenotypes(c *gin.Context) ([]*Genotype, error) {
	db := c.MustGet("db").(*sql.DB)
	rows, err := db.Query("SELECT * FROM genotypes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genotypes []*Genotype

	for rows.Next() {
		genotype := new(Genotype)
		err := rows.Scan(&genotype.Id, &genotype.Name, &genotype.Chr, &genotype.Coordinate, &genotype.VariantID, &genotype.Location, &genotype.Calls)
		if err != nil {
			log.Print("genotypelist: ", err)
			return nil, err
		}
		genotypes = append(genotypes, genotype)
	}
	if err = rows.Err(); err != nil {
		log.Print("genotypeslist: ", err)
		return nil, err
	}
	return genotypes, nil
}

// Update a genotype
func UpdateGenotype(c *gin.Context) (*Genotype, error) {
	db := c.MustGet("db").(*sql.DB)
	genotype := new(Genotype)
	genotype.Name = c.PostForm("name")
	genotype.Chr = c.PostForm("chr")
	genotype.Coordinate = c.PostForm("coordinate")
	genotype.VariantID = c.PostForm("variant_id")
	genotype.Location = c.PostForm("location")
	genotype.Calls = c.PostForm("calls")
	stmt, err := db.Prepare("UPDATE datasets set chr=?, coordinate=?, variant_id=?, location=?, calls=?  WHERE name=? ;")
	defer stmt.Close()
	stmt.Exec(&genotype.Chr, &genotype.Coordinate, &genotype.VariantID, &genotype.Location, &genotype.Calls, &genotype.Name)

	if err != nil {
		log.Print("updaterr: ", err)
		return nil, err
	}

	log.Print("update: ", stmt)
	return genotype, nil
}

// Delete a genotype
func DeleteGenotype(c *gin.Context) (bool, error) {
	db := c.MustGet("db").(*sql.DB)
	name := c.Param("name")
	stmt, err := db.Exec("DELETE FROM genotypes WHERE name=?;", name)
	if err != nil {
		log.Print("delete: ", err)
		return false, err
	}
	log.Print("deletetrue:", stmt)
	return true, nil
}

/*
CREATE TABLE `genotypes` (
	`id` BIGINT NOT NULL AUTO_INCREMENT,
	`name` char(250) NOT NULL,
	`chr` varchar(255) NOT NULL,
	`coordinate` TEXT,
	`variant_id` varchar(255) NOT NULL,
	`calls` TEXT,
	`location` TEXT,
	unique(`name`),
    PRIMARY KEY (`id`)
);

*/
