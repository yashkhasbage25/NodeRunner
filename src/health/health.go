package health

import (
	"fmt"
	"sync"
)

//var health int =1000

var mu1 sync.Mutex
var mu2 sync.Mutex
var FirstHealth = 1000
var SecondHealth = 1000

// func decayPlayer1() {
// 	for Player1.health > 0 {
// 		mu1.Lock()
// 		Player1.health = Player1.health - 1
// 		mu1.Unlock()
// 		// time.Sleep(time.Millisecond * 500)
// 	}
// }
// func decayPlayer2() {
// 	for Player2.health > 0 {
// 		mu2.Lock()
// 		Player2.health = Player2.health - 1
// 		// time.Sleep(time.Millisecond * 500)
// 		mu2.Unlock()
// 	}
// }
// func updatePlayer1Increment() {
// 	mu1.Lock()
// 	// time.Sleep(time.Millisecond * 500)
// 	health += 10
// 	mu1.Unlock()
// }
// func updatePlayer1Decrement() {
// 	mu1.Lock()
// 	// time.Sleep(time.Millisecond * 500)
// 	health -= 10
// 	mu1.Unlock()
// }
// func updatePlayer2Increment() {
// 	mu2.Lock()
// 	time.Sleep(time.Millisecond * 500)
// 	Player1.health += 10
// 	mu2.Unlock()
// }
// func updatePlayer2Decrement() {
// 	mu2.Lock()
// 	time.Sleep(time.Millisecond * 500)
// 	Player2.health -= 10
// 	mu2.Unlock()
// }

func UpdateHealth() {

	go decay()
	go update()
	fmt.Scanln()
}
