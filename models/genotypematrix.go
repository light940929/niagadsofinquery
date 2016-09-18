package models

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

const (
	CollectionGenotypematrix = "genotypematrixs"
)

// Genotypematrix model
type Genotypematrix struct {
	Id           string `json:"id"`
	NIAGADSID    string `json:"niagads_id"`
	IndividualID string `json:"individual_id, omitempty"`
	DatasetID    string `json:"dataset_id, omitempty"`
	SNPID        string `json:"snp_id, omitempty"`
	SNPnaID      string `json:"snp_na_id, omitempty"`
	Genotypes    string `json:"genotypes, omitempty"`
}

// Create a Genotypematrix
func CreateGenotypematrix(c *gin.Context) (*Genotypematrix, error) {
	db := c.MustGet("db").(*sql.DB)
	genotypematrix := new(Genotypematrix)
	err := c.Bind(genotypematrix)
	if err != nil {
		log.Print("err: ", err)
		return genotypematrix, nil
	}
	niagadsID := c.PostForm("niagads_id")
	individualID := c.PostForm("individual_id")
	datasetID := c.PostForm("dataset_id")
	snpID := c.PostForm("snp_id")
	snpnaID := c.PostForm("snp_na_id")
	genotypes := c.PostForm("genotypes")

	stmt, err := db.Prepare("INSERT INTO genotypematrix(niagads_id,individual_id,dataset_id,snp_id,snp_na_id,genotypes) VALUES(?, ?, ?, ?, ?, ?);")
	defer stmt.Close()
	log.Print("niagadsID:", niagadsID, "individualID:", individualID, "datasetID:", datasetID, "snpID:", snpID, "snpnaID:", snpnaID, "genotypes:", genotypes)
	stmt.Exec(genotypematrix.NIAGADSID, genotypematrix.IndividualID, genotypematrix.DatasetID, genotypematrix.SNPID, genotypematrix.SNPnaID, genotypematrix.Genotypes)
	if err != nil {
		log.Print("createGenotypematrixerr: ", err)
		log.Print("createGenotypematrix: ", stmt)
		return nil, err
	}
	log.Print("createGenotypematrix: ", genotypematrix)

	errget := db.QueryRow("SELECT * FROM genotypematrix WHERE niagads_id=?;", genotypematrix.NIAGADSID).Scan(&genotypematrix.Id, &genotypematrix.NIAGADSID, &genotypematrix.IndividualID, &genotypematrix.DatasetID, &genotypematrix.SNPID, &genotypematrix.SNPnaID, &genotypematrix.Genotypes)
	log.Print("query: ", errget)
	log.Print("getGenotypematrix: ", genotypematrix.Id)
	if errget != nil {
		log.Print("getGenotypematrixerr: ", errget)
		log.Print("getGenotypematrix_id: ", genotypematrix.Id)
		return nil, err
	}

	return genotypematrix, nil
}

// Get a Genotypematrix
func GetGenotypematrix(c *gin.Context) (*Genotypematrix, error) {

	db := c.MustGet("db").(*sql.DB)
	genotypematrix := new(Genotypematrix)
	id := c.Param("id")
	log.Print("id", id)
	err := db.QueryRow("SELECT * FROM genotypematrix WHERE id=?;", id).Scan(&genotypematrix.Id, &genotypematrix.NIAGADSID, &genotypematrix.IndividualID, &genotypematrix.DatasetID, &genotypematrix.SNPID, &genotypematrix.SNPnaID, &genotypematrix.Genotypes)
	log.Print("query: ", err)
	log.Print("getGenotypematrix: ", genotypematrix.Id)
	if err != nil {
		log.Print("getGenotypematrixerr: ", err)
		log.Print("getGenotypematrix_id: ", id)
		return nil, err
	}
	// rows, err := db.Query("SELECT * FROM genotypematrix WHERE id=?", id)
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()
	//
	// var genotypematrixs []*Genotypematrix
	//
	// for rows.Next() {
	// 	genotypematrix := new(Genotypematrix)
	// 	//var individualID string
	// 	//var status string
	// 	err := rows.Scan(&genotypematrix.Id, &genotypematrix.NIAGADSID, &genotypematrix.IndividualID, &genotypematrix.DatasetID, &genotypematrix.SNPID, &genotypematrix.SNPnaID, &genotypematrix.Genotypes)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	genotypematrixs = append(genotypematrixs, genotypematrix)
	// 	log.Print("GetGenotypematrixerr: ", genotypematrixs)
	// }
	// if err = rows.Err(); err != nil {
	// 	return nil, err
	// }

	return genotypematrix, nil
}

