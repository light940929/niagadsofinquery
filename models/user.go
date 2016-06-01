package models

import (
	"database/sql"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at,omitempty`  //NullTime
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
	name := ctx.PostForm("name")
	user := new(User)
	err := db.QueryRow("SELECT * FROM users WHERE name=?;", name).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Print("findusers: ", err)
		return nil, err
	}
	return user, nil
}

func AddUser(ctx *gin.Context) (*User, error) {
	db := ctx.MustGet("db").(*sql.DB)
	user := new(User)
	user.Name = ctx.PostForm("name")
	user.Email = ctx.PostForm("email")
	b, err := bcrypt.GenerateFromPassword([]byte(ctx.PostForm("password")+"niagads"), bcrypt.DefaultCost)
	//n := bytes.IndexByte(b, 0)
	user.Password = string(b) //b[:n]
	//err = db.QueryRow("INSERT INTO users(name,email,password) VALUES(?, ?, ?);", &user.Name, &user.Email, &user.Password).Scan(&user.Id)
	stmt, err := db.Prepare("INSERT INTO users(name,email,password) VALUES(?, ?, ?);")
	defer stmt.Close()
	stmt.Exec(&user.Name, &user.Email, &user.Password)
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
	`created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	unique(`email`),
    PRIMARY KEY (`id`)
);

INSERT INTO `users` (name, email, password) VALUES ('hannah','hanjl@mail.med.upenn.edu','hanjenlin'), ('niagads', 'support@niagads.org', 'niagads2016');
*/
