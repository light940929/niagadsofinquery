package datasets

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/user/niagadsofinquery/models"
)

// New datasets
func New(c *gin.Context) {
	dataset := models.Dataset{}

	c.HTML(http.StatusOK, "datasets/form", gin.H{
		"title":   "New dataset",
		"dataset": dataset,
	})
}

// Create a dataset
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	dataset := models.Dataset{}
	err := c.Bind(&dataset)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionDataset).Insert(dataset)
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/datasets")
}

// Edit a dataset
func Edit(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	dataset := models.Dataset{}
	oID := bson.ObjectIdHex(c.Param("_id"))
	err := db.C(models.CollectionDataset).FindId(oID).One(&dataset)
	if err != nil {
		c.Error(err)
	}

	c.HTML(http.StatusOK, "datasets/form", gin.H{
		"title":   "Edit dataset",
		"dataset": dataset,
	})
}

// List all datasets
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	datasets := []models.Dataset{}
	err := db.C(models.CollectionDataset).Find(nil).Sort("-title").All(&datasets)
	if err != nil {
		c.Error(err)
	}
	c.HTML(http.StatusOK, "datasets/list", gin.H{
		"title":    "datasets",
		"datasets": datasets,
	})
}

// Update a dataset
func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	dataset := models.Dataset{}
	err := c.Bind(&dataset)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
	doc := bson.M{
		"name":        dataset.Name,
		"description": dataset.Description,
		"type":        dataset.Type,
		"created":     dataset.Created,
	}
	err = db.C(models.CollectionDataset).Update(query, doc)
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/datasets")
}

// Delete a dataset
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
	err := db.C(models.CollectionDataset).Remove(query)
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/datasets")
}