// List all Genotypematrixs
func ListGenotypematrixs(c *gin.Context) ([]*Genotypematrix, error) {
	db := c.MustGet("db").(*sql.DB)
	rows, err := db.Query("SELECT * FROM genotypematrix")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//inquerysets := make([]*Inqueryset, 0)
	var genotypematrixs []*Genotypematrix

	for rows.Next() {
		genotypematrix := new(Genotypematrix)
		err := rows.Scan(&genotypematrix.Id, &genotypematrix.NIAGADSID, &genotypematrix.IndividualID, &genotypematrix.DatasetID, &genotypematrix.SNPID, &genotypematrix.SNPnaID, &genotypematrix.Genotypes)
		if err != nil {
			return nil, err
		}
		genotypematrixs = append(genotypematrixs, genotypematrix)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return genotypematrixs, nil
}

// Update a Genotypematrix
func UpdateGenotypematrix(c *gin.Context) (*Genotypematrix, error) {
	db := c.MustGet("db").(*sql.DB)
	genotypematrix := new(Genotypematrix)
	err := c.Bind(genotypematrix)
	if err != nil {
		log.Print("err: ", err)
		return genotypematrix, nil
	}
	id := c.PostForm("id")
	niagadsID := c.PostForm("niagads_id")
	individualID := c.PostForm("individual_id")
	datasetID := c.PostForm("dataset_id")
	snpID := c.PostForm("snp_id")
	snpnaID := c.PostForm("snp_na_id")
	genotypes := c.PostForm("genotypes")
	stmt, err := db.Prepare("UPDATE genotypematrix set niagads_id=?, individual_id=?, dataset_id=?, snp_id=?, snp_na_id=?, genotypes=? WHERE id=? ;")
	//defer stmt.Close()
	log.Print("query: ", err)
	log.Print("niagadsID:", niagadsID, "individualID:", individualID, "datasetID:", datasetID, "snpID:", snpID, "snpnaID:", snpnaID, "genotypes:", genotypes, "id:", id)
	stmt.Exec(genotypematrix.NIAGADSID, genotypematrix.IndividualID, genotypematrix.DatasetID, genotypematrix.SNPID, genotypematrix.SNPnaID, genotypematrix.Genotypes, genotypematrix.Id)

	if err != nil {
		log.Print("updaterr: ", err)
		return nil, err
	}

	log.Print("update: ", stmt)

	return genotypematrix, nil
}

// Delete a Genotypematrix
func DeleteGenotypematrix(c *gin.Context) (bool, error) {
	db := c.MustGet("db").(*sql.DB)
	id := c.Param("id")
	stmt, err := db.Exec("DELETE FROM genotypematrix WHERE id=?;", id)
	if err != nil {
		log.Print("delete: ", err)
		return false, err
	}
	log.Print("deleteGenotypematrix:", stmt)
	return true, nil
}

/*

Id           string   `json:"id"`
NIAGADSID    string   `json:"niagads_id"`
IndividualID string   `json:"individual_id, omitempty"`
DatasetID    string   `json:"dataset_id, omitempty"`
SNPID        string   `json:"snp_id, omitempty"`
SNPnaID      string   `json:"snp_na_id, omitempty"`
Genotypes    Genotype `json:"genotypes, omitempty"`

CREATE TABLE `genotypematrix` (
	`id` BIGINT NOT NULL AUTO_INCREMENT,
	`niagads_id` varchar(3306) NOT NULL,
	`individual_id` TEXT,
	`dataset_id` TEXT,
	`snp_id` TEXT,
	`snp_na_id` varchar(500),
    `genotypes` TEXT,
    PRIMARY KEY (`id`)
);

{"niagads_id":"NGS0008123, NGS0008124, NGS0008125, NGS0008126, NGS0008127","individual_id":"EURAD1, EURAD10, EURAD100, EURAD101, EURAD102","dataset_id":"NG00028, NG00028, NG00028, NG00028,  NG00028","snp_id":"","snp_na_id":"","genotypes":""}

{
  "success": [
    {
      "id": "2",
      "niagads_id": "NGS0008123, NGS0008124, NGS0008125, NGS0008126, NGS0008127",
      "individual_id": "EURAD1, EURAD10, EURAD100, EURAD101, EURAD102",
      "dataset_id": "NG00028, NG00028, NG00028, NG00028,  NG00028",
      "snp_id": "",
      "snp_na_id": "",
      "genotypes": ""
    }
  ]
}

*/
