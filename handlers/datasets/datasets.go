package datasets

import (
	"net/http"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

// Dataset model
type Dataset struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Created     time.Time `json:"created"`
}

// Create a dataset
func Create(c *gin.Context) (*Dataset, error) {
	db := c.MustGet("db").(*sql.DB)
	dataset := new(Dataset)
	dataset.Name = c.PostForm("name")
	dataset.Description = c.PostForm("description")
	dataset.Type = c.PostForm("type")
	//dataset.Created = c.PostForm("created")

	err := c.Bind(&dataset)
	if err != nil {
		c.Error(err)
		return nil, err
	}

	err = db.QueryRow("INSERT INTO datasets(name,description,type) VALUES($1,$2,$3) returning id;", &dataset.Name, &dataset.Description, &dataset.Type).Scan(&dataset.ID)
	if err != nil {
		return nil, err
	}
	c.Redirect(http.StatusMovedPermanently, "/datasets")
	return dataset, nil
}

// Get a dataset
func Get(c *gin.Context) (*Dataset, error) {

	db := c.MustGet("db").(*sql.DB)
	dataset := new(Dataset)
	dataset.Name = c.PostForm("name")
	err := db.QueryRow("SELECT * FROM datasets WHERE name=$1;", dataset.Name).Scan(&dataset.ID, &dataset.Name, &dataset.Description, &dataset.Type, &dataset.Created)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Not found this dataset"})
		return nil, err
	}

	c.JSON(http.StatusCreated, gin.H{"datasets": dataset})

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
		err := rows.Scan(&dataset.ID, &dataset.Name, &dataset.Description, &dataset.Type, &dataset.Created)
		if err != nil {
			return nil, err
		}
		datasets = append(datasets, dataset)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	c.JSON(http.StatusOK, gin.H{
		"datasets": datasets,
	})
	return datasets, nil
}

// Update a dataset
func Update(c *gin.Context) (*Dataset, error) {
	db := c.MustGet("db").(*sql.DB)
	dataset := new(Dataset)
	dataset.Name = c.PostForm("name")
	dataset.Description = c.PostForm("description")
	dataset.Type = c.PostForm("type")

	err := db.QueryRow("UPDATE datasets SET (name, description, type) = ($1,$2,$3) WHERE id=$4 returning id;",
		&dataset.Name,
		&dataset.Description,
		&dataset.Type,
		&dataset.ID).Scan(&dataset.ID)

	if err != nil {
		return nil, err
	}

	c.Redirect(http.StatusMovedPermanently, "/datasets")
	return dataset, nil
}

// Delete a dataset
func Delete(c *gin.Context) (bool, error) {
	db := c.MustGet("db").(*sql.DB)
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM datasets WHERE id=$1", id)
	if err != nil {
		return false, err
	}
	c.Redirect(http.StatusMovedPermanently, "/datasets")
	return true, nil
}
