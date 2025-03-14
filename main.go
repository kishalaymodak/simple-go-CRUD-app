package main

import (
	"crudapp/controllers"
	"crudapp/initializers"
	"fmt"

	"github.com/gin-gonic/gin"
)
func init()  {
	initializers.LoadEnv()
	initializers.ConnectDB()
}
func main(){
	fmt.Println("hello world")
	
	r := gin.Default()
	r.GET("/posts", controllers.PostINdex)
	r.POST("/posts/:id", controllers.PostShow)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)
	r.POST("/posts", controllers.PostCreate)
	r.Run() 
  

}