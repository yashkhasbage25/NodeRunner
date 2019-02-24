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

func SetHealth(string player, int value) {
	if player == "p1" {
		firstHealth = value
	} else if player == "p2" {
		secondHealth = value
	}
}
func GetHealth(string player) int {
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

func UpdateHealth(byte operation, int value, string player) {

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
		} else {
			secondHealth -= value
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
		} else {
			secondHealth /= value
		}
		secondMutex.Unlock()
	}

}
func DecayPlayer1() {
	for firstHealth > 0 {
		firstMutex.Lock()
		firstHealth -= rate
		time.Sleep(time.Millisecond * pause)
		firstMutex.Unlock()
	}
}
func DecayPlayer2() {
	for secondHealth > 0 {
		secondMutex.Lock()
		secondHealth -= rate
		time.Sleep(time.Millisecond * pause)
		secondMutex.Unlock()
	}
}
