package main

import (
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type GetMemesResponse struct {    
    Data MemeList `json:"data"`
}

type MemeList struct {    
    Memes []Meme `json:"memes"`
}

type Meme struct {
    Id      string  `json:"id"`
    Name    string  `json:"name"`
    URL     string  `json:"url"`
}

func getMemes(body []byte) (MemeList, error) {
    var q MemeList
    var s = new(GetMemesResponse)
    err := json.Unmarshal(body, &s)
    if(err != nil){
        // todo - return error
    }

    q = s.Data

    return q, err
}

func getList() (MemeList) {   
    var q MemeList
    response, err := http.Get("https://api.imgflip.com/get_memes")
    if err != nil {
        // todo - return error
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        s, err := getMemes([]byte(data))
        _ = err
        q = s
    }

    return q
}