package locs

import (
	"log"

	"github.com/IITH-SBJoshi/concurrency-3/src/dtypes"
)

// player array stores positions of the players
var player [2]dtypes.Position

// bot rray stores positions of bots
var bot [3]dtypes.Position

// gem array stores the positions of gems
var gem [4]dtypes.Position

// InitializeLocations initializes  positions of players, bots and gems
func InitializeLocations() {
	player[0] = dtypes.Position{
		X: 900,
		Y: 420,
	}
	player[1] = dtypes.Position{
		X: 1000,
		Y: 420,
	}
	bot[0] = dtypes.Position{
		X: 355,
		Y: 180,
	}
	bot[1] = dtypes.Position{
		X: 235,
		Y: 25,
	}
	bot[2] = dtypes.Position{
		X: 435,
		Y: 25,
	}
	gem[0] = dtypes.Position{
		X: 215,
		Y: 305,
	}
	gem[1] = dtypes.Position{
		X: 535,
		Y: 305,
	}
	gem[2] = dtypes.Position{
		X: 465,
		Y: 305,
	}
	gem[3] = dtypes.Position{
		X: 615,
		Y: 530,
	}
}

// GetPlayerPos is a getter for position of a player
func GetPlayerPos(index int) dtypes.Position {
	if index < 0 || index >= len(player) {
		log.Fatalln("GetPlayerPos index out of range", index)
	}
	return player[index]
}

// GetBotPos is a getter fot position of a bot
func GetBotPos(index int) dtypes.Position {
	if index < 0 || index >= len(bot) {
		log.Fatalln("GetBotPos index out of range", index)
	}
	return bot[index]
}

// GetGemPos is a getter for position of a gem
func GetGemPos(index int) dtypes.Position {
	if index < 0 || index >= len(gem) {
		log.Fatalln("GetGemPos index out of range", index)
	}
	return gem[index]
}

// SetPlayerPos is a setter for position of a player
func SetPlayerPos(index int, pos dtypes.Position) {
	if index < 0 || index >= len(player) {
		log.Fatalln("SetPlayerPos index out of range", index)
	}
	player[index] = pos
}

// SetBotPos is a setter for position of a bot
func SetBotPos(index int, pos dtypes.Position) {
	if index < 0 || index >= len(bot) {
		log.Fatalln("SetBotPos index out of range", index)
	}
	bot[index] = pos
}

// SetGemPos is a setter for position of a gem
func SetGemPos(index int, pos dtypes.Position) {
	if index < 0 || index >= len(gem) {
		log.Fatalln("SetGemPos index out of range", index)
	}
	gem[index] = pos
}

// GetCurrentLocations is a getter fot current location of all players,
// bots and gems
func GetCurrentLocations(event dtypes.Event) dtypes.Event {
	log.Println("Getting current locations, with prev locations as:", event.GetStr())
	return dtypes.Event{
		EventType: event.EventType,
		Object:    event.Object,
		P1Pos:     GetPlayerPos(0),
		P2Pos:     GetPlayerPos(1),
		B1Pos:     GetBotPos(0),
		B2Pos:     GetBotPos(1),
		B3Pos:     GetBotPos(2),
		G1Pos:     GetGemPos(0),
		G2Pos:     GetGemPos(1),
		G3Pos:     GetGemPos(2),
		G4Pos:     GetGemPos(3),
	}
}

// SetCurrentLocations is a setter for current locations of all player, bots and gems
func SetCurrentLocations(event dtypes.Event) {
	log.Println("Setting current locations")
	player[0] = event.P1Pos
	player[1] = event.P2Pos

	bot[0] = event.B1Pos
	bot[1] = event.B2Pos
	bot[2] = event.B3Pos

	gem[0] = event.G1Pos
	gem[1] = event.G2Pos
	gem[2] = event.G3Pos
	gem[3] = event.G4Pos
}
