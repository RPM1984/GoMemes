package main

import (
    "net/http"
    "io/ioutil"
)

func getList() string {   
    response, err := http.Get("https://api.imgflip.com/get_memes")
    if err != nil {
        // todo - return error
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        return string(data)
    }

    return ""
}