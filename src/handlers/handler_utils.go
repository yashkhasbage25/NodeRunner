package handler

import (
	"log"

	"github.com/IITH-SBJoshi/concurrency-3/src/coords"
	"github.com/IITH-SBJoshi/concurrency-3/src/dtypes"
	"github.com/IITH-SBJoshi/concurrency-3/src/health"
)

// offsetx and offsety are declared to get rectangle boundaray from center positions
const offsetX = 15
const offsetY = 20

// OnPlatform takes player's bounding rectangle and checks whether it is on a platform
func OnPlatform(player dtypes.Rect) bool {
	var i int
	log.Println("Executing OnPlatform")
	for i = 0; i < len(coords.Platform); i++ {
		//log.Println(player.YLo, "---", coords.Platform[i].YHi, "---", player.XLo, coords.Platform[i].XLo)
		if player.YLo == coords.Platform[i].YHi && player.XLo > coords.Platform[i].XHi && player.XHi < coords.Platform[i].XLo {
			log.Println("OnPlatform returns true.")
			return true
		}
	}
	log.Println("OnPlatform returns false.")
	return false
}

// AllignedWithLadder returns true if x coordinate of center point lies within x coordinates of ladder
// and y coordinate of bottom point of player lies between y coordinates of ladder
func AllignedWithLadder(player dtypes.Rect) bool {
	var i int
	var center int = (player.XLo + player.XHi) / 2
	for i = 0; i < len(coords.Ladder); i++ {
		if coords.Ladder[i].YLo >= player.YLo && coords.Ladder[i].YHi <= player.YLo {
			if center >= coords.Ladder[i].XHi && center <= coords.Ladder[i].XLo {
				log.Println("AllignedWithLadder returns true.")
				return true
			}
		}
	}
	log.Println("AllignedWithLadder returns false.")
	return false
}

// SetAccordingToLadderTop As it excedded ladder height lower y coordinate of player is set to upper y coordinate of platform
// and according top y coordinate of ladder is adjusted using offset
func SetAccordingToLadderTop(player dtypes.Rect) dtypes.Rect { // originally was alligned but not not alligned
	var i int
	var center int = (player.XLo + player.XHi) / 2
	for i = 0; i < len(coords.Ladder); i++ {
		if coords.Ladder[i].YLo >= player.YLo && coords.Ladder[i].YHi >= player.YLo {
			if center >= coords.Ladder[i].XHi && center <= coords.Ladder[i].XLo {
				log.Println("Executing", coords.Ladder[i].YHi)
				return dtypes.Rect{player.XHi, coords.Ladder[i].YHi - 2*offsetY, player.XLo, coords.Ladder[i].YHi}
			}
		}
	}
	return dtypes.Rect{}
}

// SetAccordingToLadderBottom As it excedded ladder height while moving lower y coordinate of player is set to lower y coordinate of platform
// and according top y coordinate of ladder is adjusted using offset
func SetAccordingToLadderBottom(player dtypes.Rect) dtypes.Rect {
	var i int
	var center int = (player.XLo + player.XHi) / 2
	for i = 0; i < len(coords.Ladder); i++ {
		if coords.Ladder[i].YLo <= player.YLo && coords.Ladder[i].YHi <= player.YLo {
			if center >= coords.Ladder[i].XHi && center <= coords.Ladder[i].XLo {
				//return dtypes.Rect{coords.Ladder[i].XHi-2*offset,coords.Ladder[i].YLo-2*offset,coords.Ladder[i].XHi,coords.Ladder[i].YLo}
				return dtypes.Rect{player.XHi, coords.Ladder[i].YLo - 2*offsetY, player.XLo, coords.Ladder[i].YLo}
			}
		}
	}
	return dtypes.Rect{}
}

// FallsFromBlock Checks if player fell form block
func FallsFromBlock(player dtypes.Rect) bool { // originally was on coords.Platform but not now
	var i int
	for i = 0; i < len(coords.Platform); i++ {
		log.Println("platform number", i)
		if (player.YLo == coords.Platform[i].YHi) && (player.XHi > coords.Platform[i].XLo || player.XLo < coords.Platform[i].XHi) { // foot of player collides with top of block
			log.Println("fallsfromblock returns true", player.YLo, coords.Platform[i].YHi, player.XHi, coords.Platform[i].XLo, player.XLo, coords.Platform[i].XHi)
			return true
		}
	}
	// log.Println("fallsfromblock returns false", player.XHi, coords.Platform[i].XLo, player.XLo, coords.Platform[i].XHi)
	return false
}

