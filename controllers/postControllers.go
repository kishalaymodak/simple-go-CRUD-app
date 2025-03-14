package controllers

import (
	"crudapp/initializers"
	"crudapp/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	
	var body struct{
		Title string
		Body string
	}
	c.Bind(&body)
	
	post:=models.Post{Title:body.Title,Body:body.Body}
	result:= initializers.DB.Create(&post)
	
	if result.Error!=nil{
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Post": post,
	})
}

func PostINdex(c *gin.Context)  {
	var posts[]models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"Post": posts,
	})
}

func PostShow(c *gin.Context)  {
	id:=c.Param("id")
	
	var post models.Post
	
	
	initializers.DB.First(&post,id)
	
	if post.Title=="" && post.Body==""{
		c.JSON(200, gin.H{
			"message":"post not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"Post": post,
	})
}
func PostUpdate(c *gin.Context)  {
	id:=c.Param("id")
	var body struct{
		Title string
		Body string
	}
	c.Bind(&body)
	var post models.Post
	initializers.DB.First(&post,id)

	if post.Title=="" && post.Body==""{
		c.JSON(200, gin.H{
			"message":"post not found",
		})
		return
	}
	initializers.DB.Model(&post).Updates(models.Post{Title:body.Title,Body:body.Body})
	
	c.JSON(200, gin.H{
		"Post": post,
	})
}

func PostDelete(c *gin.Context)  {
	id:=c.Param("id")
	var post models.Post
	initializers.DB.First(&post,id)
	if post.Title=="" && post.Body==""{
		c.JSON(200, gin.H{
			"message":"post not found",
		})
		return
	}
	initializers.DB.Delete(&post,id)
	
	c.JSON(200, gin.H{
		"message": post,
	})
}