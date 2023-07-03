package main

import (
	"fmt"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

// https://github.com/gomarkdown/markdown
// TODO: convert to use markdown done [x]
// TODO : render markdown file based on page name []
// TODO: add a home page with a list of pages []
// TODO: add header template to view template []

var view = "/view/"

type Page struct {
	Title string
	Body  []byte
}

var tmpl = template.Must(template.ParseFiles(
	"templates/view.html",
	"templates/header.html",
	"templates/footer.html",
	"templates/404.html",
))

func viewHandler(w http.ResponseWriter, r *http.Request) {

	//TODO: load page based on article/page name should also parse out the first image in the article and summary

	articleName := strings.Split(r.URL.Path, view)
	if len(articleName) < 2 {
		http.Redirect(w, r, view+"home", http.StatusFound)
		return
	}
	a := strings.Trim(articleName[1], " ")

	mdFile, err := os.ReadFile("markdown/" + a + ".md")

	// If the file doesn't exist, redirect to the home page
	if err != nil {
		// Todo : add a 404 page
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	mdToHTML := blackfriday.Run(mdFile)

	fmt.Println(string(mdToHTML))
	A := parseArticle(mdToHTML)

	fmt.Printf("%+v", A)

	// create article type
	// store data in article type
	// pass article type to template

	err = tmpl.Execute(w, template.HTML(mdToHTML))

	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func frontPageHandler(w http.ResponseWriter, r *http.Request) {

	// read all files in markdown directory create array of values
	files, err := os.ReadDir(`markdown/`)

	if err != nil {
		fmt.Println("error reading directory", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	var x []Article
	for _, file := range files {

		mdFile, err := os.ReadFile("markdown/" + file.Name())
		if err != nil {
			fmt.Println("error reading file", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		mdToHTML := blackfriday.Run(mdFile)
		a := parseArticle(mdToHTML)
		x = append(x, a)

	}
	fmt.Println("!!!!:x", x[2].Title)

	tmpl.ExecuteTemplate(w, "index.html", nil)

	//if err != nil {
	//	http.Error(w, "internal server error", http.StatusInternalServerError)
	//	return
	//}
}

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images/"))))

	http.HandleFunc(view, viewHandler)
	http.HandleFunc("/", frontPageHandler)

	// TODO : make admin routes for uploading articles
	//http.HandleFunc(routes["view"], makeHandler(viewHandler))
	//http.HandleFunc(routes["edit"], makeHandler(editHandler))
	//http.HandleFunc(routes["save"], makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))

}
