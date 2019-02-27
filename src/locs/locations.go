package locs

import (
	"log"

	"github.com/IITH-SBJoshi/concurrency-3/src/dtypes"
)

var playerOne dtypes.Position
var playerTwo dtypes.Position

var botOne dtypes.Position
var botTwo dtypes.Position
var botThree dtypes.Position

var gemOne dtypes.Position
var gemTwo dtypes.Position
var gemThree dtypes.Position
var gemFour dtypes.Position

var player [2]dtypes.Position
var bot [3]dtypes.Position
var gem [4]dtypes.Position

func InitializeLocations() {
	player[0] = dtypes.Position{
		X: 35,
		Y: 530,
	}
	player[1] = dtypes.Position{
		X: 415,
		Y: 530,
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
		X: 515,
		Y: 305,
	}
	gem[2] = dtypes.Position{
		X: 465,
		Y: 305,
	}
	gem[3] = dtypes.Position{
		X: 615,
		Y: 420,
	}
}

func GetPlayerPos(index int) dtypes.Position {
	if index < 0 || index >= len(player) {
		log.Fatalln("GetPlayerPos index out of range", index)
	}
	return player[index]
}

func GetBotPos(index int) dtypes.Position {
	if index < 0 || index >= len(bot) {
		log.Fatalln("GetBotPos index out of range", index)
	}
	return bot[index]
}

func GetGemPos(index int) dtypes.Position {
	if index < 0 || index >= len(gem) {
		log.Fatalln("GetGemPos index out of range", index)
	}
	return gem[index]
}

func SetPlayerPos(index int, pos dtypes.Position) {
	if index < 0 || index >= len(player) {
		log.Fatalln("SetPlayerPos index out of range", index)
	}
	player[index] = pos
}

func SetBotPos(index int, pos dtypes.Position) {
	if index < 0 || index >= len(bot) {
		log.Fatalln("SetBotPos index out of range", index)
	}
	bot[index] = pos
}

func SetGemPos(index int, pos dtypes.Position) {
	if index < 0 || index >= len(gem) {
		log.Fatalln("SetGemPos index out of range", index)
	}
	gem[index] = pos
}

func GetCurrentLocations(event dtypes.Event) dtypes.Event {
	log.Println("Getting current locations")
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