// CollidesWithBlockVertically Checks if player collided vertically
func CollidesWithBlockVertically(player dtypes.Rect) bool { // falling vertically
	var i int
	for i = 0; i < len(coords.Platform); i++ {
		log.Println(i, player.YLo, "---", coords.Platform[i].YHi, player.XLo, coords.Platform[i].XHi, player.XHi, coords.Platform[i].XLo)
		if player.YLo < coords.Platform[i].YHi && player.XLo > coords.Platform[i].XHi && player.XHi < coords.Platform[i].XLo { // foot of player collides with top of block
			log.Println("CollidesWithBlockVertically returns true.")
			return true
		}
	}
	log.Println("CollidesWithBlockVertically returns false.")
	return false
}

// CollidesWithBlockOnRightMove Checks if player collided with a block while moving right
func CollidesWithBlockOnRightMove(player dtypes.Rect) bool {
	var i int
	for i = 0; i < len(coords.Platform); i++ {
		//log.Println(player.XLo,coords.Platform[i].XHi, "---", player.YLo,"---",coords.Platform[i].YHi, "---", player.YHi,coords.Platform[i].YLo)
		if player.XLo > coords.Platform[i].XHi && player.YLo > coords.Platform[i].YHi && player.YHi < coords.Platform[i].YLo {
			log.Println("CollidesWithBlockOnRightMove returns true.")
			return true
		}
	}
	log.Println("CollidesWithBlockOnRightMove returns false.")
	return false
}

// CollidesWithBlockOnLeftMove Checks if player collided with a block while moving left
func CollidesWithBlockOnLeftMove(player dtypes.Rect) bool {
	var i int
	for i = 0; i < len(coords.Platform); i++ {
		if player.XHi < coords.Platform[i].XLo && player.YLo > coords.Platform[i].YHi && player.YHi < coords.Platform[i].YLo {
			log.Println("CollidesWithBlockOnLeftMove returns true.")
			return true
		}
	}
	log.Println("CollidesWithBlockOnLeftMove returns false.")
	return false
}

// GetPositionCollidesWithBlockOnLeft get updated Position of player when it collided with block on left move
func GetPositionCollidesWithBlockOnLeft(player dtypes.Rect) dtypes.Rect {
	var i int
	for i = 0; i < len(coords.Platform); i++ {
		if player.XHi < coords.Platform[i].XLo && player.YLo > coords.Platform[i].YHi && player.YHi < coords.Platform[i].YLo {
			return dtypes.Rect{coords.Platform[i].XLo, player.YHi, coords.Platform[i].XLo + 2*offsetX, player.YLo}
		}
	}
	return dtypes.Rect{}
}

// GetPositionCollidesWithBlockOnRight get updated Position of player when it collided with block on right move
func GetPositionCollidesWithBlockOnRight(player dtypes.Rect) dtypes.Rect {
	var i int
	for i = 0; i < len(coords.Platform); i++ {
		if player.XLo > coords.Platform[i].XHi && player.YLo > coords.Platform[i].YHi && player.YHi < coords.Platform[i].YLo {
			return dtypes.Rect{coords.Platform[i].XHi - 2*offsetX, player.YHi, coords.Platform[i].XHi, player.YLo}
		}
	}
	return dtypes.Rect{}
}

// GetPositionCollidesWithBlockVer get updated Position of player when it collided with block on vertical move
func GetPositionCollidesWithBlockVer(player dtypes.Rect) dtypes.Rect {
	var i int
	for i = 0; i < len(coords.Platform); i++ {
		if player.YLo < coords.Platform[i].YHi && player.XLo > coords.Platform[i].XHi && player.XHi < coords.Platform[i].XLo { // foot of player collides with top of block
			return dtypes.Rect{player.XHi, coords.Platform[i].YHi - 2*offsetY, player.XLo, coords.Platform[i].YHi}
		}
	}
	return dtypes.Rect{}
}

