/*
go build -o links . && ./links
*/

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/seetohjinwei/links/page"
	"github.com/seetohjinwei/links/url"
	yaml "github.com/seetohjinwei/links/yamlparser"
)

func generateLinks(filePath string) []url.Url {
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error while reading yaml file: %v", err)
	}

	return yaml.Parse(contents)
}

func redirectHandler(link url.Url) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%q: redirecting to %q", link.Short, link.Full)
		http.Redirect(w, r, link.Full, http.StatusMovedPermanently)
	}
}

func handleLinks(mux *http.ServeMux, links []url.Url) {
	for _, link := range links {
		mux.HandleFunc("/"+link.Short, redirectHandler(link))
	}
}

const defaultYaml = "links.yaml"
const defaultPage = "public/index.html"
const port = ":8085"

var yamlFilePath string
var pageFilePath string

func getFlags() {
	flag.StringVar(&yamlFilePath, "data", defaultYaml, "file path for links, see 'links.yaml' for an example")
	flag.StringVar(&pageFilePath, "main", defaultPage, "file path for generated main page")
	flag.Parse()
}

func makeAbsolutePaths(path string) string {
	absFilePath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("error file not found %q: %v", yamlFilePath, err)
	}
	return absFilePath
}

func main() {
	getFlags()
	yamlFilePath = makeAbsolutePaths(yamlFilePath)
	pageFilePath = makeAbsolutePaths(pageFilePath)

	links := generateLinks(yamlFilePath)

	page.BuildAndGenerate(links, pageFilePath)

	mux := http.NewServeMux()
	// Everything else routes here.
	mux.Handle("/", http.FileServer(http.Dir("public")))
	handleLinks(mux, links)

	log.Printf("Starting on port: %v", port)
	http.ListenAndServe(port, mux)
}
