package main

import (
    "fmt"
    "net/http"
    "log"
    "encoding/json"
    //"io/ioutil"
)

var port = "43000"

type GameUpdates struct{
	Provider struct{
		AppId int `json: "appid"`
		AppVersion int `json: "version"`
		Timestamp int `json: "timestamp"`
		 } `json: "provider"`
	Map struct{
		Mode string `json: "mode"`
		Name string `json: "name"`
		Phase string `json: "phase"`
		RoundNr int `json: "round"`
		TeamCT struct{
			Score int `json: "score"`
		       } `json: "team_ct"`
		TeamT struct{
			Score int `json: "score"`
		      } `json: "team_t"`
	    } `json: "map"`
	Round struct{
		Phase string `json: "phase"`
		Bomb string `json: "bomb"`
		RoundWinner string `json: "win_team"`
	      } `json: "round"`

}

func update(rw http.ResponseWriter, req *http.Request){
    var g GameUpdates
    
    err := json.NewDecoder(req.Body).Decode(&g)
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        rw.WriteHeader(http.StatusInternalServerError)
        return
    }

    log.Printf("Game Updates: %s\n", g)
}

func main() {
    fmt.Println("Server starting")
    http.HandleFunc("/", update)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}