package inquerysets

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/gin-gonic/gin"
	"github.com/user/niagadsofinquery/models"
)

// Create a inqueryset
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	inqueryset := models.Inqueryset{}
	err := c.Bind(&inqueryset)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionInqueryset).Insert(inqueryset)
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/inquerysets")

	var json models.Inqueryset
	c.BindJSON(&json) // This will infer what binder to use depending on the content-type header.
	if json.IndividualID != " " && json.VariantID != " " {
		c.JSON(http.StatusCreated, gin.H{"inquerysets": inqueryset})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Please must input Dataset, IndividualID, VariantID"})
	}
}

// Get a inqueryset
func Get(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	inqueryset := models.Inqueryset{}
	//oID := bson.ObjectIdHex(c.Param("id"))
	//oID := c.Param("dataset") + c.Param("individual_id") + c.Param("variant_id")
	oID := c.Param("id")
	err := db.C(models.CollectionInqueryset).FindId(oID).One(&inqueryset)
	if err != nil {
		c.Error(err)
	}

	c.HTML(http.StatusOK, "inquerysets/form", gin.H{
		"title":      "Edit inqueryset",
		"inqueryset": inqueryset,
	})

	var json models.Inqueryset
	c.BindJSON(&json)         // This will infer what binder to use depending on the content-type header.
	if oID == inqueryset.Id { //oID == inqueryset.Dataset+inqueryset.IndividualID+inqueryset.VariantID
		c.JSON(http.StatusCreated, gin.H{"inquerysets": inqueryset})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Not found this inqueryset"})
	}
}

// List all inquerysets
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	inquerysets := []models.Inqueryset{}
	err := db.C(models.CollectionInqueryset).Find(nil).Sort("-title").All(&inquerysets)
	if err != nil {
		c.Error(err)
	}
	//HTML
	// c.HTML(http.StatusOK, "inquerysets/list", gin.H{
	// 	"title":     "inquerysets",
	// 	"inquerysets": inquerysets,
	// })

	c.JSON(http.StatusOK, gin.H{
		"inquerysets": inquerysets,
	})
}

// Update a inqueryset
// func Update(c *gin.Context) {
// 	db := c.MustGet("db").(*mgo.Database)
//
// 	inqueryset := models.Inqueryset{}
// 	err := c.Bind(&inqueryset)
// 	if err != nil {
// 		c.Error(err)
// 		return
// 	}
//
// 	query := bson.M{"dataset": c.Param("dataset")}
// 	doc := bson.M{
// 		"variant_id":    inqueryset.VariantID,
// 		"individual_id": inqueryset.IndividualID,
// 	}
// 	err = db.C(models.CollectionInqueryset).Update(query, doc)
// 	if err != nil {
// 		c.Error(err)
// 	}
// 	c.Redirect(http.StatusMovedPermanently, "/inquerysets")
//
// 	var json models.Inqueryset
// 	c.BindJSON(&json) // This will infer what binder to use depending on the content-type header.
// 	if json.Dataset != " " && json.IndividualID != " " && json.VariantID != " " {
// 		c.JSON(http.StatusCreated, gin.H{"inquerysets": inqueryset})
// 	} else {
// 		c.JSON(http.StatusBadRequest, gin.H{"status": "Please must update Title, Chr, Coordinate"})
// 	}
// }

// Delete a inqueryset
// func Delete(c *gin.Context) {
// 	db := c.MustGet("db").(*mgo.Database)
// 	query := bson.M{"dataset": c.Param("dataset")}
// 	err := db.C(models.CollectionInqueryset).Remove(query)
// 	if err != nil {
// 		c.Error(err)
// 	}
// 	c.Redirect(http.StatusMovedPermanently, "/inquerysets")
//
// 	var json models.Inqueryset
// 	c.BindJSON(&json) // This will infer what binder to use depending on the content-type header.
// 	if json.Dataset != c.Params.ByName("dataset") {
// 		c.JSON(http.StatusOK, gin.H{"status": "Delete [ " + c.Params.ByName("dataset") + " ]"})
// 	} else {
// 		c.JSON(http.StatusNotFound, gin.H{"status": "Please input right id"})
// 	}
// }
