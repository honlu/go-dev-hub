package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	// 加载环境变量，从.env文件中获取数据库连接字符串
	var err error
	err = godotenv.Load("./.env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("POSTGRES_DSN") // 从环境变量中获取数据库连接字符串
	if dsn == "" {
		dsn = "host=localhost user=admin password=password_123456 dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	}

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	// 自动迁移数据表
	db.AutoMigrate(&Blog{}, &Question{})
}

func main() {
	initDB()

	// 设置gin模式
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// 路由分组
	admin := r.Group("/api/admin")
	{
		// 博客管理接口
		admin.POST("/blogs", createBlog)
		admin.GET("/blogs", getBlogs)
		admin.GET("/blogs/:id", getBlog)
		admin.PUT("/blogs/:id", updateBlog)
		admin.DELETE("/blogs/:id", deleteBlog)

		// 题库管理接口
		admin.POST("/questions", createQuestion)
		admin.GET("/questions", getQuestions)
		admin.GET("/questions/:id", getQuestion)
		admin.PUT("/questions/:id", updateQuestion)
		admin.DELETE("/questions/:id", deleteQuestion)

		// 获取分类枚举
		admin.GET("/categories", getCategories)
	}

	web := r.Group("/api/web")
	{
		// 博客只读接口
		web.GET("/blogs", getPublicBlogs)
		web.GET("/blogs/:id", getPublicBlog)

		// 题库只读接口
		web.GET("/questions", getPublicQuestions)
		web.GET("/questions/:id", getPublicQuestion)

		// 获取分类枚举
		web.GET("/categories", getCategories)
	}

	r.Run(":8080")
}
