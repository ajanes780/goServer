package main

import (
	"fmt"
	"github.com/russross/blackfriday/v2"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func getFirstChildData(n *html.Node) string {
	if n.FirstChild != nil {
		return n.FirstChild.Data
	}
	return ""
}

func getImgSrc(n *html.Node) string {
	for _, a := range n.Attr {
		if a.Key == "src" {
			return a.Val
		}
	}
	return ""
}

func findElement(n *html.Node, tag string) (string, bool) {
	if n.Type == html.ElementNode && n.Data == tag {
		if tag == "img" {
			return getImgSrc(n), true
		} else if tag == "p" {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Data == "img" {
					return getImgSrc(c), true
				}
			}
		}
		return getFirstChildData(n), true
	}
	return "", false
}

func FindNthElement(data []byte, tag string, n int) (string, error) {
	r := strings.NewReader(string(data))
	doc, err := html.Parse(r)

	if err != nil {
		return "", err
	}
	count := 0

	var f func(*html.Node) (string, bool)

	f = func(node *html.Node) (string, bool) {
		result, found := findElement(node, tag)

		if found {
			count++
			if count == n {
				return result, found
			}
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if result, ok := f(c); ok {
				return result, ok
			}
		}

		return "", false
	}

	result, _ := f(doc)
	return result, nil
}

func uploadImageToS3(p string) string {
	name := strings.Split(p, "/")
	name = name[len(name)-1:]
	path, err := AwsS3client.UploadFile(BUCKET_NAME, name[0], "/Users/aaronjanes/GolandProjects/goSever/backend"+p)

	if err != nil {
		fmt.Println("Error uploading file: ", err)
		//os.Exit(1)
	}

	fmt.Printf("File uploaded successfully at %s\n", path)
	return path
}

// https://questhenkart.medium.com/s3-image-uploads-via-aws-sdk-with-golang-63422857c548
func parseArticle(htmlData []byte) Article {

	title, err := FindNthElement(htmlData, "h1", 1)
	if err != nil {
		panic(err)
	}
	imageString, err := FindNthElement(htmlData, "img", 1)

	//hero := uploadImageToS3(imageString)
	hero := uploadImageToS3(imageString)

	summary, err := FindNthElement(htmlData, "p", 2)

	authorName, err := FindNthElement(htmlData, "h4", 2)
	// remove the words "Written by: " from the author string
	authorName = strings.Replace(authorName, "Written by: ", "", 1)

	writtenOn, err := FindNthElement(htmlData, "h4", 1)
	// remove the words "Written on: " from the writtenOn string
	writtenOn = strings.Replace(writtenOn, "Written on: ", "", 1)

	// look up the author by name
	var author Author
	result := DB.Where("name = ?", authorName).First(&author)

	if result.Error != nil {
		fmt.Println("Author not found, creating new author")
		author = Author{Name: authorName}
		result = DB.Create(&author)
		if result.Error != nil {
			fmt.Printf("Failed to create new author: %v\n", result.Error)
		}

	}

	s := string(htmlData)

	a := Article{
		Title:     title,
		Summary:   summary,
		HeroImage: hero,
		Author:    author,
		WrittenOn: writtenOn,
		Draft:     false,
		Content:   s,
	}
	// return a pointer to the article
	CreateArticle(a)
	return a
}

func createBlogPosts() []os.DirEntry {
	files, err := os.ReadDir(`markdown/`)
	if err != nil {
		fmt.Println("error reading directory", err)
		os.Exit(1)
	}

	for _, file := range files {

		mdFile, err := os.ReadFile("markdown/" + file.Name())
		if err != nil {
			fmt.Println("error reading file", err)
			os.Exit(1)
		}
		mdToHTML := blackfriday.Run(mdFile)
		_ = parseArticle(mdToHTML)
	}

	return files
}
