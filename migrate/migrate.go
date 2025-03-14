package main

import (
	"crudapp/initializers"
	"crudapp/models"
)

func init()  {
	initializers.LoadEnv()
	initializers.ConnectDB()
}
func main()  {
	initializers.DB.AutoMigrate(&models.Post{})
}