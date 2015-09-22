package helpers
import (
	"net/http"
	"io/ioutil"
	"regexp"
)

func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := LoadPage(title)
	if err != nil {
		filename := "data/"+title + ".txt"
		ioutil.WriteFile(filename, []byte(""), 0600)
		http.Redirect(w, r, "/view/"+title, http.StatusFound)
		return
	}
	RenderTemplate(w, "view", p)
}

func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := LoadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	RenderTemplate(w, "edit", p)
}

func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.Redirect(w, r, "/view/Default", http.StatusTemporaryRedirect)
			//http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func DefaultHandler(w http.ResponseWriter, r *http.Request){
	http.Redirect(w, r, "/view/Default", http.StatusTemporaryRedirect)
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
