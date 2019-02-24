package handler
import (
	"fmt"
	"dtypes"
	"log"
)
const offsetX=20
const offsetY=15
/*type Position struct {
	X int;
	Y int;
}*/
type Rect struct {
  XHi int
  YHi int
  XLo int
  YLo int
}
/*var platform[3] Rect
var ladder[3] Rect*/
func OnPlatform(player Rect) int{
  var i int
  log.Println("Executing OnPlatform")
  for i=0;i<len(platform);i++ {
  	//log.Println(player.XHi,platform[i].XHi, "---", player.YLo,"---",platform[i].YHi, "---", player.XLo,platform[i].XLo)
    if player.YLo==platform[i].YHi && player.XLo>platform[i].XHi && player.XHi<platform[i].XLo{
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
	  			log.Println("Executing",ladder[i].YHi)
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
	       		return Rect{player.XHi,ladder[i].YLo-2*offsetY,player.XLo,ladder[i].YLo}
	          }
	      }
  }
   return Rect{}
}
func FallsFromBlock(player Rect) Rect{ // originally was on platform but not now
	var i int;
	for i=0;i<len(platform);i++{
      if(player.XHi>platform[i].XLo || player.XLo<platform[i].XHi){ // foot of player collides with top of block
      	   	return Rect{platform[i].XLo,platform[i].YLo-2*offsetY,platform[i].XLo+2*offsetX,platform[i].YHi}
  	}
  }
   return Rect{}
}
func CollidesWithBlockVertically(player Rect) int{ // falling vertically
  var i int
  for i=0;i<len(platform);i++{
  	 //log.Println(player.YLo,"---",platform[i].YHi)
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
		//log.Println(player.XLo,platform[i].XHi, "---", player.YLo,"---",platform[i].YHi, "---", player.YHi,platform[i].YLo)
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
			return Rect{platform[i].XLo,player.YHi,platform[i].XLo+2*offsetX,player.YLo}
		}
	}
	return Rect{}
}
func GetPositionCollidesWithBlockOnRight(player Rect)Rect{
	var i int
	for i=0;i<len(platform);i++{
		if player.XLo>platform[i].XHi && player.YLo>platform[i].YHi && player.YHi<platform[i].YLo{
			return Rect{platform[i].XHi-2*offsetX,player.YHi,platform[i].XHi,player.YLo}
		}
	}
	return Rect{}
}
func GetPositionCollidesWithBlockVer(player Rect)Rect{
	var i int
  for i=0;i<len(platform);i++{
      if(player.YLo>platform[i].YHi && player.XLo>platform[i].XHi && player.XHi<platform[i].XLo){ // foot of player collides with top of block
      	return Rect{player.XHi,platform[i].YHi-2*offsetY,player.XLo,platform[i].YHi}
  	}
  }
  return Rect{}
}
func GetBoundary(player dtypes.Position) Rect {
  var temp Rect
  temp.XHi = player.X-offsetX
  temp.YHi = player.Y-offsetY
  temp.XLo = player.X+offsetX
  temp.YLo = player.Y+offsetY
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
/*// testing part
func main() {
	p1:=Position{45,35}
	p2:=Position{50,55}
	//platform[0]=Rect{0,39,100,40}
	var p11 Rect =GetBoundary(p1);
	//platform[1]=Rect{70,50,90,60}
	platform[0]=Rect{20,50,160,80}
	var p22 Rect =GetBoundary(p2);
	//fmt.Println(platform[0].XLo)
	//fmt.Println(OnPlatform(p11))
	//fmt.Println(CollidesWithBlockVertically(p11))
	//fmt.Println(CollidesWithBlockOnRightMove(p22))
	var p33 Rect=GetPositionCollidesWithBlockVer(p22)
	//var p33 Rect=GetPositionCollidesWithBlockOnRight(p22)

	/*	platform[1]=Rect{60,60,80,80}
	platform[2]=Rect{0,30,50,50}
	//ladder[0]=Rect{15,25,25,45}
	ladder[0]=Rect{35,12,45,42}
	//fmt.Println(AllignedWithLadder(p11))
	//fmt.Println(AllignedWithLadder(p22))
	log.Println(p33.XHi,p33.YHi,p33.XLo,p33.YLo)
	fmt.Println(p11.XHi,p11.YHi,p11.XLo,p11.YLo);
	/*

	//fmt.Println(on_platform(p11))
	fmt.Println(CollidesWithBlock_vertically(p11))

}*/

