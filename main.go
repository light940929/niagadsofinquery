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
	"net/http"
	"os"
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	"github.com/user/niagadsofinquery/db"
	"github.com/user/niagadsofinquery/gin_html_render"
	"github.com/user/niagadsofinquery/handlers/genotypes"
	"github.com/user/niagadsofinquery/handlers/phenotypes"
	"github.com/user/niagadsofinquery/middlewares"
)

const (
	// Port at which the server starts listening
	Port = "7000"
)

var (
	mysupersecretpassword = "NIAGADSisAWESOME"
)

var service *gin.Engine

func init() {
	db.Connect()
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

func main() {

	if os.Getenv("API_KEY") == "" {
		fmt.Println("WARNING: API_KEY is not set! No authentication will be required.")
	}

	// Configure
	router := gin.Default()

	setupMiddleware()

	// Do not print extra debugging information
	gin.SetMode("release")

	// simulate some private data
	var secrets = gin.H{
		"hannah":  gin.H{"email": "hanjl@upenn.edu", "phone": "2158986986"},
		"niagads": gin.H{"email": "support@niagads.org", "phone": "2158983258"},
	}

	// Group using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"hannah":  "hanjenlin",
		"niagads": "niagads2016",
	}))

	// Set html render options
	htmlRender := GinHTMLRender.New()
	htmlRender.Debug = gin.IsDebugging()
	htmlRender.Layout = "layouts/default"
	// htmlRender.TemplatesDir = "templates/" in default
	// htmlRender.Ext = ".html"               in default

	// Tell gin to use our html render
	router.HTMLRender = htmlRender.Create()

	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true

	// Middlewares
	router.Use(middlewares.Connect)
	router.Use(middlewares.ErrorHandler)
	//router.Use(middlewares.RequireAuth)

	// URL: localhost:port/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret, "token_URL": "/api/oauth2/token/"})

		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	// JWT TOKEN
	public := router.Group("/api/oauth2/token")
	public.GET("/", func(c *gin.Context) {
		// Create the token
		token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
		// Set some claims
		token.Claims["ID"] = "NIAGADS"
		token.Claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
		// Sign and get the complete encoded token as a string
		tokenString, err := token.SignedString([]byte(mysupersecretpassword))
		if err != nil {
			c.JSON(500, gin.H{"message": "Could not generate token"})
		}
		c.JSON(200, gin.H{"token": tokenString})
	})

	// JWT PRIVATE
	private := router.Group("/api")
	private.Use(jwt.Auth(mysupersecretpassword))

	/*
		Set this header in your request to get here.
		Authorization: Bearer `token`
	*/

	private.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "These are phenotypes & genotypes api."})
	})

	// // Routes
	// private.GET("/", func(c *gin.Context) {
	// 	c.Redirect(http.StatusMovedPermanently, "/phenotypes")
	// })

	// phenotypes
	private.GET("/newphenotypes", phenotypes.New)

	private.GET("/phenotypes", phenotypes.List)
	private.POST("/phenotypes", phenotypes.Create)
	private.GET("/phenotypes/:_id", phenotypes.Get)
	private.PUT("/phenotypes/:_id", phenotypes.Update)
	private.DELETE("/phenotypes/:_id", phenotypes.Delete)

	// genotypes
	private.GET("/newgenotypes", genotypes.New)

	private.GET("/genotypes", genotypes.List)
	private.POST("/genotypes", genotypes.Create)
	private.GET("/genotypes/:_id", genotypes.Get)
	private.PUT("/genotypes/:_id", genotypes.Update)
	private.DELETE("/genotypes/:_id", genotypes.Delete)

	// Start listening
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}
	router.Run(":" + port)
}
