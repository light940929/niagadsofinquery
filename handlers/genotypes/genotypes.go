package genotypes

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/user/niagadsofinquery/models"
)

// New genotypes
func New(c *gin.Context) {
	genotype := models.Genotype{}

	c.HTML(http.StatusOK, "genotypes/form", gin.H{
		"title":    "New genotype",
		"genotype": genotype,
	})
}

// Create a genotype
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	genotype := models.Genotype{}
	err := c.Bind(&genotype)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionGenotype).Insert(genotype)
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/genotypes")

	var json models.Genotype
	c.BindJSON(&json) // This will infer what binder to use depending on the content-type header.
	if json.Title != " " && json.Chr != " " && json.Coordinate != " " {
		c.JSON(http.StatusCreated, gin.H{"genotypes": genotype})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Please must input Title, Chr, Coordinate"})
	}
}

// Get a genotype
func Get(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	genotype := models.Genotype{}
	oID := bson.ObjectIdHex(c.Param("_id"))
	err := db.C(models.CollectionGenotype).FindId(oID).One(&genotype)
	if err != nil {
		c.Error(err)
	}

	c.HTML(http.StatusOK, "genotypes/form", gin.H{
		"title":    "Edit genotype",
		"genotype": genotype,
	})

	var json models.Genotype
	c.BindJSON(&json) // This will infer what binder to use depending on the content-type header.
	if oID == genotype.Id {
		c.JSON(http.StatusCreated, gin.H{"genotypes": genotype})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Not found this genotype"})
	}
}

// List all genotypes
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	genotypes := []models.Genotype{}
	err := db.C(models.CollectionGenotype).Find(nil).Sort("-title").All(&genotypes)
	if err != nil {
		c.Error(err)
	}
	//HTML
	// c.HTML(http.StatusOK, "genotypes/list", gin.H{
	// 	"title":     "genotypes",
	// 	"genotypes": genotypes,
	// })

	c.JSON(http.StatusOK, gin.H{
		"genotypes": genotypes,
	})
}

// Update a genotype
func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	genotype := models.Genotype{}
	err := c.Bind(&genotype)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
	doc := bson.M{
		"title":      genotype.Title,
		"chr":        genotype.Chr,
		"coordinate": genotype.Coordinate,
		"variant_id": genotype.VariantID,
		"location":   genotype.Location,
		"Call":       genotype.Call,
	}
	err = db.C(models.CollectionGenotype).Update(query, doc)
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/genotypes")

	var json models.Genotype
	c.BindJSON(&json) // This will infer what binder to use depending on the content-type header.
	if json.Title != " " && json.Chr != " " && json.Coordinate != " " {
		c.JSON(http.StatusCreated, gin.H{"genotypes": genotype})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Please must update Title, Chr, Coordinate"})
	}
}

// Delete a genotype
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
	err := db.C(models.CollectionGenotype).Remove(query)
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/genotypes")

	var json models.Genotype
	c.BindJSON(&json) // This will infer what binder to use depending on the content-type header.
	if json.Id.String() != c.Params.ByName("_id") {
		c.JSON(http.StatusOK, gin.H{"status": "Delete [ " + c.Params.ByName("_id") + " ]"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": "Please input right id"})
	}
}
