package models

import _ "github.com/go-sql-driver/mysql"

const (
	CollectionGenotype = "genotypes"
)

// Genotype model
type Genotype struct {
	//Id           int    `json:"id"`
	//IndividualID string `json:"individual_id"`
	//Chr          string `json:"chr"`
	//Coordinate   string `json:"coordinate, omitempty"`
	//VariantID    string `json:"variant_id"`
	//Status       string `json:"status"`
	//Location     string `json:"location, omitempty"`
	NIAGADSID string `json:"niagads_id"`
	Calls     string `json:"calls, omitempty"`
}

/*
// Create a genotype
func CreateGenotype(c *gin.Context) (*Genotype, error) {
	db := c.MustGet("db").(*sql.DB)
	genotype := new(Genotype)
	err := c.Bind(genotype)
	if err != nil {
		log.Print("err: ", err)
		return genotype, nil
	}
	individual_id := c.PostForm("individual_id")
	chr := c.PostForm("chr")
	coordinate := c.PostForm("coordinate")
	variantID := c.PostForm("variant_id")
	status := c.PostForm("status")
	location := c.PostForm("location")
	calls := c.PostForm("calls")

	stmt, err := db.Prepare("INSERT INTO genotypes(individual_id,chr,coordinate,variant_id,status,location,calls) VALUES(?, ?, ?, ?, ?, ?, ?);")
	defer stmt.Close()
	log.Print("individual_id:", individual_id, "chr", chr, "coordinate", coordinate, "variantID", variantID, "status", status, "location", location, "calls", calls)
	stmt.Exec(genotype.IndividualID, genotype.Chr, genotype.Coordinate, genotype.VariantID, genotype.Status, genotype.Location, genotype.Calls)
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
	genotype := new(Genotype)
	individual_id := c.Param("individual_id")
	log.Print("individual_id", individual_id)
	err := db.QueryRow("SELECT * FROM genotypes WHERE individual_id=?;", individual_id).Scan(&genotype.Id, &genotype.IndividualID, &genotype.Chr, &genotype.Coordinate, &genotype.VariantID, &genotype.Status, &genotype.Location, &genotype.Calls)
	log.Print("getgenotype: ", genotype)
	if err != nil {
		log.Print("getgenotyperr: ", err)
		log.Print("getgenotypeindividual_id: ", individual_id)
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
		err := rows.Scan(&genotype.Id, &genotype.IndividualID, &genotype.Chr, &genotype.Coordinate, &genotype.VariantID, &genotype.Status, &genotype.Location, &genotype.Calls)
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
	err := c.Bind(genotype)
	if err != nil {
		log.Print("err: ", err)
		return genotype, nil
	}
	individual_id := c.PostForm("individual_id")
	chr := c.PostForm("chr")
	coordinate := c.PostForm("coordinate")
	variantID := c.PostForm("variant_id")
	status := c.PostForm("status")
	location := c.PostForm("location")
	calls := c.PostForm("calls")

	stmt, err := db.Prepare("UPDATE genotypes set chr=?, coordinate=?, variant_id=?, status=?, location=?, calls=?  WHERE individual_id=? ;")
	defer stmt.Close()

	log.Print("individual_id", individual_id, "chr", chr, "coordinate", coordinate, "variantID", variantID, "status", status, "location", location, "calls", calls)
	stmt.Exec(genotype.Chr, genotype.Coordinate, genotype.VariantID, genotype.Status, genotype.Location, genotype.Calls, genotype.IndividualID)

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
	individual_id := c.Param("individual_id")
	stmt, err := db.Exec("DELETE FROM genotypes WHERE individual_id=?;", individual_id)
	if err != nil {
		log.Print("delete: ", err)
		return false, err
	}
	log.Print("deletetrue:", stmt)
	return true, nil
}
*/

/*
CREATE TABLE `genotypes` (
    `niagads_id` varchar(500) NOT NULL,
	`calls` varchar(500),
	CONSTRAINT pk_genotypes PRIMARY KEY (`niagads_id`, `calls`)
);


CREATE TABLE `genotypes` (
	`id` BIGINT NOT NULL AUTO_INCREMENT,
	`individual_id` char(255) NOT NULL,
	`chr` varchar(255) NOT NULL,
	`coordinate` TEXT,
	`variant_id` varchar(500) NOT NULL,
	`status` varchar(50) NOT NULL,
	`calls` TEXT,
	`location` TEXT,
	KEY (`variant_id`, `status`),
	unique(`variant_id`),
    PRIMARY KEY (`id`)
);


mysql> select * from genotypes;

+----+---------------+-----+------------+------------+--------+-------------------------+----------+
| id | individual_id | chr | coordinate | variant_id | status | calls                   | location |
+----+---------------+-----+------------+------------+--------+-------------------------+----------+
|  1 | NGS0002190    | 1   | 500830     | rs1000071  | AA     | G T G T G G T T G T T T | 0        |
+----+---------------+-----+------------+------------+--------+-------------------------+----------+



*/
