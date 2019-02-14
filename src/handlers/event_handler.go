package handlers

import (
	"dtypes"
	"log"
	"utils"
)

// Handle handles all kinds of events
func Handle(event dtypes.Event) dtypes.Event {

	// var spriteMotionEventTypes = getMotionEventTypes("p1", "p2")
	var spriteMotionEventTypes = []string{"up", "down", "right", "left"}

	if utils.InArray(event.EventType, spriteMotionEventTypes) {
		return handleMotionEvent(event)
	} else {
		log.Fatalf("Invalid event detected '%s'", event.EventType)
		return dtypes.Event{}
	}
}

// handleMotionEvent specifically handles motion events
// maybe it should be coded in a separate file
func handleMotionEvent(event dtypes.Event) dtypes.Event {

	var replyEvent dtypes.Event

	direction :=[4]string{"up","down","left","right"}
	var step uint32=2;
	dx:=[4]uint32{0,0,-step,step}
	dy:=[4]uint32{-step,step,0,0}
	for i:=0;i<4;i++{
		if direction[i]==event.EventType{
			replyEvent = dtypes.Event{
			EventType: "update",
			Object:    event.Object,
			B1Pos:     event.B1Pos,
			B2Pos:     event.B2Pos,
			B3Pos:     event.B3Pos,

			G1Pos: event.G1Pos,
			G2Pos: event.G2Pos,
			G3Pos: event.G3Pos,
			G4Pos: event.G4Pos,
		}
		if event.Object == "p1" {
			replyEvent.P1Pos = dtypes.Position{X: event.P1Pos.X+dx[i], Y: event.P1Pos.Y +dy[i]}
			replyEvent.P2Pos = dtypes.Position{X: event.P2Pos.X, Y: event.P2Pos.Y}
		} else {
			replyEvent.P1Pos = dtypes.Position{X: event.P1Pos.X, Y: event.P1Pos.Y}
			replyEvent.P2Pos = dtypes.Position{X: event.P2Pos.X+dx[i], Y: event.P2Pos.Y +dy[i]}
		}
	}	
 }
	return replyEvent
}
