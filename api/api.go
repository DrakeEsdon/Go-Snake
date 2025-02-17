package api

import (
	"encoding/json"
	"fmt"
	"github.com/DrakeEsdon/Go-Snake/datatypes"
	"github.com/DrakeEsdon/Go-Snake/snake"
	"log"
	"net/http"
)

type RequestLogger map[string][]datatypes.GameRequest

var requestLogger RequestLogger
var latestGame string

func getRequestLogger() RequestLogger {
	if requestLogger == nil {
		requestLogger = make(RequestLogger)
	}
	return requestLogger
}

func logRequest(request datatypes.GameRequest) {
	logger := getRequestLogger()
	latestGame = request.Game.ID
	logger[request.Game.ID] = append(logger[request.Game.ID], request)
}

func HandleLatestLog(w http.ResponseWriter, r *http.Request) {
	response := getRequestLogger()[latestGame]
	if response == nil {
		response = []datatypes.GameRequest{}
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}

func GetServerInfo() datatypes.BattlesnakeInfoResponse {
	return datatypes.BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "DrakeEsdon & HugoKlepsch",
		Color:      "#dcc010",
		Head:       "fang",
		Tail:       "pixel",
	}
}

// HandleIndex is called when your Battlesnake is created and refreshed
// by play.battlesnake.com. BattlesnakeInfoResponse contains information about
// your Battlesnake, including what it should look like on the game board.
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	response := GetServerInfo()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}

// HandleStart is called at the start of each game your Battlesnake is playing.
// The GameRequest object contains information about the game that's about to start.
// TODO: Use this function to decide how your Battlesnake is going to look on the board.
func HandleStart(w http.ResponseWriter, r *http.Request) {
	request := datatypes.GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	// Nothing to respond with here
	fmt.Print("START\n")
}

// HandleMove is called for each turn of each game.
// Valid responses are "up", "down", "left", or "right".
// TODO: Use the information in the GameRequest object to determine your next move.
func HandleMove(w http.ResponseWriter, r *http.Request) {
	request := datatypes.GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	move, shout := snake.ChooseMove(request)

	response := datatypes.MoveResponse{
		Move: move,
		Shout: shout,
	}

	fmt.Printf("MOVE: %s\n", response.Move)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}

// HandleEnd is called when a game your Battlesnake was playing has ended.
// It's purely for informational purposes, no response required.
func HandleEnd(w http.ResponseWriter, r *http.Request) {
	request := datatypes.GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	// Nothing to respond with here
	fmt.Print("END\n")
}
