package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"github.com/simensollie/CSGO_GameState_Tool/structure"
)

var port = "43000"

func update(rw http.ResponseWriter, req *http.Request) {
	var g structure.GameUpdates

	err := json.NewDecoder(req.Body).Decode(&g)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Bomb: %s\n", g.Round.Bomb)
	log.Printf("CT Score: %s - T Score: %s\n", g.Map.TeamCT.Score, g.Map.TeamT.Score)
}

func main() {
	fmt.Println("Server starting")
	http.HandleFunc("/", update)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}