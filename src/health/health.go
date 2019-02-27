package health

import (
	"log"
	"sync"
	"time"
)

// firstMutex provides safety while updating health of first 
// player
var firstMutex sync.Mutex

// secondMutex provides safety while updating health of second 
// player
var secondMutex sync.Mutex

// firstHealth is the health of first player
var firstHealth int

// secondHealth is the health of second player
var secondHealth int

// pause is the time for which health updater pauses while 
// updating health
var pause int

// rate is the rate by which health decreases in every update
var rate int

// gameWinChannel communicates the winner info to server
var gameWinChannel chan int

// SetGameWinChannel is a setter for gameWinChannel in this
// package
func SetGameWinChannel(winChannel chan int) {
	gameWinChannel = winChannel
}

// SetHealth is the setter for initial health of a player
func SetHealth(player string, value int) {
	if player == "p1" {
		firstHealth = value
	} else if player == "p2" {
		secondHealth = value
	}
}

// GetHealth is a getter for health of a player
func GetHealth(player string) int {
	if player == "p1" {
		return firstHealth
	} else {
		return secondHealth
	}
}

// SetDecayParams is a setter for rate and pause as explained 
// above
func SetDecayParams(rate_val, pause_val int) {
	rate = rate_val
	pause = pause_val
}

// UpdateHealth updated health on collision with gem. Health can be 
// updated in various ways depending on the gem
func UpdateHealth(operation byte, value int, player string) {
	if operation == '+' {
		log.Println("plus detected")
		firstMutex.Lock()
		if player == "p1" {
			firstHealth += value
		} else {
			secondHealth += value
		}
		firstMutex.Unlock()
	} else if operation == '-' {
		secondMutex.Lock()
		if player == "p2" {
			firstHealth -= value
			if firstHealth <= 0 {
				gameWinChannel <- 1
			}
		} else {
			secondHealth -= value
			if secondHealth <= 0 {
				gameWinChannel <- 0
			}
		}
		secondMutex.Unlock()
	} else if operation == '*' {

		firstMutex.Lock()
		if player == "p1" {
			firstHealth *= value
		} else {
			secondHealth *= value
		}
		firstMutex.Unlock()

	} else if operation == '/' {

		secondMutex.Lock()
		if player == "p2" {
			firstHealth /= value
			if firstHealth <= 0 {
				gameWinChannel <- 1
			}
		} else {
			secondHealth /= value
			if secondHealth <= 0 {
				gameWinChannel <- 0
			}
		}
		secondMutex.Unlock()
	}

}

// DecayPlayer1 decays health of first player
func DecayPlayer1() {
	for firstHealth > 0 {
		firstMutex.Lock()
		firstHealth -= rate
		time.Sleep(time.Duration(pause) * time.Millisecond)
		firstMutex.Unlock()
	}
	gameWinChannel <- 1
}

// DecayPlayer2 decays health of second player
func DecayPlayer2() {
	for secondHealth > 0 {
		secondMutex.Lock()
		secondHealth -= rate
		time.Sleep(time.Millisecond * time.Duration(pause))
		secondMutex.Unlock()
	}
	gameWinChannel <- 0
}
