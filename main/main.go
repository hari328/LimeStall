package main

import (
	"net/http"
	"html/template"
	"os"
	"strings"
	"bufio"
	"github.com/LimeStall/viewmodels"
	"fmt"
)

func main() {
	templates := getTemplateCache()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		requestedFile := r.URL.Path[1:]
		requestedTemplate := templates.Lookup(requestedFile+".html")

		var context interface{} = nil

		if requestedFile == "home" {
			context = viewmodels.GetHome()
		}

		fmt.Println(context)
		if requestedTemplate != nil {
			requestedTemplate.Execute(w, context)
		} else {
			w.WriteHeader(404)
		}

	})

	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)
	http.ListenAndServe(":8000", nil)
}

func serveResource(w http.ResponseWriter, r *http.Request) {
	resourcePath := "public" + r.URL.Path

	var contentType string

	if strings.HasSuffix(resourcePath, ".css") {
		contentType = "text/css"
	}else if strings.HasSuffix(resourcePath, ".png") {
		contentType = "image/png"
	}else {
		contentType = "text/plain"
	}

	f, err := os.Open(resourcePath)

	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)
		br := bufio.NewReader(f)
		br.WriteTo(w)
	}else {
		w.WriteHeader(404)
	}
}

func getTemplateCache() *template.Template {
	result := template.New("templates")
	basePath := "templates"

	folder,_ := os.Open(basePath)
	rawPaths, _ := folder.Readdir(-1);

	actualPaths := new([]string);

	for _, childPath := range rawPaths {
		if ! childPath.IsDir() {
			*actualPaths = append(*actualPaths, basePath+"/"+childPath.Name())
		}
	}

	result.ParseFiles(*actualPaths...)
	return result
}