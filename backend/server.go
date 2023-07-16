package main

import (
	"encoding/json"
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
		http.Redirect(w, r, "/home", http.StatusFound)
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

//func frontPageHandler(w http.ResponseWriter, r *http.Request) {
//	// read all files in markdown directory create array of values
//	files, err := os.ReadDir(`markdown/`)
//
//	if err != nil {
//		fmt.Println("error reading directory", err)
//		http.Error(w, "internal server error", http.StatusInternalServerError)
//		return
//	}
//
//	var listOfMostRecentArticles []ArticleWithAuthorName
//
//	for _, file := range files {
//
//		mdFile, err := os.ReadFile("markdown/" + file.Name())
//		if err != nil {
//			fmt.Println("error reading file", err)
//			http.Error(w, "internal server error", http.StatusInternalServerError)
//			return
//		}
//
//		mdToHTML := blackfriday.Run(mdFile)
//		_ = parseArticle(mdToHTML)
//
//		result := GetAllArticles()
//		listOfMostRecentArticles = result
//		fmt.Printf("length of : %v \n", len(listOfMostRecentArticles))
//
//	}
//
//	// return json of all articles
//
//}

func catchAllHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "404.html", nil)
}

func getAllArticlesHandler(w http.ResponseWriter, r *http.Request) {
	result := GetAllArticles()

	//return json of all articles
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Set this to the methods you want to allow (e.g. "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	// Set this to the headers you want to allow, e.g. "Authorization, Content-Type"
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		fmt.Println("error encoding json", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func getArticle(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/api/article/")
	fmt.Printf("id: %v \n", id)

	result, err := GetArticleById(id)

	if err != nil {
		fmt.Println("error getting article", err)
		http.Error(w, "Can not find article with the id", http.StatusNotFound)
		return
	} else {

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Set this to the methods you want to allow (e.g. "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Methods", "GET")

		// Set this to the headers you want to allow, e.g. "Authorization, Content-Type"
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		err = json.NewEncoder(w).Encode(result)

		if err != nil {
			fmt.Println("error encoding json", err)
			http.Error(w, "Intenal server error", http.StatusInternalServerError)
		}
	}

}

func main() {
	InitDB()
	InitAws()
	createBlogPosts()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images/"))))

	http.HandleFunc(view, viewHandler)
	//http.HandleFunc("/home", frontPageHandler)
	//http.HandleFunc("/", catchAllHandler)
	http.HandleFunc("/api/articles", getAllArticlesHandler)
	http.HandleFunc("/api/article/", getArticle)
	// TODO : make admin routes for uploading articles
	//http.HandleFunc(routes["view"], makeHandler(viewHandler))
	//http.HandleFunc(routes["edit"], makeHandler(editHandler))
	//http.HandleFunc(routes["save"], makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))

}
