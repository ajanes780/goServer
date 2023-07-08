package main

import (
	"fmt"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// TODO: page caching via redis []
// TODO: save articles to database []
// TODO: add admin page for uploading articles []
// TODO: login system cognito? []
// TODO: Dockerize app []
// TODO: Deploy to AWS container []
var view = "/view/"

func removePunctuation(s string) string {
	// Replace white space with hyphen
	s = strings.ReplaceAll(s, " ", "-")

	reg := regexp.MustCompile(`[^\w\s-]+`)
	s = reg.ReplaceAllString(s, "")

	// Replace multiple hyphens with a single hyphen
	regDoubleHyphen := regexp.MustCompile(`-{2,}`)
	s = regDoubleHyphen.ReplaceAllString(s, "-")

	// Convert to lowercase for standardization
	s = strings.ToLower(s)

	return s
}

var funcMap = template.FuncMap{
	"removePunctuation": removePunctuation,
}
var tmpl = template.Must(template.New("").Funcs(funcMap).ParseFiles(
	"templates/view.html",
	"templates/header.html",
	"templates/footer.html",
	"templates/404.html",
	"templates/index.html",
))

func viewHandler(w http.ResponseWriter, r *http.Request) {

	articleName := strings.Split(r.URL.Path, view)
	fmt.Println(articleName)
	if len(articleName) < 2 {
		http.Redirect(w, r, view+"home", http.StatusFound)
		return
	}

	a := strings.Trim(articleName[1], " ")

	mdFile, err := os.ReadFile("markdown/" + a + ".md")

	// If the file doesn't exist, redirect to the home page
	if err != nil {
		//TODO : add a 404 page
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	mdToHTML := blackfriday.Run(mdFile)

	err = tmpl.ExecuteTemplate(w, "view.html", template.HTML(mdToHTML))

	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
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

	var listOfMostRecentArticles []Article
	for _, file := range files {

		mdFile, err := os.ReadFile("markdown/" + file.Name())
		if err != nil {
			fmt.Println("error reading file", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		mdToHTML := blackfriday.Run(mdFile)
		parsedArticle := parseArticle(mdToHTML)
		listOfMostRecentArticles = append(listOfMostRecentArticles, parsedArticle)

	}

	tmpl.ExecuteTemplate(w, "index.html", listOfMostRecentArticles)

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
