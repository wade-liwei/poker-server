package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"encoding/json"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/public/login", PublicLogin)
	r.HandleFunc("/lobby/rooms", LobbyRooms)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	http.ListenAndServe("0.0.0.0:8083", c.Handler(r))

}

type Room struct {
	IdRoom           int64  `json:"id_room"`
	ServerIp         string `json:"server_ip"`
	Name             string `json:"name"`
	Gproto           string `json:"gproto"`
	Desc             string `json:"description"`
	MaxPlayers       int64  `json:"max_players"`
	MinCoinForAccess int64  `json:"minCoinForAccess"`
	IsOfficial       bool   `json:"isOfficial"`
	Players          int64  `json:"players"`
}
type Rooms struct {
	Rooms []Room `json:"rooms"`
}

func LobbyRooms(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("22222222222222\n")
	w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte("{\"id_room\":333333,\"server_ip\":\"4.4.4.4\",\"name\":\"rommName\",\"gproto\":\"22222\",\"description\":\"desc111111\",\"max_players\":10},\"minCoinForAccess\":100,\"\"))
	// //w.WriteString(`{}`)
	room := Room{
		IdRoom:           1111,
		ServerIp:         "1.1.1.1",
		Name:             "tableName",
		Gproto:           "g proto",
		Desc:             "room desc",
		MaxPlayers:       10,
		MinCoinForAccess: 10000,
		IsOfficial:       true,
		Players:          10,
	}

	rooms := Rooms{Rooms: []Room{room}}

	jsonRooms, _ := json.Marshal(rooms)

	w.Write([]byte(jsonRooms))
}

func PublicLogin(w http.ResponseWriter, r *http.Request) {
	token, _ := CreateToken("1")
	fmt.Printf("111111111111111\n")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"jwtToken\":\"" + token + "\",\"sessionID\":\"2222222\",\"userID\":\"1\",\"operationSuccess\":true}"))
}

func CreateToken(userId string) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}
