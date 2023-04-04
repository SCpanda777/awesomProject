package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}

type User struct {
	gorm.Model
	Name string `json:"name"`
}

var db *gorm.DB

// 创建帖子
func createPost(c *gin.Context) {
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Create(&post)
	c.JSON(200, gin.H{"data": post})
}

// 根据用户ID查询帖子
func getPostsByUserID(c *gin.Context) {
	var posts []Post
	user_id := c.Param("user_id")
	db.Where("user_id = ?", user_id).Find(&posts)
	c.JSON(200, gin.H{"data": posts})
}

func main() {
	dconn := "user:password@tcp(127.0.0.1:3306)/forum?charset=utf8mb4&parseTime=True&loc=Local" //password需要改为自己mysql的密码

	db, err := gorm.Open("mysql", dconn)
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Post{}, &User{})

	r := gin.Default()

	// 发布帖子路由
	r.POST("/posts", createPost)

	// 根据用户ID查询帖子路由
	r.GET("/posts/:user_id", getPostsByUserID)

	r.Run()
}
