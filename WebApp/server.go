package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
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

//-------------------------------
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s", r.URL.Path[1:])
}

// func viewHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/view/"):]
// 	log.Println("viewHandler: title " + title)
// 	p, _ := loadPage(title)
// 	fmt.Fprintf(w, "<h1>%s</h1> <div>%s</div>", p.Title, p.Body)
// }

// func editHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/edit/"):]
// 	p, err := loadPage(title)
// 	if err != nil {
// 		p = &Page{Title: title}
// 	}
// 	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
// 		"<form action=\"/save/%s\" method=\"POST\">"+
// 		"<textarea name=\"body\">%s</textarea><br>"+
// 		"<input type=\"submit\" value=\"Save\">"+
// 		"</form>",
// 		p.Title, p.Title, p.Body)
// }

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	log.Println("viewHandler: title " + title)
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		log.Println("redirect to: " + "/edit/" + title)
		return
	}
	renderTemplate(w, "view", p)
	// t, _ := template.ParseFiles("view.html")
	// t.Execute(w, p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
	// t, _ := template.ParseFiles("edit.html")
	// t.Execute(w, p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func main() {
	fmt.Println("Server is running now...")

	/* Seed the page */
	p1 := &Page{Title: "FirstPage", Body: []byte("This is a sample Page.")}
	p1.save()

	/* Local load */
	//p2, _ := loadPage("FirstPage")
	//fmt.Println(string(p2.Body))

	/* Deprecated */
	//http.HandleFunc("/", handler)

	/* Handlers */
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	//http.HandleFunc("/save/", saveHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
