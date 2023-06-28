package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

type router struct {
	path    string
	handler func(w http.ResponseWriter, r *http.Request)
}

var viewRoute = router{path: "/view/", handler: viewHandler}
var editRoute = router{path: "/edit/", handler: editHandler}

// var saveRoute = router{path: "/save/", handler: saveHandler}
var routes = map[string]router{
	"view": viewRoute,
	"edit": editRoute,
	//"save": editRoute,
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)

}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len(routes["view"].path):]
	p, err := loadPage(title)

	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusNotFound)
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len(routes["edit"].path):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}

func main() {
	http.HandleFunc(routes["view"].path, routes["view"].handler)
	http.HandleFunc(routes["edit"].path, routes["edit"].handler)
	//http.HandleFunc("/save/", saveHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
