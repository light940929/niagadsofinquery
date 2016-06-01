package models

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

const (
	CollectionInqueryset = "inquerysets"
)

// Inqueryset model
type Inqueryset struct {
	Id             string `json:"id"`
	InquerysetName string `json:"name"`
	IndividualID   string `json:"individual_id"`
	VariantID      string `json:"variant_id"`
}

// Create a Inqueryset
func CreateInqueryset(c *gin.Context) (*Inqueryset, error) {
	db := c.MustGet("db").(*sql.DB)
	inqueryset := new(Inqueryset)
	inqueryset.InquerysetName = c.PostForm("name")
	inqueryset.IndividualID = c.PostForm("individual_id")
	inqueryset.VariantID = c.PostForm("variant_id")

	stmt, err := db.Prepare("INSERT INTO inquerysets(name,individual_id,variant_id) VALUES(?, ?, ?);")
	defer stmt.Close()
	stmt.Exec(&inqueryset.InquerysetName, &inqueryset.IndividualID, &inqueryset.VariantID)
	if err != nil {
		log.Print("createInqueryseterr: ", err)
		log.Print("createInqueryset: ", stmt)
		return nil, err
	}
	log.Print("createInqueryset: ", inqueryset)
	return inqueryset, nil
}

// Get a Inqueryset
func GetInqueryset(c *gin.Context) (*Inqueryset, error) {

	db := c.MustGet("db").(*sql.DB)
	name := c.PostForm("name")
	inqueryset := new(Inqueryset)
	err := db.QueryRow("SELECT * FROM inquerysets WHERE name=?;", name).Scan(&inqueryset.Id, &inqueryset.InquerysetName, &inqueryset.IndividualID, &inqueryset.VariantID)
	log.Print("getInqueryset: ", inqueryset)
	if err != nil {
		log.Print("getInqueryseterr: ", err)
		log.Print("getInquerysetname: ", name)
		return nil, err
	}
	return inqueryset, nil
}

// List all Inquerysets
func ListInquerysets(c *gin.Context) ([]*Inqueryset, error) {
	db := c.MustGet("db").(*sql.DB)
	rows, err := db.Query("SELECT * FROM inquerysets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//inquerysets := make([]*Inqueryset, 0)
	var inquerysets []*Inqueryset

	for rows.Next() {
		inqueryset := new(Inqueryset)
		err := rows.Scan(&inqueryset.Id, &inqueryset.InquerysetName, &inqueryset.IndividualID, &inqueryset.VariantID)
		if err != nil {
			return nil, err
		}
		inquerysets = append(inquerysets, inqueryset)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return inquerysets, nil
}

// Update a Inqueryset
func UpdateInqueryset(c *gin.Context) (*Inqueryset, error) {
	db := c.MustGet("db").(*sql.DB)
	inqueryset := new(Inqueryset)
	inqueryset.InquerysetName = c.PostForm("name")
	inqueryset.IndividualID = c.PostForm("individual_id")
	inqueryset.VariantID = c.PostForm("variant_id")
	stmt, err := db.Prepare("UPDATE inquerysets set individual_id=?, variant_id=? WHERE name=? ;")
	defer stmt.Close()
	stmt.Exec(&inqueryset.IndividualID, &inqueryset.VariantID, &inqueryset.InquerysetName)

	if err != nil {
		log.Print("updaterr: ", err)
		return nil, err
	}

	log.Print("update: ", stmt)
	return inqueryset, nil
}

// Delete a Inqueryset
func DeleteInqueryset(c *gin.Context) (bool, error) {
	db := c.MustGet("db").(*sql.DB)
	name := c.Param("name")
	stmt, err := db.Exec("DELETE FROM inquerysets WHERE name=?;", name)
	if err != nil {
		log.Print("delete: ", err)
		return false, err
	}
	log.Print("deleteInquery:", stmt)
	return true, nil
}

/*
CREATE TABLE `inquerysets` (
	`id` BIGINT NOT NULL AUTO_INCREMENT,
	`name` char(50) NOT NULL,
	`individual_id` varchar(255) NOT NULL,
	`variant_id` varchar(255) NOT NULL,
	unique(`name`),
    PRIMARY KEY (`id`)
);

*/
