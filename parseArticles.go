package main

import (
	"golang.org/x/net/html"
	"strings"
)

type Article struct {
	Title     string
	HeroImage string
	Summary   string
	Author    string
	WrittenOn string
	Draft     bool
	MdFile    string
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

func parseArticle(htmlData []byte) *Article {

	title, err := FindNthElement(htmlData, "h1", 1)
	if err != nil {
		panic(err)
	}
	hero, err := FindNthElement(htmlData, "img", 1)
	summary, err := FindNthElement(htmlData, "p", 2)

	author, err := FindNthElement(htmlData, "h4", 2)
	// remove the words "Written by: " from the author string
	author = strings.Replace(author, "Written by: ", "", 1)

	writtenOn, err := FindNthElement(htmlData, "h4", 1)
	// remove the words "Written on: " from the writtenOn string
	writtenOn = strings.Replace(writtenOn, "Written on: ", "", 1)

	s := string(htmlData)

	a := Article{
		Title:     title,
		Summary:   summary,
		HeroImage: hero,
		Author:    author,
		WrittenOn: writtenOn,
		Draft:     false,
		MdFile:    s,
	}
	// return a pointer to the article
	return &a
}
