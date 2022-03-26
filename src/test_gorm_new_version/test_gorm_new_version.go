package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type UserInfo struct {
	ID       uint
	Name     string
	Password string
	Role     string
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:4466@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})

	fmt.Println(err)
	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)
}
