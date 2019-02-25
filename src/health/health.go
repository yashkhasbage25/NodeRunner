package health

import (
	"log"
	"sync"
	"time"
)

var firstMutex sync.Mutex
var secondMutex sync.Mutex
var firstHealth int
var secondHealth int
var pause int
var rate int
var gameWinChannel chan int

func SetGameWinChannel(winChannel chan int) {
	gameWinChannel = winChannel
}

func SetHealth(player string, value int) {
	if player == "p1" {
		firstHealth = value
	} else if player == "p2" {
		secondHealth = value
	}
}
func GetHealth(player string) int {
	if player == "p1" {
		return firstHealth
	} else {
		return secondHealth
	}
}
func SetDecayParams(rate_val, pause_val int) {
	rate = rate_val
	pause = pause_val
}

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
func DecayPlayer1() {
	for firstHealth > 0 {
		firstMutex.Lock()
		firstHealth -= rate
		time.Sleep(time.Duration(pause) * time.Millisecond)
		firstMutex.Unlock()
	}
	gameWinChannel <- 1
}
func DecayPlayer2() {
	for secondHealth > 0 {
		secondMutex.Lock()
		secondHealth -= rate
		time.Sleep(time.Millisecond * time.Duration(pause))
		secondMutex.Unlock()
	}
	gameWinChannel <- 0
}
