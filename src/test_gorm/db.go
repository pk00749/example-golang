package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"fmt"
)

// UserInfo 用户信息
type UserInfo struct {
	ID       uint
	Name     string
	Password string
	Role     string
}

func main() {
	db, err := gorm.Open("mysql", "root:4466@/test?charset=utf8&parseTime=True&loc=Local")
	fmt.Println(err)
	defer db.Close()

	db.SingularTable(true)

	db.AutoMigrate(&UserInfo{})

	// u1 := UserInfo{1, "a", "qwer1234", "admin"}
	// u2 := UserInfo{2, "b", "asdf1234", "admin"}

	// db.Create(&u1)
	// db.Create(&u2)

	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)
	var uu UserInfo
	db.Find(&uu, "name=?", "a")
	fmt.Printf("%#v\n", uu)

}
