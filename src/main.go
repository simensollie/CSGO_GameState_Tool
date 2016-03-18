package main

import (
    "fmt"
    "net/http"
    "log"
    "encoding/json"
    //"io/ioutil"
)

var port = "43000"

type Match struct {
	Mode string 	`json: "map.mode"`
	Name string 	`json: "map.name"`
	Phase string 	`json: "map.phase"`
	Round int	`json: "map.round"`
	CT_Score int 	`json: "map.team_ct.score"`
	T_Score int 	`json: "map.team_t.score"`
}

type Round struct  {
	Phase string	`json: "round.phase"`
	Bomb string	`json: "round.bomb"`
	Winner string	`json: "round.win_team"`
}

type Gamestate struct {
    	Match []Match
	Round []Round
}

func update(rw http.ResponseWriter, req *http.Request){
    var g Gamestate
    
    err := json.NewDecoder(req.Body).Decode(&g)
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        rw.WriteHeader(http.StatusInternalServerError)
        return
    }

    log.Printf("Mapname: %s\n", )
    log.Printf("Round phase: %s\n", g.RoundPhase)
    log.Printf("Bomb planted: %s\n", g.BombPlanted)
}

func main() {
    fmt.Println("Server starting")
    http.HandleFunc("/", update)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}