// Package middlewares contains gin middlewares
// Usage: router.Use(middlewares.Connect)
package middlewares

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/satori/go.uuid"
)

// Connect middleware clones the database session for each request and
// makes the `db` object available for each handler
// func Connect(c *gin.Context) {
//
// 	//c.Set("db", s.DB(db.Mongo.Database))
//
// 	c.Set("db", db)
// 	c.Next()
// }

// ErrorHandler is a middleware to handle errors encountered during requests
// func ErrorHandler(c *gin.Context) {
// 	c.Next()
//
// 	// TODO: Handle it in a better way
// 	if len(c.Errors) > 0 {
// 		c.HTML(http.StatusBadRequest, "400", gin.H{"errors": c.Errors})
// 	}
// }

// Check response
func ErrorResponse(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	//c.JSON(401, gin.H{"message": "Empty Session.", "status": 401})
	c.JSON(code, resp)
	//c.Abort(code)

}

// Require valid api authentication token
func RequireAuth(c *gin.Context) {
	token := c.Request.FormValue("api_key")

	if token == "" {
		ErrorResponse(401, "Missing 'api_key' patameter", c)
		c.Abort()
		return
	}

	if token != os.Getenv("API_KEY") {
		ErrorResponse(401, "Invalid api key", c)
		c.Abort()
		return
	}

	c.Next()
}

func RevisionMiddleware(c *gin.Context) {
	// Revision file contents will be only loaded once per process
	data, err := ioutil.ReadFile("REVISION")

	// If we cant read file, just skip to the next request handler
	// This is pretty much a NOOP middlware :)
	if err != nil {
		c.Next()
	}

	// Clean up the value since it could contain line breaks
	revision := strings.TrimSpace(string(data))

	// Set out header value for each response
	c.Writer.Header().Set("X-Revision", revision)
	c.Next()

}

func RequestIdMiddleware(c *gin.Context) {

	c.Writer.Header().Set("X-Request-Id", uuid.NewV4().String())
	c.Next()

}
