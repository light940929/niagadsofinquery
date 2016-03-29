package references

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/user/niagadsofinquery/models"
)

// New references
func New(c *gin.Context) {
	reference := models.Reference{}

	c.HTML(http.StatusOK, "references/form", gin.H{
		"title":     "New reference",
		"reference": reference,
	})
}

// Create a reference
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	reference := models.Reference{}
	err := c.Bind(&reference)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionReference).Insert(reference)
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/references")
}

// Edit a reference
func Edit(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	reference := models.Reference{}
	oID := bson.ObjectIdHex(c.Param("_id"))
	err := db.C(models.CollectionReference).FindId(oID).One(&reference)
	if err != nil {
		c.Error(err)
	}

	c.HTML(http.StatusOK, "references/form", gin.H{
		"title":     "Edit reference",
		"reference": reference,
	})
}

// List all references
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	references := []models.Reference{}
	err := db.C(models.CollectionReference).Find(nil).Sort("-title").All(&references)
	if err != nil {
		c.Error(err)
	}
	c.HTML(http.StatusOK, "references/list", gin.H{
		"title":      "references",
		"references": references,
	})
}

// Update a reference
func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	reference := models.Reference{}
	err := c.Bind(&reference)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
	doc := bson.M{
		"name":        reference.Name,
		"description": reference.Description,
		"plateform":   reference.Plateform,
		"length":      reference.Length,
		"source_uri":  reference.SourceUri,
		"format":      reference.Format,
	}
	err = db.C(models.CollectionReference).Update(query, doc)
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/references")
}

// Delete a reference
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
	err := db.C(models.CollectionReference).Remove(query)
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/references")
}
