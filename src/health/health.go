package health
import (
<<<<<<< HEAD
	"fmt"
	"sync"
=======
	"log"
>>>>>>> bc4e7480c194d46d26fc1f97c1f97a9b012a7c2a
)
var firstMutex sync.Mutex
var secondMutex sync.Mutex
var firstHealth int
var secondHealth int 
var pause int
var rate int

<<<<<<< HEAD
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
=======
func SetHealth(string player,int value){
	if player=="p1"{
		firstHealth=value
	}else if player=="p2"{
		secondHealth=value
	}
}
func GetHealth(string player) int{
	if player=="p1"{
		return firstHealth
	}else{
		return secondHealth
	}
}
func SetDecayParams(rate_val, pause_val int){
	rate=rate_val
	pause=pause_val;
}
>>>>>>> bc4e7480c194d46d26fc1f97c1f97a9b012a7c2a

func UpdateHealth(byte operation,int value,string player){

	if operation=='+'{
		log.Println("plus detected")
		firstMutex.Lock()
		if player=="p1"{
			firstHealth+=value;
		}else{
			secondHealth+=value;
		}
		firstMutex.Unlock()
	}else if operation=='-'{
		secondMutex.Lock()
		if player=="p2"{
			firstHealth-=value;
		}else{
			secondHealth-=value;
		}
		secondMutex.Unlock()
	}else if operation=='*'{

		firstMutex.Lock()
		if player=="p1"{
			firstHealth*=value;
		}else{
			secondHealth*=value;
		}
		firstMutex.Unlock()

	}else if operation=='/'{

		secondMutex.Lock()
		if player=="p2"{
			firstHealth/=value;
		}else{
			secondHealth/=value;
		}
		secondMutex.Unlock()
	}

}
func DecayPlayer1(){
	for firstHealth>0 {
		firstMutex.Lock()
		firstHealth-=rate;
		time.Sleep(time.Millisecond*pause)
		firstMutex.Unlock()		
	}
}
func DecayPlayer2(){
	for secondHealth>0 {
		secondMutex.Lock()
		secondHealth-=rate;
		time.Sleep(time.Millisecond*pause)
		secondMutex.Unlock()		
	}
}