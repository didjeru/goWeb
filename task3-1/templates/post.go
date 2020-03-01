package templates

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/russross/blackfriday"
)

// Post - struct
type Post struct {
	Title   string
	Body    template.HTML
	ModTime int64
}

// PostArray - struct
type PostArray struct {
	Items map[string]Post
	sync.RWMutex
}

// NewPostArray - array of post
func NewPostArray() *PostArray {
	p := PostArray{}
	p.Items = make(map[string]Post)
	return &p
}

// Get Загружает markdown-файл и конвертирует его в HTML
// Возвращает объект типа Post
// Если путь не существует или является каталогом, то возвращаем ошибку
func (p *PostArray) Get(md string) (Post, int, error) {
	info, err := os.Stat(md)
	if err != nil {
		if os.IsNotExist(err) {
			// файл не существует
			return Post{}, 404, err
		}
		return Post{}, 500, err
	}
	if info.IsDir() {
		// не файл, а папка
		return Post{}, 404, fmt.Errorf("dir")
	}
	val, ok := p.Items[md]
	if !ok || (ok && val.ModTime != info.ModTime().UnixNano()) {
		p.RLock()
		defer p.RUnlock()
		fileread, _ := ioutil.ReadFile(md)
		lines := strings.Split(string(fileread), "\n")
		title := string(lines[0])
		body := strings.Join(lines[1:len(lines)], "\n")
		body = string(blackfriday.MarkdownCommon([]byte(body)))
		p.Items[md] = Post{title, template.HTML(body), info.ModTime().UnixNano()}
	}
	Post := p.Items[md]
	return Post, 200, nil
}
