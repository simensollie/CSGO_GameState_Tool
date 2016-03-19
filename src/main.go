package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	//"io/ioutil"
)

var port = "43000"

type GameUpdates struct {
	Provider struct {
			 AppId      int `json: "appid"`
			 AppVersion int `json: "version"`
			 Timestamp  int `json: "timestamp"`
		 }`json: "provider"`
	Map      struct {
			 Mode    string `json: "mode"`
			 Name    string `json: "name"`
			 Phase   string `json: "phase"`
			 RoundNr int    `json: "round"`
			 TeamCT  struct {
					 Score int `json: "score"`
				 }`json: "team_ct"`
			 TeamT   struct {
					 Score int `json: "score"`
				 }`json: "team_t"`
		 } `json: "map"`
	Round    struct {
			 Phase       string `json: "phase"`
			 Bomb        string `json: "bomb"`
			 RoundWinner string `json: "win_team"`
		 }`json: "round"`
	Player   struct {
			 SteamID     string `json: "steamid"`
			 Name        string `json: "name"`
			 State       struct {
					     Health     int  `json: "health"`
					     Armor      int  `json: "armor"`
					     hasHelmet  bool `json: "helmet"`
					     Money      int  `json: "money"`
					     RoundKills int  `json: "round_kills"`
					     RoundHS    int  `json: "round_killhs"`
				     }`json: "state"`
			 Weapons     struct {
					     Weapon_0 struct {
							      Name     string `json: "name"`
							      Paintkit string `json: "paintkit"`
							      Type     string `json: "type"`
							      State    string `json: "state"`
						      }`json: "weapon_0"`
					     Weapon_1 struct {
							      Name        string `json: "name"`
							      Paintkit    string `json: "paintkit"`
							      Type        string `json: "type"`
							      AmmoClip    int 	 `json: "ammo_clip"`
							      AmmoClipMax int 	 `json: "ammo_clip_max"`
							      AmmoReserve int 	 `json: "ammo_reserve"`
							      State       string `json: "state"`
						      }`json: "weapon_1"`
					     Weapon_2 struct {
							      Name        string `json: "name"`
							      Paintkit    string `json: "paintkit"`
							      Type        string `json: "type"`
							      AmmoClip    int 	 `json: "ammo_clip"`
							      AmmoClipMax int 	 `json: "ammo_clip_max"`
							      AmmoReserve int 	 `json: "ammo_reserve"`
							      State       string `json: "state"`
						      }`json: "weapon_2"`
					     Weapon_3 struct {
							      Name        string `json: "name"`
							      Paintkit    string `json: "paintkit"`
							      Type        string `json: "type"`
							      AmmoClip    int 	 `json: "ammo_clip"`
							      AmmoClipMax int 	 `json: "ammo_clip_max"`
							      AmmoReserve int 	 `json: "ammo_reserve"`
							      State       string `json: "state"`
						      }`json: "weapon_3"`
					     Weapon_4 struct {
							      Name        string `json: "name"`
							      Paintkit    string `json: "paintkit"`
							      Type        string `json: "type"`
							      AmmoClip    int 	 `json: "ammo_clip"`
							      AmmoClipMax int 	 `json: "ammo_clip_max"`
							      AmmoReserve int 	 `json: "ammo_reserve"`
							      State       string `json: "state"`
						      }`json: "weapon_4"`
					     Weapon_5 struct {
							      Name        string `json: "name"`
							      Paintkit    string `json: "paintkit"`
							      Type        string `json: "type"`
							      AmmoClip    int 	 `json: "ammo_clip"`
							      AmmoClipMax int 	 `json: "ammo_clip_max"`
							      AmmoReserve int 	 `json: "ammo_reserve"`
							      State       string `json: "state"`
						      }`json: "weapon_5"`
					     Weapon_6 struct {
							      Name     string `json: "name"`
							      Paintkit string `json: "paintkit"`
							      Type     string `json: "type"`
							      State    string `json: "state"`
						      }`json: "weapon_6"`
				     }`json: "weapons"`
			 Match_stats struct {
					     Kills   int        `json: "kills"`
					     Assists int        `json: "assists"`
					     Deaths  int        `json: "deaths"`
					     MVPs    int        `json: "mvps"`
					     Score   int        `json: "score"`
				     }`json: "match_stats"`
		 }`json: "player"`
}

func update(rw http.ResponseWriter, req *http.Request) {
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
	log.Fatal(http.ListenAndServe(":" + port, nil))
}