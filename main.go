/*
go build -o links . && ./links
*/

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"

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
const defaultPage = "bin/index.html"
const port = ":8085"

var yamlFilePath string
var pageFilePath string

func getFlags() {
	flag.StringVar(&yamlFilePath, "data", defaultYaml, "file path for links, see 'links.yaml' for an example")
	flag.StringVar(&pageFilePath, "main", defaultPage, "file path for generated main page")
	flag.Parse()
}

func main() {
	getFlags()

	links := generateLinks(yamlFilePath)

	page.BuildAndGenerate(links, pageFilePath)

	mux := http.NewServeMux()
	// Everything else routes here.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, defaultPage)
	})
	handleLinks(mux, links)

	log.Printf("Starting on port: %v", port)
	http.ListenAndServe(port, mux)
}
