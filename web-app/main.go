package main

import (
  // "fmt"
  "io/ioutil"
  "log"
  "html/template"
  "regexp"
  "net/http"
)

type Page struct {
  Title string
  Body  []byte
}

func (p *Page) save() error {
  filename := p.Title + ".txt"
  // "io/ioutil"が必要
  return ioutil.WriteFile(filename, p.Body, 0600 )
}

func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return &Page{Title: title, Body: body}, nil
}

// キャッシュに入れる系
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page)  {
  // t, _ := template.ParseFiles(tmpl + ".html")
  // t.Execute(w, p)

  err := templates.ExecuteTemplate(w, tmpl+".html", p)
  if err := nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}


// viewHandlerってのがコントローラーっぽいことしてくれるみたいだね
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
  // /view/test
  p, err := loadPage(title)
  if err != nil {
    http.Redirect(w, r, "/edit/"+title, http.StatusFound)
  }
  // htmlを返してくれるらしいよ
  renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
  p, err := loadPage(title)
  if err != nil {
    p = &Page{Title: title}
  }
  // htmlを返してくれるらしいよ("html/template"が必要)
  renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
  body := r.FormValue("body")
  p := &Page{Title:title, Body: []byte(body)}
  err := p.save()
  if err != nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// 有効なパスであるかを識別
// "regexp"をimport
var validPath = regexp.MustCompile("^/(edit|view|save)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandleFunc{
  return func (w http.ResponseWriter, r *http.Request, string)  {
    m := validPath.FindStringSubmutch(r.URL.Path)
    if m == nil {
      http.NotFound(w, r)
      return
    }
    fn(w, r, m[2])
  }
}


func main()  {
  // p1 := &Page{Title: "test", Body: []byte("This is a sample Page.")}
  // p1.save()

  // p2, _ := loadPage(p1.Title)
  // fmt.Println(string(p2.Body))

  // HandleFuncってのがルーティングのイメージかな？
  http.HandleFunc("/view/", makeHandler(viewHandler))
  http.HandleFunc("/edit/", makeHandler(editHandler))
  http.HandleFunc("/save/", makeHandler(saveHandler))
  // "log"が必要
  log.Fatal(http.ListenAndServe(":3000", nil))

}
