package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    var PORT string
    
    if PORT = os.Getenv("PORT"); PORT == "" {
        PORT = "3001"
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Get meme list
        memeList := getList()
        fmt.Fprintf(w, memeList)
    })

    http.ListenAndServe(":" + PORT, nil)
}