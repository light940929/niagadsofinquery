// Package main is the CLI.
// You can use the CLI via Terminal.
// @APIVersion 0.0.1
// @APITitle title
// @APIDescription description
// @Contact hanjl@mail.med.upenn.edu
// @TermsOfServiceUrl http://...
// @License MIT NIAGADS
// @LicenseUrl http://osensource.org/licenses/MIT
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	"github.com/user/niagadsofinquery/db"
	"github.com/user/niagadsofinquery/gin_html_render"
	"github.com/user/niagadsofinquery/middlewares"
	"github.com/user/niagadsofinquery/models"
)

const (
	// Port at which the server starts listening
	Port       = "7000"
	MysqlDBUrl = "@tcp(localhost:3306)/test?charset=utf8&parseTime=true" //user:password@tcp(ip:port)/database //root:
)

var (
	mysupersecretpassword = "NIAGADSisAWESOME"
)

var service *gin.Engine

func dbWare() gin.HandlerFunc {

	db, err := db.NewDB(MysqlDBUrl)
	if err != nil {
		log.Panic(err)
		fmt.Printf("Can't connect to mysql")
	}
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func getEnv(name string, def string) string {
	val := os.Getenv(name)

	if val == "" {
		return def
	}

	return val
}

func setupMiddleware() {
	if os.Getenv("API_KEY") != "" {
		service.Use(middlewares.RequireAuth)
	}
}

func index(c *gin.Context) {
	content := gin.H{"Hello": "World"}
	c.JSON(200, content)
}

func Usertest(ctx *gin.Context) {
	// user, err := models.AllUsers(ctx)
	// if err != nil {
	// 	ctx.JSON(404, gin.H{"error": err})
	// 	return
	// }
	// ctx.JSON(200, gin.H{"users": user})

	user, err := models.FindUser(ctx)
	if err != nil {
		log.Print("userlogin: ", err)
		ctx.JSON(404, gin.H{"error": "error loging in"})
		return
	}
	ctx.JSON(200, gin.H{"users": user})
}

func UserLogin(ctx *gin.Context) {
	user, err := models.FindUser(ctx)
	if err != nil {
		log.Print("userlogin: ", err)
		ctx.JSON(404, gin.H{"error": "error loging in"})
		return
	}

	dbpassword := user.Password
	log.Print("userpassword: ", dbpassword)
	formpassword := ctx.PostForm("password")
	log.Print("userform: ", formpassword)

	if bcrypt.CompareHashAndPassword([]byte(dbpassword), []byte(formpassword+"niagads")) != nil {
		log.Print("userloginAuthorized: ", err)
		log.Print("password: ", formpassword)
		ctx.JSON(401, gin.H{"error": "User not Authorized"})
		return
	}
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	token.Claims["ID"] = user.Email
	token.Claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // expires in a hour
	tokenString, err2 := token.SignedString([]byte(mysupersecretpassword))
	if err2 != nil {
		log.Print("tokenString: ", tokenString)
		log.Print("err2: ", err2)
		ctx.JSON(500, gin.H{"error": "Problem generating token"})
	}
	user.Password = ""
	log.Print("time: ", user.UpdatedAt)
	ctx.JSON(200, gin.H{"user": user, "token": tokenString})
}

func UserAdd(ctx *gin.Context) {
	user, err := models.AddUser(ctx)
	if err != nil {
		log.Print("user: ", user)
		log.Print("err: ", err)
		ctx.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	token.Claims["ID"] = user.Email
	token.Claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // expires in a hour
	tokenString, err2 := token.SignedString([]byte(mysupersecretpassword))
	if err2 != nil {

		ctx.JSON(500, gin.H{"error": "Problem generating token"})
	}
	user.Password = ""
	ctx.JSON(200, gin.H{"user": user, "token": tokenString, "success": "New user added"})
}

func UserRemove(ctx *gin.Context) {
	_, err := models.RemoveUser(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	ctx.JSON(200, gin.H{"success": "User removed"})
}

func DatasetAdd(c *gin.Context) {

	dataset, err := models.CreateDataset(c)
	if err != nil {
		log.Print("dataset: ", dataset)
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	c.JSON(201, gin.H{"success": dataset})

}
func InquerysetAdd(c *gin.Context) {

	inqueryset, err := models.CreateInqueryset(c)
	if err != nil {
		log.Print("inqueryset: ", inqueryset)
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	c.JSON(201, gin.H{"success": inqueryset})

}

func PhenotypeAdd(c *gin.Context) {

	phenotype, err := models.CreatePhenotype(c)
	if err != nil {
		log.Print("phenotype: ", phenotype)
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	c.JSON(201, gin.H{"success": phenotype})

}
func GenotypeAdd(c *gin.Context) {

	genotype, err := models.CreateGenotype(c)
	if err != nil {
		log.Print("genotype: ", genotype)
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	c.JSON(201, gin.H{"success": genotype})

}

func DatasetsGet(c *gin.Context) {
	datasets, err := models.ListDatasets(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error", "message": err})
		return
	}
	c.JSON(200, gin.H{"success": datasets})

}
func InquerysetsGet(c *gin.Context) {
	inquerysets, err := models.ListInquerysets(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error", "message": err})
		return
	}
	c.JSON(200, gin.H{"success": inquerysets})

}
func PhenotypesGet(c *gin.Context) {
	phenotypes, err := models.ListPhenotypes(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error", "message": err})
		return
	}
	c.JSON(200, gin.H{"success": phenotypes})

}
func GenotypesGet(c *gin.Context) {
	genotypes, err := models.ListGenotypes(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error", "message": err})
		return
	}
	c.JSON(200, gin.H{"success": genotypes})

}

func DatasetGet(c *gin.Context) {
	dataset, err := models.GetDataset(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	c.JSON(200, gin.H{"success": dataset})

}

func InquerysetGet(c *gin.Context) {
	inqueryset, err := models.GetInqueryset(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	c.JSON(200, gin.H{"success": inqueryset})

}
func PhenotypeGet(c *gin.Context) {
	phenotype, err := models.GetPhenotype(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	c.JSON(200, gin.H{"success": phenotype})

}
func GenotypeGet(c *gin.Context) {
	genotype, err := models.GetGenotype(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	c.JSON(200, gin.H{"success": genotype})

}

func DatasetUpdate(c *gin.Context) {
	dataset, err := models.UpdateDataset(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	c.JSON(201, gin.H{"success": dataset})

}

func InquerysetUpdate(c *gin.Context) {
	inqueryset, err := models.UpdateInqueryset(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	c.JSON(201, gin.H{"success": inqueryset})

}

func PhenotypeUpdate(c *gin.Context) {
	phenotype, err := models.UpdatePhenotype(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	c.JSON(201, gin.H{"success": phenotype})

}
func GenotypeUpdate(c *gin.Context) {
	genotype, err := models.UpdateGenotype(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	c.JSON(201, gin.H{"success": genotype})

}

func DatasetRemove(c *gin.Context) {
	dataset, err := models.DeleteDataset(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	log.Print("datasetRemove", dataset)
	c.JSON(200, gin.H{"success": "dataset removed"})

}

func InquerysetRemove(c *gin.Context) {
	inqueryset, err := models.DeleteInqueryset(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	log.Print("inquerysetRemove", inqueryset)
	c.JSON(200, gin.H{"success": "inqueryset removed"})

}
func PhenotypeRemove(c *gin.Context) {
	phenotype, err := models.DeletePhenotype(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	log.Print("phenotypeRemove", phenotype)
	c.JSON(200, gin.H{"success": "phenotype removed"})

}
func GenotypeRemove(c *gin.Context) {
	genotype, err := models.DeleteGenotype(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	log.Print("genotypeRemove", genotype)
	c.JSON(200, gin.H{"success": "genotype removed"})

}

func main() {

	if os.Getenv("API_KEY") == "" {
		fmt.Println("WARNING: API_KEY is not set! No authentication will be required.")
	}

	// Configure
	router := gin.Default()
	router.Use(dbWare())
	fmt.Print(router.Use(dbWare()))
	setupMiddleware()

	// Do not print extra debugging information
	gin.SetMode("release")

	// Set html render options
	htmlRender := GinHTMLRender.New()
	htmlRender.Debug = gin.IsDebugging()
	htmlRender.Layout = "layouts/default"
	// htmlRender.TemplatesDir = "templates/" in default
	// htmlRender.Ext = ".html"               in default

	// Tell gin to use our html render
	//router.HTMLRender = htmlRender.Create()

	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true

	// Middlewares
	//router.Use(middlewares.Connect)
	//router.Use(middlewares.ErrorHandler)
	//router.Use(middlewares.RequireAuth)

	//login public routes
	router.GET("/", index)
	//router.GET("/users", Usertest)
	//router.POST("/users", Usertest)

	router.POST("/login", UserLogin)
	router.POST("/signup", UserAdd)

	// JWT PRIVATE
	/*
		Set this header in your request to get here.
		Authorization: Bearer `token`
	*/

	private := router.Group("/api")
	private.Use(jwt.Auth(mysupersecretpassword))
	private.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from private. These are phenotypes & genotypes APIs."})
	})
	private.DELETE("/users/:email", UserRemove)

	//datasets
	private.GET("/datasets", DatasetsGet)
	private.POST("/datasets", DatasetAdd)
	private.GET("/datasets/:name", DatasetGet)
	private.PUT("/datasets/:name", DatasetUpdate)
	private.DELETE("/datasets/:name", DatasetRemove)

	//inquerysets
	private.GET("/inqueries", InquerysetsGet)
	private.POST("/inqueries", InquerysetAdd)
	private.GET("/inqueries/:name", InquerysetGet)
	private.PUT("/inqueries/:name", InquerysetUpdate)
	private.DELETE("/inqueries/:name", InquerysetRemove)

	// phenotypes
	private.GET("/phenotypes", PhenotypesGet)
	private.POST("/phenotypes", PhenotypeAdd)
	private.GET("/phenotypes/:name", PhenotypeGet)
	private.PUT("/phenotypes/:name", PhenotypeUpdate)
	private.DELETE("/phenotypes/:name", PhenotypeRemove)

	// genotypes
	private.GET("/genotypes", GenotypesGet)
	private.POST("/genotypes", GenotypeAdd)
	private.GET("/genotypes/:name", GenotypeGet)
	private.PUT("/genotypes/:name", GenotypeUpdate)
	private.DELETE("/genotypes/:name", GenotypeRemove)

	// Start listening
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}
	router.Run(":" + port)
}
