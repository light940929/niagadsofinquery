package phenotypes

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/user/niagadsofinquery/models"
)

// New phenotypes
func New(c *gin.Context) {
	phenotype := models.Phenotype{}

	c.HTML(http.StatusCreated, "phenotypes/form", gin.H{
		"title":     "New phenotype",
		"phenotype": phenotype,
	})

}

// Create a phenotype
func Create(c *gin.Context) {

	phenotype := models.Phenotype{}
	db := c.MustGet("db").(*mgo.Database)

	err := c.Bind(&phenotype)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionPhenotype).Insert(phenotype)
	if err != nil {
		c.Error(err)
	}

	c.Redirect(http.StatusMovedPermanently, "/phenotypes")

	var json models.Phenotype
	c.BindJSON(&json) // This will infer what binder to use depending on the content-type header.
	if json.Title != " " && json.FamilyID != " " && json.IndividualID != " " {
		c.JSON(http.StatusCreated, gin.H{"phenotypes": phenotype})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Please must input Title, FamilyID, IndividualID"})
	}
}

// Get a phenotype
func Get(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	phenotype := models.Phenotype{}
	oID := bson.ObjectIdHex(c.Param("_id"))
	err := db.C(models.CollectionPhenotype).FindId(oID).One(&phenotype)
	if err != nil {
		c.Error(err)
	}

	c.HTML(http.StatusOK, "phenotypes/form", gin.H{
		"title":     "Edit phenotype",
		"phenotype": phenotype,
	})

	var json models.Phenotype
	c.BindJSON(&json) // This will infer what binder to use depending on the content-type header.
	if oID == phenotype.Id {
		c.JSON(http.StatusCreated, gin.H{"phenotypes": phenotype})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Not found this phenotype"})
	}
}

// List all phenotypes
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	phenotypes := []models.Phenotype{}
	err := db.C(models.CollectionPhenotype).Find(nil).Sort("-title").All(&phenotypes)
	if err != nil {
		c.Error(err)
	}
	c.HTML(http.StatusOK, "phenotypes/list", gin.H{
		"title":      "Phenotypes",
		"phenotypes": phenotypes,
	})

	c.JSON(http.StatusOK, gin.H{
		"phenotypes": phenotypes,
	})
}

// Update a phenotype
func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	phenotype := models.Phenotype{}
	err := c.Bind(&phenotype)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
	doc := bson.M{
		"title":            phenotype.Title,
		"sex":              phenotype.Sex,
		"birth":            phenotype.Birth,
		"age_on_set":       phenotype.Ageonset,
		"family_id":        phenotype.FamilyID,
		"individual_id":    phenotype.IndividualID,
		"paternal_id":      phenotype.PaternalID,
		"maternal_id":      phenotype.MaternalID,
		"affection_status": phenotype.AffectionStatus,
	}
	err = db.C(models.CollectionPhenotype).Update(query, doc)
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/phenotypes")

	var json models.Phenotype
	c.BindJSON(&json) // This will infer what binder to use depending on the content-type header.
	if json.Title != " " && json.FamilyID != " " && json.IndividualID != " " {
		c.JSON(http.StatusCreated, gin.H{"phenotypes": phenotype})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Please must update Title, FamilyID, IndividualID"})
	}
}

// Delete a phenotype
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
	err := db.C(models.CollectionPhenotype).Remove(query)
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/phenotypes")

	var json models.Phenotype
	c.BindJSON(&json) // This will infer what binder to use depending on the content-type header.
	if json.Id.String() != c.Params.ByName("_id") {
		c.JSON(http.StatusOK, gin.H{"status": "Delete [ " + c.Params.ByName("_id") + " ]"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": "Please input right id"})
	}
}
