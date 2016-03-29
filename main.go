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
	"net/http"
	"os"

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

func init() {
	db.Connect()
}

func main() {

	// Configure
	router := gin.Default()

	// Set html render options
	htmlRender := GinHTMLRender.New()
	htmlRender.Debug = gin.IsDebugging()
	htmlRender.Layout = "layouts/default"
	// htmlRender.TemplatesDir = "templates/" // default
	// htmlRender.Ext = ".html"               // default

	// Tell gin to use our html render
	router.HTMLRender = htmlRender.Create()

	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true

	// Middlewares
	router.Use(middlewares.Connect)
	router.Use(middlewares.ErrorHandler)

	// Statics
	router.Static("/public", "./public")

	// Routes
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/phenotypes")
	})

	// phenotypes
	router.GET("/newphenotypes", phenotypes.New)

	router.GET("/phenotypes", phenotypes.List)
	router.POST("/phenotypes", phenotypes.Create)
	router.GET("/phenotypes/:_id", phenotypes.Get)
	router.PUT("/phenotypes/:_id", phenotypes.Update)
	router.DELETE("/phenotypes/:_id", phenotypes.Delete)

	// genotypes
	router.GET("/newgenotypes", genotypes.New)

	router.GET("/genotypes", genotypes.List)
	router.POST("/genotypes", genotypes.Create)
	router.GET("/genotypes/:_id", genotypes.Get)
	router.PUT("/genotypes/:_id", genotypes.Update)
	router.DELETE("/genotypes/:_id", genotypes.Delete)

	// Start listening
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}
	router.Run(":" + port)
}
