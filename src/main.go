package main

import (
    "html/template"
    "log"
    "net/http"
    "os"
)

type PageVariables struct {
    Name    string
    Caption string
    URL     string
}

func main() {
    var PORT string
    
    if PORT = os.Getenv("PORT"); PORT == "" {
        PORT = "3001"
    }

    http.HandleFunc("/", HomePage)
    http.HandleFunc("/generate", Generate)
    http.ListenAndServe(":" + PORT, nil)
}

func HomePage(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("index.html")
    if err != nil {
  	  log.Print("template parsing error: ", err)
  	}
    err = t.Execute(w, nil)
    if err != nil {
  	  log.Print("template executing error: ", err)
  	}
}

func Generate(w http.ResponseWriter, r *http.Request){
    memeList := getList()
    r.ParseForm()

    HomePageVars := PageVariables{
      Name: memeList.Memes[0].Name,
      URL: memeList.Memes[0].URL,
      Caption: r.Form["caption"][0],
    }

    t, err := template.ParseFiles("meme.html")
    if err != nil {
  	  log.Print("template parsing error: ", err)
  	}
    err = t.Execute(w, HomePageVars)
    if err != nil {
  	  log.Print("template executing error: ", err)
  	}
}