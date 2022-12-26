// Helped out a bunch: https://betterprogramming.pub/how-to-generate-html-with-golang-templates-5fad0d91252

package page

import (
	"bufio"
	"bytes"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/seetohjinwei/links/url"
)

// Generate the page's content.
func Generate(templateDirectory string, mainTemplate string, templateNames []string, links []url.Url) []byte {
	templatePaths := make([]string, len(templateNames))

	for i, t := range templateNames {
		path, err := filepath.Abs(templateDirectory + "/" + t)
		if err != nil {
			log.Fatalf("error while parsing template %q: %v", t, err)
		}
		templatePaths[i] = path
	}

	templates, err := template.New("").ParseFiles(templatePaths...)
	if err != nil {
		log.Fatalf("error while parsing template: %v", err)
	}

	filteredLinks := url.RemoveDuplicates(links)

	var processed bytes.Buffer
	err = templates.ExecuteTemplate(&processed, mainTemplate, filteredLinks)
	if err != nil {
		log.Fatalf("error while executing template: %v", err)
	}

	return processed.Bytes()
}

func Build(content []byte, filePath string) {
	os.Remove(filePath)
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("error while creating page file: %v", err)
	}
	w := bufio.NewWriter(f)
	w.Write(content)
	w.Flush()
}

func BuildAndGenerate(links []url.Url, filePath string) {
	content := Generate("templates", "page", []string{"page.tmpl", "header.tmpl", "content.tmpl", "footer.tmpl"}, links)

	Build(content, filePath)
}
