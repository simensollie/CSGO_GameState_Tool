package main

import (
    "fmt"
    "net/http"
    "log"
    "encoding/json"
    //"io/ioutil"
)

var port = "43000"

type Gamestate struct {
    ScoreCT int `json: "map.team_ct.score"`
    ScoreT int `json: "map.team_t.score"`
}

func update(rw http.ResponseWriter, req *http.Request){
    decoder := json.NewDecoder(req.Body)
    
    var g Gamestate
    
    err := decoder.Decode(&g)
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        rw.WriteHeader(http.StatusInternalServerError)
        return
    }
    log.Printf("Score CT: %v - Score T: %v", g.ScoreCT, g.ScoreT)
}

func main() {
    fmt.Println("Server starting")
    http.HandleFunc("/", update)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}