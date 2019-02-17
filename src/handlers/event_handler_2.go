package events
import (
	"handlerfunctions"
	"dtypes"
)

func handleMotionEvent(event dtypes.Event) dtypes.Event {

	var replyEvent dtypes.Event

	direction :=[4]string{"up","down","left","right"}
	var freefallP1 bool=false
	var freefallP2 bool=false
	var step int=2;
	dx:=[4]int{0,0,-step,step}
	dy:=[4]int{-step,step,0,0}
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
			var p11 Rect =GetBoundary(event.P1Pos)
			var p22 Rect =GetBoundary(replyEvent.P1Pos)
			var updated Rect
			if i==0 {
				if(!AllignedWithLadder(p11)){
					// no change
					log.Println("up but not alligned with ladder")
					updated Rect=p11
				}else if(AllignedWithLadder(p11) && AllignedWithLadder(p22)){
					// success
					log.Println("up and not alligned with ladder")
					updated Rect=p22
				}else{
					// get ladder top
					log.Println("up and not alligned with ladder restricted")
					updated Rect=SetAccordingToLadderTop(p11)
				}
			}else if i==1 {
				if(!AllignedWithLadder(p11)){
					// no change
					log.Println("down but not alligned with ladder")
					updated Rect=p11
				}else if(AllignedWithLadder(p11) && AllignedWithLadder(p22)){
					// success
					log.Println("down and alligned with ladder")
					updated Rect=p22
				}else{
					// get ladder bottom
						log.Println("down and alligned with ladder restricted")
					updated Rect=SetAccordingToLadderBottom(p11)
				}
			}else if i==2{
				if AllignedWithLadder(p11) && AllignedWithLadder(p22){
					updated Rect=p22
					log.Println("was alligned with ladder on pressing left")
				}else if AllignedWithLadder(p11) && !AllignedWithLadder(p22){
					log.Println("freefall")
					freefallP1=true;
				}else if !OnPlatform(p11){
					log.Println("not on platform")
					updated Rect=p11
				}else if(CollidesWithBlockOnLeftMove(p22)){
					log.Println("collided with block on left")
					updated Rect=GetPositionCollidesWithBlockOnLeft(p22)
				}else if(FallsFromBlock(p22)){
						log.Println("fell from block and freefall")
						freefallP1=true;
				}else{
						log.Println("successfull left move")
					updated Rect=p22
				}
			}else if i==3{
				if AllignedWithLadder(p11) && AllignedWithLadder(p22){
					updated Rect=p22
					log.Println("was alligned with ladder on pressing right")
				}else if AllignedWithLadder(p11) && !AllignedWithLadder(p22){
					freefallP1=true
					log.Println("freefall")
				}else if !OnPlatform(p11){
					log.Println("not on platform")
					updated Rect=p11
				}else if(CollidesWithBlockOnRightMove(p22)){
					log.Println("collided with block on right")
					updated Rect=GetPositionCollidesWithBlockOnRight(p22)
				}else if(FallsFromBlock(p22)){
					log.Println("fell from block and freefall")
					freefallP1=true
				}else{
					log.Println("successfull right move")
					updated Rect=p22
				}
			}
			var  dtypes.Position temporary =getposition(updated)

			if (freefallP1){
				log.Println("freefall")
				var  dtypes.Position temporary2 ={temporary.X,temporary.Y+2*step}
				var p11 Rect =GetBoundary(temporary)
				var p22 Rect =GetBoundary(temporary2)
				if CollidesWithBlockVertically(p22){
					log.Println("collidedwith block while freefalling")
					freefallP1=false;
					updated Rect=GetPositionCollidesWithBlockVer(p22)
				}else{
					updated Rect=p22
				}
			}
			replyEvent.P1Pos=getposition(updated)
		}
		/*else {
			replyEvent.P2Pos = dtypes.Position{X: event.P2Pos.X, Y: event.P2Pos.Y}
			replyEvent.P1Pos = dtypes.Position{X: event.P1Pos.X, Y: event.P1Pos.Y}
			replyEvent.P2Pos = dtypes.Position{X: event.P2Pos.X+dx[i], Y: event.P2Pos.Y +dy[i]}
		}*/
	}
 }
	return replyEvent
}
