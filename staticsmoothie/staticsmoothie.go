package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/eknkc/amber"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

type PageStruct struct {
	Title  string
	Footer string
	Body   string
	Pages  []string
}

func readPagesDir() []string {
	dir, err := os.Open("./pages")
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	var pages []string
	re := regexp.MustCompile(".md")
	for _, fi := range fileInfos {
		pages = append(pages, re.ReplaceAllString(fi.Name(), ""))
	}
	return pages
}

func loadMarkdownPage(title string) ([]byte, error) {
	filename := title + ".md"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	output := blackfriday.MarkdownBasic(body)
	return output, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[1:]
	p, err := loadMarkdownPage(title)
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}

	file, _ := ioutil.ReadFile("./config.json")

	ps := PageStruct{
		"",
		"",
		string(p[0:]),
		readPagesDir()}

	json.Unmarshal(file, &ps)

	buff := bytes.NewBufferString("")

	compiler := amber.New()
	// Parse the input file
	cerr := compiler.ParseFile("./layout.amber")
	if cerr == nil {
		// Compile input file to Go template
		tmplHead, err := compiler.Compile()
		if err == nil {
			tmplHead.Execute(buff, ps)
			fmt.Fprintf(w, buff.String())
		} else {
			log.Fatal(err)
		}
	}
}

func main() {
	pH := http.HandlerFunc(viewHandler)
	http.ListenAndServe(":8080", pH)
}
