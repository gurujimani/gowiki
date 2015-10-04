package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	wiki_list := make([]string, 3)
	fmt.Println("Inside index handler")
	files, _ := filepath.Glob("*.txt")
	for _, f := range files {
		wiki_list = append(wiki_list, f[:(len(f)-4)])
		//fmt.Println(f[:(len(f) - 4)])
	}
	fmt.Println(wiki_list)
	t, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, wiki_list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusNotFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	fmt.Println("Inside editHandler")
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside newhandler")
	t, err := template.ParseFiles("new.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside saveHandler")

	title := r.URL.Path[len("/save/"):]
	if title == "" {
		title = r.FormValue("title")
	}
	fmt.Println("Title is " + title)
	body := r.FormValue("body")
	fmt.Println("Body is " + body)
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	fmt.Println("Starting the webserver...\n")
	http.HandleFunc("/index/", indexHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/new/", newHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Stopping the webserver...\n")

}
