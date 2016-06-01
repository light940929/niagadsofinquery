package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// var (
// Session stores mongo session
//Session *mgo.Session

// Mongo stores the mongodb connection string information
//Mongo *mgo.DialInfo
// )

const (
	// MongoDBUrl is the default mongodb url that will be used to connect to the
	// database.
	//MongoDBUrl = "mongodb://localhost:27017/test" //mongodb://username:password@l
	MysqlDBUrl = "root:@tcp(localhost:3306)/test/?parseTime=true" //?charset=utf8?
)

// Connect connects to mysql
func NewDB(dataSourceName string) (*sql.DB, error) {
	// Connect connects to mongodb
	// uri := os.Getenv("MONGODB_URL")
	//
	// if len(uri) == 0 {
	// 	uri = MongoDBUrl
	// }
	//
	// mongo, err := mgo.ParseURL(uri)
	// s, err := mgo.Dial(uri)
	// if err != nil {
	// 	fmt.Printf("Can't connect to mongo, go error %v\n", err)
	// 	panic(err.Error())
	// }
	// s.SetSafe(&mgo.Safe{})
	// fmt.Println("Connected to", uri)
	// Session = s
	// Mongo = mongo

	// Connect connects to mysql
	uri := os.Getenv("MysqlDBUrl")
	if len(uri) == 0 {
		uri = MysqlDBUrl
	}
	db, err := sql.Open("mysql", dataSourceName)
	log.Print(dataSourceName)
	log.Print(db.Stats())
	//log.Print(db.Query("SELECT * FROM users"))
	if err != nil {
		log.Fatal("err: ", err)
		fmt.Printf("Can't connect to mysql, error %v\n", err)
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	//fmt.Println("Connected to " + uri)

	//defer db.Close()

	return db, nil
}
