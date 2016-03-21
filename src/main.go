package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"structure"
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

	//log.Printf("Bomb: %s\n", g.Round.Bomb)
	//log.Printf("CT Score: %v - T Score: %v\n", g.Map.TeamCT, g.Map.TeamT)
	log.Printf("Phase: %v \nBomb: %v \nWinner: %v", g.Round.Phase, g.Round.Bomb, g.Round.RoundWinner)
}

func main() {
	fmt.Println("Server starting")
	http.HandleFunc("/", update)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}