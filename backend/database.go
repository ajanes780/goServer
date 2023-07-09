package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Article struct {
	gorm.Model
	Title     string `gorm:"uniqueIndex"`
	HeroImage string
	Summary   string
	Author    string
	WrittenOn string
	Draft     bool
	Content   string
}

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=Admin password=example  port=8001 sslmode=disable TimeZone=mst"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	err = db.AutoMigrate(&Article{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = db.AutoMigrate(&Article{})

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	DB = db
}

func CreateArticle(a Article) {
	// Check if article already exists before creating it cheecking by title
	var article Article
	result := DB.Where("title = ?", a.Title).First(&article)

	if result.Error != nil {
		result = DB.Create(&a)
		if result.Error != nil {
			fmt.Println("Failed to create article: ", result.Error)
		}
	} else {
		fmt.Println("Article already exists with title: ", a.Title)
		return
	}
}

func GetArticle(title string) Article {
	var article Article
	result := DB.Where("title = ?", title).First(&article)
	if result.Error != nil {
		log.Fatalf("Failed to get article: %v", result.Error)

	}
	return article
}

func GetAllArticles() []Article {
	// TODO add pagination
	var articles []Article
	result := DB.Find(&articles)
	fmt.Println("result: ", result)

	if result.Error != nil {
		log.Fatalf("Failed to get all articles: %v", result.Error)
	}
	return articles
}
