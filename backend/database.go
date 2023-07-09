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
	AuthorID  uint
	Author    Author `gorm:"foreignKey:AuthorID"`
	WrittenOn string
	Draft     bool
	Content   string
}

type ArticleWithAuthorName struct {
	gorm.Model
	Title      string
	HeroImage  string
	Summary    string
	AuthorName string `gorm:"column:author_name"`
	WrittenOn  string
	Draft      bool
	Content    string
}

type Author struct {
	gorm.Model
	Name     string
	Email    string
	Articles []Article `gorm:"foreignKey:AuthorID"`
}

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=Admin password=example  port=8001 sslmode=disable TimeZone=mst"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Drop tables if they exist
	err = db.Migrator().DropTable(&Author{}, &Article{})
	if err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
	}

	// Create new tables
	err = db.AutoMigrate(&Author{}, &Article{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	DB = db
}

func CreateArticle(a Article) {
	// Check if article already exists before creating it checking by title
	var article Article
	result := DB.Where("title = ?", a.Title).First(&article)

	if result.Error != nil {
		result = DB.Create(&a)
		if result.Error != nil {
			fmt.Printf("Failed to create article: %v\n ", result.Error)
		}
		return
	} else {
		fmt.Printf("Article already exists with title:%v\n  ", a.Title)
		return
	}
}

func GetArticle(title string) Article {
	var article Article
	result := DB.Where("title = ?", title).First(&article)
	if result.Error != nil {
		fmt.Printf("Failed to get article: %v\n", result.Error)

	}
	return article
}

func GetAllArticles() []ArticleWithAuthorName {
	var articles []ArticleWithAuthorName
	result := DB.Model(&Article{}).
		Select("articles.*, authors.name as author_name").
		Joins("left join authors on authors.id = articles.author_id").
		Limit(10).
		Find(&articles)

	if result.Error != nil {
		fmt.Printf("Failed to get all articles: %v\n", result.Error)
		return nil
	}
	return articles
}

func DeleteArticle(title string) {
	var article Article
	result := DB.Where("title = ?", title).Delete(&article)
	if result.Error != nil {
		log.Fatalf("Failed to delete article: %v", result.Error)
		return
	}
}
