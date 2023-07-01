package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

type Article struct {
	Title     string
	HeroImage string
	Summary   string
	mdFile    []byte
}

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
		}
		return getFirstChildData(n), true
	}
	return "", false
}

func FindFirstElement(data []byte, tag string) (string, error) {
	r := strings.NewReader(string(data))
	doc, err := html.Parse(r)
	if err != nil {
		return "", err
	}
	var f func(*html.Node) (string, bool)
	f = func(n *html.Node) (string, bool) {
		result, found := findElement(n, tag)
		if found {
			return result, found
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if result, ok := f(c); ok {
				return result, ok
			}
		}
		return "", false
	}
	result, _ := f(doc)
	return result, nil
}

func parseArticle(htmlData []byte) *Article {
	//htmlData := []byte(`<html><body><h1>Hello, world!</h1><h2>Some words on the subject</h2><img src="https://example.com/image.jpg" alt="An example image"><p>This is an example.</p></body></html>`)

	tags := []string{"h2", "h3", "h4", "h5", "h6", "p"}

	title, err := FindFirstElement(htmlData, "h1")
	if err != nil {
		panic(err)
	}
	hero, err := FindFirstElement(htmlData, "img")
	summary := ""

	for _, tag := range tags {
		result, err := FindFirstElement(htmlData, tag)

		if err != nil {
			panic(err)
		}

		if result != "" {
			summary += fmt.Sprintf("%s ", result)
		}

	}
	summary = strings.TrimSpace(summary)

	a := Article{
		Title:     title,
		Summary:   summary,
		HeroImage: hero,
		mdFile:    htmlData,
	}

	return &a
}
