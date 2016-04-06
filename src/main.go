package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"structure"
	"os"
	"io/ioutil"
)

var port = "43000"

func main() {
	fmt.Println("Server starting")
	http.HandleFunc("/", indexPage)
	//http.Handle("/", http.FileServer(http.Dir("./public")))

	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", fs)
	fs = http.StripPrefix("/templates/", http.FileServer(http.Dir("templates")))
	http.Handle("/templates/", fs)

	http.HandleFunc("/update", update)
	http.HandleFunc("/index.html", indexPage)

	err := http.ListenAndServe(":" + port, nil)
	checkError(err)
}

func indexPage(rw http.ResponseWriter, req *http.Request){
	index, err := ioutil.ReadFile("src/templates/index.html")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Write([]byte(index))
}

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

func checkError(err error){
	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
		os.Exit(1)
	}
}