package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"

	"./templates"

	"github.com/bmizerany/pat"
)

const (
	configFileName = "config.json"
)

var (
	// компилируем шаблоны, если не удалось, то выходим
	postTemplate  = template.Must(template.ParseFiles(path.Join("./templates", "layout.html"), path.Join("./templates", "post.html")))
	errorTemplate = template.Must(template.ParseFiles(path.Join("./templates", "layout.html"), path.Join("./templates", "error.html")))
	posts         = templates.NewPostArray()
)

type config struct {
	Port string `json:"port"`
}

func readConfig(ConfigName string) (x *config, err error) {
	var file []byte
	if file, err = ioutil.ReadFile(ConfigName); err != nil {
		return nil, err
	}

	file = bytes.TrimPrefix(file, []byte("\xef\xbb\xbf"))
	x = new(config)
	if err = json.Unmarshal(file, x); err != nil {
		return nil, err
	}
	return x, nil
}

func main() {
	cfg, err := readConfig(configFileName)
	if err != nil {
		log.Fatalln(err)
	}
	// для отдачи сервером статичных файлов из папки public/static
	fs := noDirListing(http.FileServer(http.Dir("/public/static")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	mux := pat.New()
	mux.Get("/:page", http.HandlerFunc(postHandler))
	mux.Get("/:page/", http.HandlerFunc(postHandler))
	mux.Get("/", http.HandlerFunc(postHandler))

	http.Handle("/", mux)
	log.Printf("Listening %s...", cfg.Port)
	log.Fatalln(http.ListenAndServe(":"+cfg.Port, nil))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	// Извлекаем параметр
	// Например, в http://127.0.0.1:8080/p1 page = "p1"
	// в http://127.0.0.1:8080/ page = ""
	page := params.Get(":page")
	// Путь к файлу (без расширения)
	// Например, posts/p1
	p := path.Join("./public/posts", page)

	var postMD string

	if page == "" {
		// если page пусто, то выдаем главную
		postMD = p + "/index.md"
	} else if page == "add" {
		log.Println(p)
		postMD = "./public/posts/index.md"
	} else if page == "edit" {
		log.Println(p)
		postMD = "./public/posts/index.md"
	} else {
		// если page не пусто, то считаем, что запрашивается файл
		// получим posts/p1.md
		postMD = p + ".md"
	}
	post, status, err := posts.Get(postMD)
	if err != nil {
		errorHandler(w, r, status)
		return
	}
	if err := postTemplate.ExecuteTemplate(w, "layout", post); err != nil {
		log.Println(err.Error())
		errorHandler(w, r, 500)
	}

}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	log.Printf("error %d %s %s\n", status, r.RemoteAddr, r.URL.Path)
	w.WriteHeader(status)
	if err := errorTemplate.ExecuteTemplate(w, "layout", map[string]interface{}{"Error": http.StatusText(status), "Status": status}); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
}

// обертка для http.FileServer, чтобы она не выдавала список файлов
// например, если открыть http://127.0.0.1:8080/static/,
// то будет видно список файлов внутри каталога.
// noDirListing - вернет 404 ошибку в этом случае.
func noDirListing(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") || r.URL.Path == "" {
			http.NotFound(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}