// GetBoundary get bounding rectangle using coordinates of center point
func GetBoundary(player dtypes.Position) dtypes.Rect {
	var temp dtypes.Rect
	temp.XHi = player.X - offsetX
	temp.YHi = player.Y - offsetY
	temp.XLo = player.X + offsetX
	temp.YLo = player.Y + offsetY
	return temp
}

// GetPosition get center point using bounding rectangle
func GetPosition(player dtypes.Rect) dtypes.Position {
	var temp dtypes.Position
	temp.X = (player.XHi + player.XLo) / 2
	temp.Y = (player.YHi + player.YLo) / 2
	return temp
}

// CollidesGem Check whether collided with gem
func CollidesGem(player dtypes.Rect, id string) {
	for i := 0; i < len(coords.Gems); i++ {
		if player.XLo >= coords.Gems[i].Pos.XHi && player.XHi <= coords.Gems[i].Pos.XLo && player.YLo == coords.Gems[i].Pos.YLo {
			health.UpdateHealth(coords.Gems[i].Gemtype, coords.Gems[i].Value, id)
			log.Println("gem collected")
			// collided with gem
			// now change position of gems
			for j := 0; j < len(coords.Freepositions); j++ {
				if coords.Freepositions[j].Available == true {
					coords.Gems[i].Pos.XHi = coords.Freepositions[j].Pos.XHi
					coords.Gems[i].Pos.XLo = coords.Freepositions[j].Pos.XLo
					coords.Gems[i].Pos.YHi = coords.Freepositions[j].Pos.YHi
					coords.Gems[i].Pos.YLo = coords.Freepositions[j].Pos.YLo
					coords.Freepositions[j].Available = false
					break
				}
			}
		}
	}
}

// Check if player collides with bots
func CollidesWithBot(player dtypes.Rect, b11 dtypes.Rect, b22 dtypes.Rect, b33 dtypes.Rect) bool {
	if (player.YLo >= b11.YHi && player.YHi <= b11.YLo && player.XHi <= b11.XLo && player.XLo >= b11.XHi) ||
		(player.YLo >= b22.YHi && player.YHi <= b22.YLo && player.XHi <= b22.XLo && player.XLo >= b22.XHi) ||
		(player.YLo >= b33.YHi && player.YHi <= b33.YLo && player.XHi <= b33.XLo && player.XLo >= b33.XHi) {
		log.Println("Collision with bot returns true")
		return true
	}
	log.Println("Collision with bot returns false")
	return false
}

/*// testing part
func main() {
	p1:=Position{45,35}
	p2:=Position{50,55}
	//coords.Platform[0]=dtypes.Rect{0,39,100,40}
	var p11 dtypes.Rect =GetBoundary(p1);
	//coords.Platform[1]=dtypes.Rect{70,50,90,60}
	coords.Platform[0]=dtypes.Rect{20,50,160,80}
	var p22 dtypes.Rect =GetBoundary(p2);
	//fmt.Println(coords.Platform[0].XLo)
	//fmt.Println(OnPlatform(p11))
	//fmt.Println(CollidesWithBlockVertically(p11))
	//fmt.Println(CollidesWithBlockOnRightMove(p22))
	var p33 dtypes.Rect=GetPositionCollidesWithBlockVer(p22)
	//var p33 dtypes.Rect=GetPositionCollidesWithBlockOnRight(p22)

	/*	coords.Platform[1]=dtypes.Rect{60,60,80,80}
	coords.Platform[2]=dtypes.Rect{0,30,50,50}
	//coords.Ladder[0]=dtypes.Rect{15,25,25,45}
	coords.Ladder[0]=dtypes.Rect{35,12,45,42}
	//fmt.Println(AllignedWithLadder(p11))
	//fmt.Println(AllignedWithLadder(p22))
	log.Println(p33.XHi,p33.YHi,p33.XLo,p33.YLo)
	fmt.Println(p11.XHi,p11.YHi,p11.XLo,p11.YLo);
	/*

	//fmt.Println(on_coords.Platform(p11))
	fmt.Println(CollidesWithBlock_vertically(p11))

}*/
