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

//var routes = map[string]string{
//	"view": "/view/",
//	"edit": "/edit/",
//	"save": "/save/",
//}
//
//var route = map[string]string{
//	"home":  "/",
//	"view":  "/view/",
//	"about": "/about/",
//}

var tmpl = template.Must(template.ParseFiles(
	"templates/view.html",
	"templates/header.html",
	"templates/footer.html",
	"templates/404.html",
))

//var validPath = regexp.MustCompile("^/(edit|save|view|)/([a-zA-Z0-9]+)$")

//func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		//m := validPath.FindStringSubmatch(r.URL.Path)
//		//p := loadPage(r.URL.Path)
//
//		if m == nil {
//			http.NotFound(w, r)
//			return
//		}
//
//		fn(w, r, m[2])
//	}
//}

//func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
//	m := validPath.FindStringSubmatch(r.URL.Path)
//	if m == nil {
//		http.NotFound(w, r)
//		return "", errors.New("invalid Page Title")
//	}
//	return m[2], nil // The title is the second subexpression.
//}

//func (p *Page) save() error {
//	filename := p.Title + ".txt"
//	return os.WriteFile(filename, p.Body, 0600)
//}

//func loadPage(title string) (*Page, error) {
//	filename := title + ".md"
//	body, err := os.ReadFile(filename)
//
//	if err != nil {
//		return nil, err
//	}
//	return &Page{Title: title, Body: body}, nil
//}

//func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
//
//	err := templates.ExecuteTemplate(w, tmpl+".html", p)
//
//	if err != nil {
//		log.Printf("Error reading .md file: %v", err)
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//
//}

//func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
//
//	p, err := loadPage(title)
//
//	if err != nil {
//		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
//	}
//
//	renderTemplate(w, "view", p)
//}

//func editHandler(w http.ResponseWriter, r *http.Request, title string) {
//
//	p, err := loadPage(title)
//	if err != nil {
//		p = &Page{Title: title}
//	}
//
//	renderTemplate(w, "edit", p)
//}

//func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
//
//	body := r.FormValue("body")
//	p := &Page{Title: title, Body: []byte(body)}
//	err := p.save()
//
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	http.Redirect(w, r, "/view/"+title, http.StatusFound)
//}

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
		// todo : add a 404 page
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	mdToHTML := blackfriday.Run(mdFile)

	parseArticle(mdFile)

	fmt.Println(string(&a.mdFile))

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
	tmpl.ExecuteTemplate(w, "404.html", nil)
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
