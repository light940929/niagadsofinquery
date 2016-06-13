package models

import (
	"database/sql"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

const (
	CollectionUser = "users"
)

type User struct {
	Id        int        `json:"id",omitempty`
	Name      string     `json:"name" `                //binding:"required"
	Email     string     `json:"email" `               //binding:"required"
	Password  string     `json:"password" `            //binding:"required"
	CreatedAt *time.Time `json:"created_at,omitempty"` //NullTime
	UpdatedAt *time.Time `json:"updated_at,omitempty"` //NullTime
}

func AllUsers(ctx *gin.Context) ([]*User, error) {
	db := ctx.MustGet("db").(*sql.DB)
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Print("query: ", err)
		return nil, err
	}
	//defer rows.Close()

	var users []*User //users := make([]*User, 0)

	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Print("rows: ", err)
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func FindUser(ctx *gin.Context) (*User, error) {
	db := ctx.MustGet("db").(*sql.DB)
	user := new(User)
	err := ctx.Bind(user)
	if err != nil {
		log.Print("err: ", err)
		return user, nil
	}
	log.Print("bindjson: ", &user)
	log.Print("testname: ", user.Name)
	log.Print("testpassword: ", user.Password)

	//err := db.QueryRow("SELECT * FROM users WHERE name=?;", name).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	errjson := db.QueryRow("SELECT * FROM users WHERE name=?;", user.Name).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if errjson != nil {
		log.Print("findusersjson: ", errjson)
		return nil, errjson
	}

	return user, nil
}

func AddUser(ctx *gin.Context) (*User, error) {
	db := ctx.MustGet("db").(*sql.DB)
	user := new(User)
	err := ctx.Bind(user)
	if err != nil {
		log.Print("err: ", err)
		return user, nil
	}

	log.Print("bindjson: ", &user)
	name := ctx.PostForm("name")
	email := ctx.PostForm("email")
	b, err := bcrypt.GenerateFromPassword([]byte(ctx.PostForm("password")+"niagads"), bcrypt.DefaultCost)
	password := string(b) //b[:n]
	log.Print("passwordb: ", password)
	log.Print("user: ", user)
	//err = db.QueryRow("INSERT INTO users(name,email,password) VALUES(?, ?, ?);", &user.Name, &user.Email, &user.Password).Scan(&user.Id)
	stmt, err := db.Prepare("INSERT INTO users(name,email,password) VALUES(?, ?, ?);")
	defer stmt.Close()
	log.Print("username: ", name, "useremail: ", email, "userpassword: ", password)
	//ctx.BindJSON(&user)
	stmt.Exec(user.Name, user.Email, password)
	if err != nil {
		return nil, err
	}
	//user.Password = ""
	return user, nil
}

func RemoveUser(ctx *gin.Context) (bool, error) {
	db := ctx.MustGet("db").(*sql.DB)
	email := ctx.Param("email")
	user, err := db.Exec("DELETE FROM users WHERE email=?", email)
	if err != nil {
		log.Print("errToRemoveUser: ", err)
		return false, err
	}
	log.Print("removeUser", user)
	return true, nil
}

/*
CREATE TABLE `users` (
	`id` BIGINT NOT NULL AUTO_INCREMENT,
	`name` char(50) NOT NULL,
	`email` varchar(255) NOT NULL,
	`password` char(200) NOT NULL,
	`created_at` TIMESTAMP NOT NULL default 0,
	`updated_at` TIMESTAMP ON UPDATE now(),
	unique(`email`),
    PRIMARY KEY (`id`)
);

INSERT INTO `users` (name, email, password) VALUES ('hannah','hanjl@mail.med.upenn.edu','hanjenlin'), ('niagads', 'support@niagads.org', 'niagads2016');
*/
