package handler

import (
	"log"
	"math/rand"

	"github.com/IITH-SBJoshi/concurrency-3/src/coords"
	"github.com/IITH-SBJoshi/concurrency-3/src/dtypes"
	"github.com/IITH-SBJoshi/concurrency-3/src/health"
)

var gameWinChannel chan int

func SetGameWinChannel(winChannel chan int) {
	gameWinChannel = winChannel
	health.SetGameWinChannel(winChannel)
}

func Handle(event dtypes.Event) dtypes.Event {
	log.Println("event obtained in handler.Handle is ", event.GetStr())
	var replyEvent dtypes.Event
	if event.EventType == "Update" {
		replyEvent = dtypes.Event{
			EventType: "Update",
			Object:    event.Object,
			B1Pos:     event.B1Pos,
			B2Pos:     event.B2Pos,
			B3Pos:     event.B3Pos,
			P1Pos:     event.P1Pos,
			P2Pos:     event.P2Pos,
			G1Pos:     event.G1Pos,
			G2Pos:     event.G2Pos,
			G3Pos:     event.G3Pos,
			G4Pos:     event.G4Pos,
			P1Health:  health.GetHealth("p1"),
			P2Health:  health.GetHealth("p2"),
		}
		log.Println("handler replies with ordinary update eventtype", replyEvent.GetStr())
		return replyEvent
	}
	if event.EventType == "SocketClosedUnexpectedly" {
		log.Fatal("Unexpectedly closed:", event.Object)
	}
	if event.EventType == "Teleport" {
		if event.Object == "p1" {
			j := rand.Intn(10)
			replyEvent = dtypes.Event{
				EventType: "Update",
				Object:    event.Object,
				B1Pos:     event.B1Pos,
				B2Pos:     event.B2Pos,
				B3Pos:     event.B3Pos,
				P1Pos:     coords.Randompos[j],
				P2Pos:     event.P2Pos,
				G1Pos:     event.G1Pos,
				G2Pos:     event.G2Pos,
				G3Pos:     event.G3Pos,
				G4Pos:     event.G4Pos,
				P1Health:  health.GetHealth("p1"),
				P2Health:  health.GetHealth("p2"),
			}
		} else if event.Object == "p2" {
			j := rand.Intn(10)
			replyEvent = dtypes.Event{
				EventType: "Update",
				Object:    event.Object,
				B1Pos:     event.B1Pos,
				B2Pos:     event.B2Pos,
				B3Pos:     event.B3Pos,
				P1Pos:     event.P1Pos,
				P2Pos:     coords.Randompos[j],
				G1Pos:     event.G1Pos,
				G2Pos:     event.G2Pos,
				G3Pos:     event.G3Pos,
				G4Pos:     event.G4Pos,
				P1Health:  health.GetHealth("p1"),
				P2Health:  health.GetHealth("p2"),
			}
		}
		log.Println("handler replies for teleport", replyEvent.GetStr())
		return replyEvent
	}
	direction := [4]string{"Up", "Down", "Left", "Right"}
	var freeFallP1 bool = false
	var freeFallP2 bool = false
	var step int = 2
	dx := [5]int{0, 0, -step, step}
	dy := [5]int{-step, step, 0, 0, 0}
	for i := 0; i < 4; i++ {
		if direction[i] == event.EventType {
			var b11 = GetBoundary(event.B1Pos)
			var b22 = GetBoundary(event.B2Pos)
			var b33 = GetBoundary(event.B3Pos)
			log.Println("Direction detected:", direction[i])
			replyEvent = dtypes.Event{
				EventType: "Update",
				Object:    event.Object,
				B1Pos:     event.B1Pos,
				B2Pos:     event.B2Pos,
				B3Pos:     event.B3Pos,
			}
			log.Println("Set default attr for replyEvent")
			if event.Object == "p1" {
				log.Println("Object of this event is p1")
				replyEvent.P1Pos = dtypes.Position{
					X: event.P1Pos.X + dx[i],
					Y: event.P1Pos.Y + dy[i],
				}
				var p11 = GetBoundary(event.P1Pos)
				var p22 = GetBoundary(replyEvent.P1Pos)
				if(p22.XHi<0){
					log.Println("out of bounds p1")
					p22.XHi=0
					p22.XLo=30
				}
				if(p22.XLo>1200){
					log.Println("out of bounds p1")
					p22.XLo=1200
					p22.XHi=1170
				}
				var updatedRect dtypes.Rect
				if i == 0 {
					if !AllignedWithLadder(p11) {
						// no change
						log.Println("up but not alligned with ladder")
						updatedRect = p11
					} else if AllignedWithLadder(p11) && AllignedWithLadder(p22) {
						// success
						log.Println("up and not alligned with ladder")
						updatedRect = p22
					} else {
						// get ladder top
						log.Println("up and not alligned with ladder restricted")
						updatedRect = SetAccordingToLadderTop(p11)
					}
				} else if i == 1 {
					if !AllignedWithLadder(p11) {
						// no change
						log.Println("down but not alligned with ladder")
						updatedRect = p11
					} else if AllignedWithLadder(p11) && AllignedWithLadder(p22) {
						// success
						log.Println("down and alligned with ladder")
						updatedRect = p22
					} else {
						// get ladder bottom
						log.Println("down and alligned with ladder restricted")
						updatedRect = SetAccordingToLadderBottom(p11)
					}
				} else if i == 2 {
					if AllignedWithLadder(p11) && AllignedWithLadder(p22) {
						updatedRect = p22
						log.Println("was alligned with ladder on pressing left")
					} else if !OnPlatform(p22) && AllignedWithLadder(p11) && !AllignedWithLadder(p22) {
						freeFallP1 = true
						log.Println("freefall")
						updatedRect = p22
					} else if FallsFromBlock(p22) {
						log.Println("fell from block and freefall")
						freeFallP1 = true
						updatedRect = p22
					} else if !OnPlatform(p11) {
						log.Println("not on Platform")
						updatedRect = p11
					} else if CollidesWithBlockOnLeftMove(p22) {
						log.Println("collided with block on left")
						updatedRect = GetPositionCollidesWithBlockOnLeft(p22)
					} else {
						log.Println("successfull left move")
						updatedRect = p22
					}
					CollidesGem(updatedRect, "p1")
				} else if i == 3 {
					if AllignedWithLadder(p11) && AllignedWithLadder(p22) {
						updatedRect = p22
						log.Println("was alligned with ladder on pressing right")
					} else if !OnPlatform(p22) && AllignedWithLadder(p11) && !AllignedWithLadder(p22) {
						freeFallP1 = true
						log.Println("freefall")
						updatedRect = p22
					} else if FallsFromBlock(p22) {
						log.Println("fell from block and freefall")
						freeFallP1 = true
						updatedRect = p22
					} else if !OnPlatform(p11) {
						log.Println("not on Platform")
						updatedRect = p11
					} else if CollidesWithBlockOnRightMove(p22) {
						log.Println("collided with block on right")
						updatedRect = GetPositionCollidesWithBlockOnRight(p22)
					} else {
						log.Println("successfull right move")
						updatedRect = p22
					}
					CollidesGem(updatedRect, "p1")
				}
				var temporary dtypes.Position = GetPosition(updatedRect)

				if freeFallP1 {
					log.Println("freefall")
					log.Println(temporary.X, temporary.Y, "---------")
					temporary2 := dtypes.Position{temporary.X, temporary.Y + 2*step}
					p11 = GetBoundary(temporary)
					p22 = GetBoundary(temporary2)
					if CollidesWithBlockVertically(p22) {
						log.Println("collidedwith block while freefalling")
						freeFallP1 = false
						updatedRect = GetPositionCollidesWithBlockVer(p22)
					} else {
						updatedRect = p22
					}
				}

				if CollidesWithBot(updatedRect, b11, b22, b33) {
					replyEvent.EventType = "P1Looses"
				}
				replyEvent.P1Pos = GetPosition(updatedRect)
				replyEvent.P2Pos = event.P2Pos
				replyEvent.P1Health = health.GetHealth("p1")
				replyEvent.P2Health = health.GetHealth("p2")
				replyEvent.G1Pos = GetPosition(coords.Gems[0].Pos)
				replyEvent.G2Pos = GetPosition(coords.Gems[1].Pos)
				replyEvent.G3Pos = GetPosition(coords.Gems[2].Pos)
				replyEvent.G4Pos = GetPosition(coords.Gems[3].Pos)
			}

			if event.Object == "p2" {
				log.Println("Object of this event is p2")
				replyEvent.P2Pos = dtypes.Position{
					X: event.P2Pos.X + dx[i],
					Y: event.P2Pos.Y + dy[i],
				}
				p11 := GetBoundary(event.P2Pos)
				p22 := GetBoundary(replyEvent.P2Pos)
				if(p22.XHi<0){
					log.Println("out of bounds p2")
					p22.XHi=0
					p22.XLo=30
				}
				if(p22.XLo>1200){
					log.Println("out of bounds p2")
					p22.XLo=1200
					p22.XHi=1170
				}
				var updatedRect dtypes.Rect
				if i == 0 {
					if !AllignedWithLadder(p11) {
						// no change
						log.Println("up but not alligned with ladder")
						updatedRect = p11
					} else if AllignedWithLadder(p11) && AllignedWithLadder(p22) {
						// success
						log.Println("up and not alligned with ladder")
						updatedRect = p22
					} else {
						// get ladder top
						log.Println("up and not alligned with ladder restricted")
						updatedRect = SetAccordingToLadderTop(p11)
					}
				} else if i == 1 {
					if !AllignedWithLadder(p11) {
						// no change
						log.Println("down but not alligned with ladder")
						updatedRect = p11
					} else if AllignedWithLadder(p11) && AllignedWithLadder(p22) {
						// success
						log.Println("down and alligned with ladder")
						updatedRect = p22
					} else {
						// get ladder bottom
						log.Println("down and alligned with ladder restricted")
						updatedRect = SetAccordingToLadderBottom(p11)
					}
				} else if i == 2 {
					if AllignedWithLadder(p11) && AllignedWithLadder(p22) {
						updatedRect = p22
						log.Println("was alligned with ladder on pressing left")
					} else if !OnPlatform(p22) && AllignedWithLadder(p11) && !AllignedWithLadder(p22) {
						freeFallP2 = true
						log.Println("freefall")
						updatedRect = p22
					} else if FallsFromBlock(p22) {
						log.Println("fell from block and freefall")
						freeFallP2 = true
						updatedRect = p22
					} else if !OnPlatform(p11) {
						log.Println("not on Platform")
						updatedRect = p11
					} else if CollidesWithBlockOnLeftMove(p22) {
						log.Println("collided with block on left")
						updatedRect = GetPositionCollidesWithBlockOnLeft(p22)
					} else {
						log.Println("successfull left move")
						updatedRect = p22
					}
					CollidesGem(updatedRect, "p2")
				} else if i == 3 {
					if AllignedWithLadder(p11) && AllignedWithLadder(p22) {
						updatedRect = p22
						log.Println("was alligned with ladder on pressing right")
					} else if !OnPlatform(p22) && AllignedWithLadder(p11) && !AllignedWithLadder(p22) {
						freeFallP2 = true
						log.Println("freefall")
						updatedRect = p22
					} else if FallsFromBlock(p22) {
						log.Println("fell from block and freefall")
						freeFallP2 = true
						updatedRect = p22
					} else if !OnPlatform(p11) {
						log.Println("not on Platform")
						updatedRect = p11
					} else if CollidesWithBlockOnRightMove(p22) {
						log.Println("collided with block on right")
						updatedRect = GetPositionCollidesWithBlockOnRight(p22)
					} else {
						log.Println("successfull right move")
						updatedRect = p22
					}
					CollidesGem(updatedRect, "p2")
				}
				var temporary dtypes.Position = GetPosition(updatedRect)

				if freeFallP2 {
					log.Println("freefall")
					log.Println(temporary.X, temporary.Y, "---------")
					temporary2 := dtypes.Position{temporary.X, temporary.Y + 2*step}
					p11 = GetBoundary(temporary)
					p22 = GetBoundary(temporary2)
					if CollidesWithBlockVertically(p22) {
						log.Println("collidedwith block while freefalling")
						freeFallP2 = false
						updatedRect = GetPositionCollidesWithBlockVer(p22)
					} else {
						updatedRect = p22
					}
				}
				if CollidesWithBot(updatedRect, b11, b22, b33) {
					replyEvent.EventType = "P2Looses"
				}
				replyEvent.P2Pos = GetPosition(updatedRect)
				replyEvent.P1Pos = event.P1Pos
				replyEvent.P1Health = health.GetHealth("p1")
				replyEvent.P2Health = health.GetHealth("p2")
				replyEvent.G1Pos = GetPosition(coords.Gems[0].Pos)
				replyEvent.G2Pos = GetPosition(coords.Gems[1].Pos)
				replyEvent.G3Pos = GetPosition(coords.Gems[2].Pos)
				replyEvent.G4Pos = GetPosition(coords.Gems[3].Pos)
			}

		}
	}
	log.Println("handler replied with event ", event.GetStr())
	return replyEvent
}
