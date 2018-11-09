package main

import (
    "html/template"
    "log"
    "net/http"
    "os"
    "math/rand"
    "time"
)

type PageVariables struct {
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

func shuffle(arr []Meme) {
    t := time.Now()
    rand.Seed(int64(t.Nanosecond())) // no shuffling without this line

    for i := len(arr) - 1; i > 0; i-- {
        j := rand.Intn(i)
        arr[i], arr[j] = arr[j], arr[i]
    }
}

func Generate(w http.ResponseWriter, r *http.Request){
    memeList := getList()
    shuffle(memeList.Memes)
    r.ParseForm()

    url := captionMeme(memeList.Memes[0].Id, r.Form["caption"][0])

    HomePageVars := PageVariables{
      URL: url,
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