package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	xd "github.com/casbin/xorm-adapter"
)

// Interceptor
func myAuth(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		obj := c.Request.URL.RequestURI()
		act := c.Request.Method
		sub := "root"

		// judge if policy exist
		if ok := e.Enforce(sub, obj, act); ok {
			log.Println("Check successfully")
			c.Next()
		} else {
			log.Println("Sorry, check failed")
			c.Abort()
		}
	}
}

func main() {
	a := xd.NewAdapter("mysql", "root:4466@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true&loc=Local", true)
	e := casbin.NewEnforcer("./auth_model.conf", a)

	e.LoadPolicy()

	r := gin.New()

	r.POST("/api/v1/add", func(c *gin.Context) {
		log.Println("add a policy")
		if ok := e.AddPolicy("root", "/api/v1/hello", "GET"); !ok {
			log.Println("policy exists")
		} else {
			log.Println("add successfully")
		}
	})

	r.Use(myAuth(e))

	r.GET("/api/v1/hello", func(c *gin.Context) {
		fmt.Println("hello")
	})

	r.Run(":8888")
}
