package models

import (
	"database/sql"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const (
	// 	VARIANTS          = "variants"
	// 	ANNOTATIONS       = "annotations"
	CollectionDataset = "datasets"
)

// Dataset model
type Dataset struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Type        string     `json:"type"`
	Created     *time.Time `json:"created_at,omitempty"`
}

// Create a dataset
func CreateDataset(c *gin.Context) (*Dataset, error) {
	db := c.MustGet("db").(*sql.DB)
	dataset := new(Dataset)
	dataset.Name = c.PostForm("name")
	dataset.Description = c.PostForm("description")
	dataset.Type = c.PostForm("type")

	stmt, err := db.Prepare("INSERT INTO datasets(name,description,type) VALUES(?, ?, ?);")
	defer stmt.Close()
	stmt.Exec(&dataset.Name, &dataset.Description, &dataset.Type)
	if err != nil {
		log.Print("createdataseterr: ", err)
		log.Print("createdataset: ", stmt)
		return nil, err
	}
	log.Print("createdataset: ", dataset)
	return dataset, nil
}

// Get a dataset
func GetDataset(c *gin.Context) (*Dataset, error) {

	db := c.MustGet("db").(*sql.DB)
	name := c.PostForm("name")
	dataset := new(Dataset)
	err := db.QueryRow("SELECT * FROM datasets WHERE name=?;", name).Scan(&dataset.Id, &dataset.Name, &dataset.Description, &dataset.Type, &dataset.Created)
	log.Print("getdataset: ", dataset)
	if err != nil {
		log.Print("getdataseterr: ", err)
		log.Print("getdatasetname: ", name)
		return nil, err
	}
	return dataset, nil
}

// List all datasets
func ListDatasets(c *gin.Context) ([]*Dataset, error) {
	db := c.MustGet("db").(*sql.DB)
	rows, err := db.Query("SELECT * FROM datasets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//datasets := make([]*Dataset, 0)
	var datasets []*Dataset

	for rows.Next() {
		dataset := new(Dataset)
		err := rows.Scan(&dataset.Id, &dataset.Name, &dataset.Description, &dataset.Type, &dataset.Created)
		if err != nil {
			return nil, err
		}
		datasets = append(datasets, dataset)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return datasets, nil
}

// Update a dataset
func UpdateDataset(c *gin.Context) (*Dataset, error) {
	db := c.MustGet("db").(*sql.DB)
	dataset := new(Dataset)
	dataset.Name = c.PostForm("name")
	dataset.Description = c.PostForm("description")
	dataset.Type = c.PostForm("type")
	stmt, err := db.Prepare("UPDATE datasets set description=?, type=? WHERE name=? ;")
	defer stmt.Close()
	stmt.Exec(&dataset.Description, &dataset.Type, &dataset.Name)

	if err != nil {
		log.Print("updaterr: ", err)
		return nil, err
	}

	log.Print("update: ", stmt)
	return dataset, nil
}

// Delete a dataset
func DeleteDataset(c *gin.Context) (bool, error) {
	db := c.MustGet("db").(*sql.DB)
	name := c.Param("name")
	stmt, err := db.Exec("DELETE FROM datasets WHERE name=?;", name)
	if err != nil {
		log.Print("delete: ", err)
		return false, err
	}
	log.Print("deletetrue:", stmt)
	return true, nil
}

/*
CREATE TABLE `datasets` (
	`id` BIGINT NOT NULL AUTO_INCREMENT,
	`name` char(50) NOT NULL,
	`description` varchar(255) NOT NULL,
	`type` char(200) NOT NULL,
	`created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	unique(`name`),
    PRIMARY KEY (`id`)
);

*/
