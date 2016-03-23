package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {
	serveWeb()
}

var themeName = getThemeName()
var staticPages = populateStaticPages()

func serveWeb() {

	gorillaRoute := mux.NewRouter()

	gorillaRoute.HandleFunc("/", serveContent)
	gorillaRoute.HandleFunc("/{page_alias}", serveContent)

	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8080", nil)
}

func serveContent(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	page_alias := urlParams["page_alias"]
	if page_alias == "" {
		page_alias = "home"
	}

	staticPage := staticPages.Lookup(page_alias + ".html")
	if staticPage == nil {
		staticPage = staticPages.Lookup("404.html")
		w.WriteHeader(404)
	}
	staticPage.Execute(w, nil)
}

func getThemeName() string {
	return "bs4"
}

func populateStaticPages() *template.Template {
	result := template.New("templates")
	templatePaths := new([]string)

	basePath := "pages"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()
	templatePathsRaw, _ := templateFolder.Readdir(-1)

	for _, pathInfo := range templatePathsRaw {
		log.Println(pathInfo.Name())
		*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
	}
	result.ParseFiles(*templatePaths...)
	return result
}
