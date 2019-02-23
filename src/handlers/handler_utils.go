package handler

import (
	"dtypes"
	"log"
)

const offset = 5

// type Position struct {
// 	X int;
// 	Y int;
// }

// Rect represents a bounding box rectangle
type Rect struct {
	XHi int
	YHi int
	XLo int
	YLo int
}

var platform [3]Rect
var ladder [3]Rect

func OnPlatform(player Rect) bool {
	var i int
	log.Println("Executing OnPlatform")
	for i = 0; i < len(platform); i++ {
		if player.YLo == platform[i].YHi && player.XHi >= platform[i].XHi && player.XLo <= platform[i].XLo {
			log.Println("OnPlatform returns true.")
			return true
		}
	}
	log.Println("OnPlatform returns false.")
	return false
}
func AllignedWithLadder(player Rect) bool {
	var i int
	var center int = (player.XLo + player.XHi) / 2
	for i = 0; i < len(ladder); i++ {
		if ladder[i].YLo >= player.YLo && ladder[i].YHi <= player.YLo {
			if center >= ladder[i].XHi && center <= ladder[i].XLo {
				log.Println("AllignedWithLadder returns true.")
				return true
			}
		}
	}
	log.Println("AllignedWithLadder returns false.")
	return false
}
func SetAccordingToLadderTop(player Rect) Rect { // originally was alligned but not not alligned
	var i int
	var center int = (player.XLo + player.XHi) / 2
	for i = 0; i < len(ladder); i++ {
		if ladder[i].YLo >= player.YLo && ladder[i].YHi <= player.YLo {
			if center >= ladder[i].XHi && center <= ladder[i].XLo {
				return Rect{player.XHi, ladder[i].YHi - 2*offset, player.XLo, ladder[i].YHi}
			}
		}
	}
	return Rect{}
}
func SetAccordingToLadderBottom(player Rect) Rect {
	var i int
	var center int = (player.XLo + player.XHi) / 2
	for i = 0; i < len(ladder); i++ {
		if ladder[i].YLo >= player.YLo && ladder[i].YHi <= player.YLo {
			if center >= ladder[i].XHi && center <= ladder[i].XLo {
				//return Rect{ladder[i].XHi-2*offset,ladder[i].YLo-2*offset,ladder[i].XHi,ladder[i].YLo}
				return Rect{player.XHi, ladder[i].YLo - 2*offset, player.XLo, ladder[i].YLo}
			}
		}
	}
	return Rect{}
}
func FallsFromBlock(player Rect) *Rect { // originally was on platform but not now
	var i int
	for i = 0; i < len(platform); i++ {
		if player.XHi > platform[i].XLo || player.XLo < platform[i].XHi { // foot of player collides with top of block
			return &Rect{platform[i].XLo, platform[i].YLo - 2*offset, platform[i].XLo + 2*offset, platform[i].YHi}
		}
	}
	return nil
}
func CollidesWithBlockVertically(player Rect) bool { // falling vertically
	var i int
	for i = 0; i < len(platform); i++ {
		if player.YLo > platform[i].YHi && player.XLo > platform[i].XHi && player.XHi < platform[i].XLo { // foot of player collides with top of block
			log.Println("CollidesWithBlockVertically returns true.")
			return true
		}
	}
	log.Println("CollidesWithBlockVertically returns false.")
	return false
}
func CollidesWithBlockOnRightMove(player Rect) bool {
	var i int
	for i = 0; i < len(platform); i++ {
		if player.XLo > platform[i].XHi && player.YLo > platform[i].YHi && player.YHi < platform[i].YLo {
			log.Println("CollidesWithBlockOnRightMove returns true.")
			return true
		}
	}
	log.Println("CollidesWithBlockOnRightMove returns false.")
	return false
}
func CollidesWithBlockOnLeftMove(player Rect) bool {
	var i int
	for i = 0; i < len(platform); i++ {
		if player.XHi < platform[i].XLo && player.YLo > platform[i].YHi && player.YHi < platform[i].YLo {
			log.Println("CollidesWithBlockOnLeftMove returns true.")
			return true
		}
	}
	log.Println("CollidesWithBlockOnLeftMove returns false.")
	return false
}
func GetPositionCollidesWithBlockOnLeft(player Rect) Rect {
	var i int
	for i = 0; i < len(platform); i++ {
		if player.XHi < platform[i].XLo && player.YLo > platform[i].YHi && player.YHi < platform[i].YLo {
			return Rect{platform[i].XHi, player.YLo, platform[i].XHi + 2*offset, player.YHi}
		}
	}
	return Rect{}
}
func GetPositionCollidesWithBlockOnRight(player Rect) Rect {
	var i int
	for i = 0; i < len(platform); i++ {
		if player.XLo > platform[i].XHi && player.YLo > platform[i].YHi && player.YHi < platform[i].YLo {
			return Rect{platform[i].XLo - 2*offset, player.YLo, platform[i].XLo, player.YHi}
		}
	}
	return Rect{}
}
func GetPositionCollidesWithBlockVer(player Rect) Rect {
	var i int
	for i = 0; i < len(platform); i++ {
		if player.YLo > platform[i].YHi && player.XLo > platform[i].XHi && player.XHi < platform[i].XLo { // foot of player collides with top of block
			return Rect{player.XLo, platform[i].YLo - 2*offset, player.XHi, platform[i].YLo}
		}
	}
	return Rect{}
}
func GetBoundary(player dtypes.Position) Rect {

	return Rect{
		XHi: player.X - offset,
		YHi: player.Y - offset,
		XLo: player.X + offset,
		YLo: player.Y + offset,
	}
}

func GetPosition(player Rect) dtypes.Position {
	return dtypes.Position{
		X: (player.XHi + player.XLo) / 2,
		Y: (player.YHi + player.YLo) / 2,
	}
}

// testing part
/*func main() {
	p1:=Position{54,15}
	platform[0]=Rect{0,14,50,50}
	platform[0]=Rect{24,56,898,66}
	fmt.Println(platform[0].XLo)
	/*platform[1]=Rect{60,60,80,80}
	platform[2]=Rect{0,30,50,50}
	/*ladder[0]=Rect{15,25,25,45}
	ladder[1]=Rect{30,30,50,60}
	ladder[0]=Rect{10,25,20,35}
	var p11 Rect =GetBoundary(p1);
	fmt.Println(p11.XHi,p11.YHi,p11.XLo,p11.YLo);
	//fmt.Println(on_platform(p11))
	//fmt.Println(alligned_with_ladder(p11))
	fmt.Println(CollidesWithBlock_vertically(p11))

}*/
