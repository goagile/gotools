package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/bmizerany/pat"
	"github.com/russross/blackfriday"
)

var (
	post_template = template.Must(template.ParseFiles(
		path.Join("templates", "layout.html"),
		path.Join("templates", "post.html"),
	))

	error_template = template.Must(template.ParseFiles(
		path.Join("templates", "layout.html"),
		path.Join("templates", "error.html"),
	))
)

func main() {
	fs := http.FileServer(http.Dir("./public/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// http.HandleFunc("/", handler)

	mux := pat.New()
	mux.Get("/:page", http.HandlerFunc(handler))
	mux.Get("/:page/", http.HandlerFunc(handler))
	mux.Get("/", http.HandlerFunc(handler))
	http.Handle("/", mux)

	fmt.Println("Server was running at :3000")
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	page := params.Get(":page")
	p := path.Join("posts", page)
	var post_md string
	if page != "" {
		post_md = p + ".md"
	} else {
		post_md = p + "/index.md"
	}
	post, status, err := load_post(post_md)
	// post, status, err := load_post("posts/index.md")
	if err != nil {
		// http.Error(w, http.StatusText(status), status)
		errorHandler(w, r, status)
		return
	}
	// filebyt, err := ioutil.ReadFile("posts/index.md")
	// if err != nil {
	// 	log.Println("read .md error:", err)
	// 	http.Error(w, http.StatusText(500), 500)
	// }
	// lines := strings.Split(string(filebyt), "\n")
	// title := string(lines[0])
	// body := strings.Join(lines[1:len(lines)], "\n")
	// body = string(blackfriday.MarkdownCommon([]byte(body)))
	// post := Post{title, template.HTML(body)}
	if err := post_template.ExecuteTemplate(w, "layout", post); err != nil {
		log.Println(err)
		// http.Error(w, http.StatusText(500), 500)
		errorHandler(w, r, status)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if err := error_template.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Error":  http.StatusText(status),
		"Status": status,
	}); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
}

type Post struct {
	Title string
	Body  template.HTML
}

func load_post(md string) (Post, int, error) {
	info, err := os.Stat(md)
	if err != nil {
		if os.IsNotExist(err) {
			// файл не существует
			return Post{}, http.StatusNotFound, err
		}
	}
	if info.IsDir() {
		// не файл, а папка
		return Post{}, http.StatusNotFound, fmt.Errorf("dir")
	}
	fileread, err := ioutil.ReadFile(md)
	if err != nil {
		log.Println("read .md error:", err)
		return Post{}, http.StatusNotFound, nil
	}
	lines := strings.Split(string(fileread), "\n")
	title := string(lines[0])
	body := strings.Join(lines[1:len(lines)], "\n")
	body = string(blackfriday.MarkdownCommon([]byte(body)))
	post := Post{title, template.HTML(body)}
	return post, 200, nil
}
