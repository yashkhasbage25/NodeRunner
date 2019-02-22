package handler

import (
	"fmt"
	"dtypes"
	"log"
)

const offset=5
// type Position struct {
// 	X int;
// 	Y int;
// }
type Rect struct {
  XHi int
  YHi int
  XLo int
  YLo int
}
var platform[3] Rect
var ladder[3] Rect
func OnPlatform(player Rect) int{
  var i int
  log.Println("Executing OnPlatform")
  for i=0;i<len(platform);i++ {
    if player.YLo==platform[i].YHi && player.XHi>=platform[i].XHi && player.XLo<=platform[i].XLo{
	  log.Println("OnPlatform returns true.")
	  return 1
    }
  }
  log.Println("OnPlatform returns false.")
  return 0
}
func AllignedWithLadder(player Rect) int{
  var i int
  var center int=(player.XLo+player.XHi)/2
  for i=0;i<len(ladder);i++{
    if ladder[i].YLo>=player.YLo && ladder[i].YHi<=player.YLo {
       if center>=ladder[i].XHi && center<=ladder[i].XLo{
		   log.Println("AllignedWithLadder returns true.")
			return 1
          }
      }
  }
  log.Println("AllignedWithLadder returns false.")
  return 0
}
func SetAccordingToLadderTop(player Rect) Rect{ // originally was alligned but not not alligned
	var i int
	var center int = (player.XLo+player.XHi)/2
	  for i=0;i<len(ladder);i++{
	    if ladder[i].YLo>=player.YLo && ladder[i].YHi<=player.YLo {
	       if center>=ladder[i].XHi && center<=ladder[i].XLo{
	          	return Rect{player.XHi,ladder[i].YHi-2*offset,player.XLo,ladder[i].YHi}
	          }
	      }
  }
  return Rect{}
}
func SetAccordingToLadderBottom(player Rect) Rect{
	var i int
	var center int=(player.XLo+player.XHi)/2
	  for i=0;i<len(ladder);i++{
	    if ladder[i].YLo>=player.YLo && ladder[i].YHi<=player.YLo {
	       if center>=ladder[i].XHi && center<=ladder[i].XLo{
	          	//return Rect{ladder[i].XHi-2*offset,ladder[i].YLo-2*offset,ladder[i].XHi,ladder[i].YLo}
	       		return Rect{player.XHi,ladder[i].YLo-2*offset,player.XLo,ladder[i].YLo}
	          }
	      }
  }
   return Rect{}
}
func FallsFromBlock(player Rect) Rect{ // originally was on platform but not now
	var i int;
	for i=0;i<len(platform);i++{
      if(player.XHi>platform[i].XLo || player.XLo<platform[i].XHi){ // foot of player collides with top of block
      	   	return Rect{platform[i].XLo,platform[i].YLo-2*offset,platform[i].XLo+2*offset,platform[i].YHi}
  	}
  }
   return Rect{}
}
func CollidesWithBlockVertically(player Rect) int{ // falling vertically
  var i int
  for i=0;i<len(platform);i++{
      if(player.YLo>platform[i].YHi && player.XLo>platform[i].XHi && player.XHi<platform[i].XLo){ // foot of player collides with top of block
	  log.Println("CollidesWithBlockVertically returns true.")
	  return 1
  	}
  }
  log.Println("CollidesWithBlockVertically returns false.")
  return 0
}
func CollidesWithBlockOnRightMove(player Rect) int{
	var i int
	for i=0;i<len(platform);i++{
		if player.XLo>platform[i].XHi && player.YLo>platform[i].YHi && player.YHi<platform[i].YLo{
			log.Println("CollidesWithBlockOnRightMove returns true.")
			return 1
		}
	}
	log.Println("CollidesWithBlockOnRightMove returns false.")
	return 0
}
func CollidesWithBlockOnLeftMove(player Rect) int{
	var i int
	for i=0;i<len(platform);i++{
		if player.XHi<platform[i].XLo && player.YLo>platform[i].YHi && player.YHi<platform[i].YLo{
			log.Println("CollidesWithBlockOnLeftMove returns true.")
			return 1
		}
	}
	log.Println("CollidesWithBlockOnLeftMove returns false.")
	return 0
}
func GetPositionCollidesWithBlockOnLeft(player Rect) Rect{
	var i int
	for i=0;i<len(platform);i++{
		if player.XHi<platform[i].XLo && player.YLo>platform[i].YHi && player.YHi<platform[i].YLo{
			return Rect{platform[i].XHi,player.YLo,platform[i].XHi+2*offset,player.YHi}
		}
	}
	return Rect{}
}
func GetPositionCollidesWithBlockOnRight(player Rect)Rect{
	var i int
	for i=0;i<len(platform);i++{
		if player.XLo>platform[i].XHi && player.YLo>platform[i].YHi && player.YHi<platform[i].YLo{
			return Rect{platform[i].XLo-2*offset,player.YLo,platform[i].XLo,player.YHi}
		}
	}
	return Rect{}
}
func GetPositionCollidesWithBlockVer(player Rect)Rect{
	var i int
  for i=0;i<len(platform);i++{
      if(player.YLo>platform[i].YHi && player.XLo>platform[i].XHi && player.XHi<platform[i].XLo){ // foot of player collides with top of block
      	return Rect{player.XLo,platform[i].YLo-2*offset,player.XHi,platform[i].YLo}
  	}
  }
  return Rect{}
}
func GetBoundary(player Position) Rect {
  var temp Rect
  temp.XHi = player.X-offset
  temp.YHi = player.Y-offset
  temp.XLo = player.X+offset
  temp.YLo = player.Y+offset
  return temp
}
func GetPosition(player Rect) dtypes.Position{
	var temp dtypes.Position
	temp.X=(player.XHi+player.XLo)/2
	temp.Y=(player.YHi+player.YLo)/2
	return temp
}
func CollidesGem(player Rect) int{
	for i:=0;i<len(gems);i++{
		if player.XLo>=gems[i].XHi || player.XHi<=gems[i].XLo{
			if gems[i].active==true && gems[i].type==1 {
				// increase own health atomically
				UpdatePlayer1Increment()
				for j:=0;j<len(freepositions);j++{
					if freepositions[i]==true  {
						gems[i].XHi=freepositions[i].XHi
						gems[i].XLo=freepositions[i].XLo
						gems[i].YHi=freepositions[i].YHi
						gems[i].XLo=freepositions[i].YLo
						break
					}
				}
			}else if gems[i].active==true && gems[i].type==2 {
				// decrease opponent's health atomically
				UpdatePlayer2Decrement()
				for j:=0;j<len(freepositions);j++{
					if freepositions[i]==true {
						gems[i].XHi=freepositions[i].XHi
						gems[i].XLo=freepositions[i].XLo
						gems[i].YHi=freepositions[i].YHi
						gems[i].XLo=freepositions[i].YLo
						break
					}
				}				
			}
		}
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

