package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
)

func main() {
    var PORT string
    
    if PORT = os.Getenv("PORT"); PORT == "" {
        PORT = "3001"
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        
        // Get a list of memes
        response, err := http.Get("https://api.imgflip.com/get_memes")
        if err != nil {
            // todo - return error
        } else {
            data, _ := ioutil.ReadAll(response.Body)
            fmt.Fprintf(w, string(data))
        }

        //fmt.Fprintf(w, "Hello World from path: %s\n", r.URL.Path)
    })

    http.ListenAndServe(":" + PORT, nil)
}