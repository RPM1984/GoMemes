package main

import (
    "net/http"
    "io/ioutil"
    "log"
    "net/url"
    "encoding/json"
)

type GetMemesResponse struct {    
    Data MemeList `json:"data"`
}

type CaptionMemeResponse struct {    
    Data CaptionedMeme `json:"data"`
}

type MemeList struct {    
    Memes []Meme `json:"memes"`
}

type CaptionedMeme struct {    
    URL     string  `json:"url"`
}

type Meme struct {
    Id      string  `json:"id"`
    Name    string  `json:"name"`
    URL     string  `json:"url"`
}

func readGetList(body []byte) (MemeList, error) {
    var q MemeList
    var s = new(GetMemesResponse)
    err := json.Unmarshal(body, &s)
    if(err != nil){
        log.Print("error: ", err)
    }

    q = s.Data

    return q, err
}

func getList() (MemeList) {   
    var q MemeList
    response, err := http.Get("https://api.imgflip.com/get_memes")
    if err != nil {
        log.Print("error: ", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        s, err := readGetList([]byte(data))
        _ = err
        q = s
    }

    return q
}

func readCaptionMeme(body []byte) (string, error) {
    var q string
    var s = new(CaptionMemeResponse)
    err := json.Unmarshal(body, &s)
    if(err != nil){
        log.Print("error: ", err)
    }

    q = s.Data.URL

    return q, err
}

func captionMeme(templateId string, caption string) (string) {   
    var q string

    formData := url.Values{
        "template_id": {templateId},
        "text0": {caption},
        "username": {"gomeme"},
        "password": {"gomeme2018"},
	}
    
    response, err := http.PostForm("https://api.imgflip.com/caption_image", formData)

    if err != nil {
        log.Print("error: ", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        s, err := readCaptionMeme([]byte(data))
        _ = err
        q = s
    }

    return q
}