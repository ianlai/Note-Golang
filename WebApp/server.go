package main

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

/* ParseFiles function takes any number of arguments for the file names */
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

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

	/* Not safe */
	//title := r.URL.Path[len("/view/"):]

	/* Safer */
	title, err := checkAndGetTitle(w, r)
	if err != nil {
		return
	}

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
	/* Not safe */
	//title := r.URL.Path[len("/view/"):]

	/* Safer */
	title, err := checkAndGetTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
	// t, _ := template.ParseFiles("edit.html")
	// t.Execute(w, p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	/* Not safe */
	//title := r.URL.Path[len("/view/"):]

	/* Safer */
	title, err := checkAndGetTitle(w, r)
	if err != nil {
		return
	}
	log.Println("saveHandler: title " + title)

	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err = p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func checkAndGetTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	log.Println("checkAndGetTitle: title " + m[2])
	return m[2], nil // The title is the second subexpression.
}

func main() {
	fmt.Println("Server is running now...")

	/* Seed the page */
	p1 := &Page{Title: "SamplePage", Body: []byte("This is a sample Page.")}
	p1.save()

	/* Local load */
	//p2, _ := loadPage("FirstPage")
	//fmt.Println(string(p2.Body))

	/* Deprecated */
	//http.HandleFunc("/", handler)

	/* Handlers */
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
