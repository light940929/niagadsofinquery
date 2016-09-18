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
	Id           string `json:"id"`
	NIAGADSID    string `json:"niagads_id"`
	IndividualID string `json:"individual_id, omitempty"`
	DatasetID    string `json:"dataset_id, omitempty"`
	SNPpos       string `json:"snp_pos, omitempty"`
	SNPposna     string `json:"snp_pos_na, omitempty"`
	Genotypes    string `json:"genotypes, omitempty"`
}

// Create a Inqueryset
func CreateInqueryset(c *gin.Context) (*Inqueryset, error) {
	db := c.MustGet("db").(*sql.DB)
	inqueryset := new(Inqueryset)
	err := c.Bind(inqueryset)
	if err != nil {
		log.Print("err: ", err)
		return inqueryset, nil
	}
	niagadsID := c.PostForm("niagads_id")
	individualID := c.PostForm("individual_id")
	datasetID := c.PostForm("dataset_id")
	snppos := c.PostForm("snp_pos")
	snpposna := c.PostForm("snp_pos_na")
	genotypes := c.PostForm("genotypes")

	stmt, err := db.Prepare("INSERT INTO inquerysets(niagads_id, individual_id, dataset_id, snp_pos, snp_pos_na, genotypes) VALUES(?, ?, ?, ?, ?, ?);")
	defer stmt.Close()
	log.Print("niagadsID:", niagadsID, "individualID:", individualID, "datasetID:", datasetID, "snp_pos:", snppos, "snp_pos_na:", snpposna, "genotypes:", genotypes)
	stmt.Exec(inqueryset.NIAGADSID, inqueryset.IndividualID, inqueryset.DatasetID, inqueryset.SNPpos, inqueryset.SNPposna, inqueryset.Genotypes)
	if err != nil {
		log.Print("createInqueryseterr: ", err)
		log.Print("createInqueryset: ", stmt)
		return nil, err
	}
	log.Print("createInqueryset: ", inqueryset)
	errget := db.QueryRow("SELECT * FROM inquerysets WHERE niagads_id=?;", inqueryset.NIAGADSID).Scan(&inqueryset.Id, &inqueryset.NIAGADSID, &inqueryset.IndividualID, &inqueryset.DatasetID, &inqueryset.SNPpos, &inqueryset.SNPposna, &inqueryset.Genotypes)
	log.Print("getInqueryset: ", inqueryset)
	if errget != nil {
		log.Print("getInqueryseterr: ", errget)
		log.Print("getInqueryset_id: ", inqueryset.Id)
		return nil, err
	}

	return inqueryset, nil
}

// Get a Inqueryset
func GetInqueryset(c *gin.Context) (*Inqueryset, error) {

	db := c.MustGet("db").(*sql.DB)
	inqueryset := new(Inqueryset)
	id := c.Param("id")
	log.Print("id", id)
	err := db.QueryRow("SELECT * FROM inquerysets WHERE id=?;", id).Scan(&inqueryset.Id, &inqueryset.NIAGADSID, &inqueryset.IndividualID, &inqueryset.DatasetID, &inqueryset.SNPpos, &inqueryset.SNPposna, &inqueryset.Genotypes)
	log.Print("getInqueryset: ", inqueryset)
	if err != nil {
		log.Print("getInqueryseterr: ", err)
		log.Print("getInqueryset_id: ", id)
		return nil, err
	}
	// rows, err := db.Query("SELECT * FROM inquerysets WHERE id=?", id)
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()
	//
	// var inquerysets []*Inqueryset
	//
	// for rows.Next() {
	// 	inqueryset := new(Inqueryset)
	// 	//var individualID string
	// 	//var status string
	// 	err := rows.Scan(&inqueryset.Id, &inqueryset.NIAGADSID, &inqueryset.IndividualID, &inqueryset.DatasetID, &inqueryset.SNPpos, &inqueryset.SNPposna, &inqueryset.Genotypes)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	inquerysets = append(inquerysets, inqueryset)
	// }
	// if err = rows.Err(); err != nil {
	// 	return nil, err
	// }

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
		err := rows.Scan(&inqueryset.Id, &inqueryset.NIAGADSID, &inqueryset.IndividualID, &inqueryset.DatasetID, &inqueryset.SNPpos, &inqueryset.SNPposna, &inqueryset.Genotypes)
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
	err := c.Bind(inqueryset)
	if err != nil {
		log.Print("err: ", err)
		return inqueryset, nil
	}
	id := c.PostForm("id")
	niagadsID := c.PostForm("niagads_id")
	individualID := c.PostForm("individual_id")
	datasetID := c.PostForm("dataset_id")
	snppos := c.PostForm("snp_pos")
	snpposna := c.PostForm("snp_pos_na")
	genotypes := c.PostForm("genotypes")
	stmt, err := db.Prepare("UPDATE inquerysets set niagads_id=?, individual_id=?, dataset_id=?, snp_pos=?, snp_pos_na=?, genotypes=? WHERE id=? ;")
	defer stmt.Close()
	log.Print("niagadsID:", niagadsID, "individualID:", individualID, "datasetID:", datasetID, "snppos:", snppos, "snpposna:", snpposna, "genotypes:", genotypes, "id:", id)
	stmt.Exec(inqueryset.NIAGADSID, inqueryset.IndividualID, inqueryset.DatasetID, inqueryset.SNPpos, inqueryset.SNPposna, inqueryset.Genotypes, inqueryset.Id)

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
	id := c.Param("id")
	stmt, err := db.Exec("DELETE FROM inquerysets WHERE id=?;", id)
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
	`individual_id` varchar(1024) NOT NULL,
	`variant_id` varchar(1024) NOT NULL,
	unique(`individual_id`),
    PRIMARY KEY (`id`)
);

CREATE TABLE `inquerysets` (
	`id` BIGINT NOT NULL AUTO_INCREMENT,
	`niagads_id` varchar(500) NOT NULL,
	`individual_id` varchar(500) NOT NULL,
	`dataset_id` varchar(500),
	`snp_pos` varchar(500),
	`snp_pos_na` varchar(500),
    `genotypes` TEXT,
    PRIMARY KEY (`id`)
);



+----+------------------------+----------------------+
| id | individual_id          | variant_id           |
+----+------------------------+----------------------+
|  1 | NGS0002190             | rs1000071            |
|  2 | NGS0000904             | rs1000072            |
|  3 | NGS0002190, NGS0000904 | rs1000075, rs1000072 |
+----+------------------------+----------------------+

/api/genotypebyposition
{"success"[
  {"NIAGADS_id" : "NGSXXXXX, NGSXXXXX, NGSXXXXX",
   "individual_id": "DUKE8002, DUKE8211, 100902",
   "dataset_id": "NGXXXX,NGXXXX,NGXXXX",
   "SNP_pos":"chr1:300,chr1:123",
   "SNP_pos_na":"chr7:12345",
   "genotypes": ["C C A T ", "A A T T","C C A A"]}}

*/
