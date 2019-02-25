package handler
import (
	"fmt"
	"github.com/IITH-SBJoshi/concurrency-3/src/dtypes"
	"log"
	"github.com/IITH-SBJoshi/concurrency-3/src/coords"
)
const offsetX=20
const offsetY=15
/*type Position struct {
	X int;
	Y int;
}*/

/*var coords.Platform[3] dtypes.Rect
var coords.Ladder[3] dtypes.Rect*/
func OnPlatform(player dtypes.Rect) int{
  var i int
  log.Println("Executing OnPlatform")
  for i=0;i<len(coords.Platform);i++ {
  	//log.Println(player.XHi,coords.Platform[i].XHi, "---", player.YLo,"---",coords.Platform[i].YHi, "---", player.XLo,coords.Platform[i].XLo)
    if player.YLo==coords.Platform[i].YHi && player.XLo>coords.Platform[i].XHi && player.XHi<coords.Platform[i].XLo{
	  log.Println("OnPlatform returns true.")
	  return 1
    }
  }
  log.Println("OnPlatform returns false.")
  return 0
}
func AllignedWithLadder(player dtypes.Rect) int{
  var i int
  var center int=(player.XLo+player.XHi)/2
  for i=0;i<len(coords.Ladder);i++{
    if coords.Ladder[i].YLo>=player.YLo && coords.Ladder[i].YHi<=player.YLo {
       if center>=coords.Ladder[i].XHi && center<=coords.Ladder[i].XLo{
		   log.Println("AllignedWithLadder returns true.")
			return 1
          }
      }
  }
  log.Println("AllignedWithLadder returns false.")
  return 0
}
func SetAccordingToLadderTop(player dtypes.Rect) dtypes.Rect{ // originally was alligned but not not alligned
	var i int
	var center int = (player.XLo+player.XHi)/2
	  for i=0;i<len(coords.Ladder);i++{
	    if coords.Ladder[i].YLo>=player.YLo && coords.Ladder[i].YHi<=player.YLo {
	       if center>=coords.Ladder[i].XHi && center<=coords.Ladder[i].XLo{
	  			log.Println("Executing",coords.Ladder[i].YHi)
	          	return dtypes.Rect{player.XHi,coords.Ladder[i].YHi-2*offset,player.XLo,coords.Ladder[i].YHi}
	          }
	      }
  }
  return dtypes.Rect{}
}
func SetAccordingToLadderBottom(player dtypes.Rect) dtypes.Rect{
	var i int
	var center int=(player.XLo+player.XHi)/2
	  for i=0;i<len(coords.Ladder);i++{
	    if coords.Ladder[i].YLo>=player.YLo && coords.Ladder[i].YHi<=player.YLo {
	       if center>=coords.Ladder[i].XHi && center<=coords.Ladder[i].XLo{
	          	//return dtypes.Rect{coords.Ladder[i].XHi-2*offset,coords.Ladder[i].YLo-2*offset,coords.Ladder[i].XHi,coords.Ladder[i].YLo}
	       		return dtypes.Rect{player.XHi,coords.Ladder[i].YLo-2*offsetY,player.XLo,coords.Ladder[i].YLo}
	          }
	      }
  }
   return dtypes.Rect{}
}
func FallsFromBlock(player dtypes.Rect) dtypes.Rect{ // originally was on coords.Platform but not now
	var i int;
	for i=0;i<len(coords.Platform);i++{
      if(player.XHi>coords.Platform[i].XLo || player.XLo<coords.Platform[i].XHi){ // foot of player collides with top of block
      	   	return dtypes.Rect{coords.Platform[i].XLo,coords.Platform[i].YLo-2*offsetY,coords.Platform[i].XLo+2*offsetX,coords.Platform[i].YHi}
  	}
  }
   return dtypes.Rect{}
}
func CollidesWithBlockVertically(player dtypes.Rect) int{ // falling vertically
  var i int
  for i=0;i<len(coords.Platform);i++{
  	 //log.Println(player.YLo,"---",coords.Platform[i].YHi)
      if(player.YLo>coords.Platform[i].YHi && player.XLo>coords.Platform[i].XHi && player.XHi<coords.Platform[i].XLo){ // foot of player collides with top of block
	  log.Println("CollidesWithBlockVertically returns true.")
	  return 1
  	}
  }
  log.Println("CollidesWithBlockVertically returns false.")
  return 0
}
func CollidesWithBlockOnRightMove(player dtypes.Rect) int{
	var i int
	for i=0;i<len(coords.Platform);i++{
		//log.Println(player.XLo,coords.Platform[i].XHi, "---", player.YLo,"---",coords.Platform[i].YHi, "---", player.YHi,coords.Platform[i].YLo)
		if player.XLo>coords.Platform[i].XHi && player.YLo>coords.Platform[i].YHi && player.YHi<coords.Platform[i].YLo{
			log.Println("CollidesWithBlockOnRightMove returns true.")
			return 1
		}
	}
	log.Println("CollidesWithBlockOnRightMove returns false.")
	return 0
}
func CollidesWithBlockOnLeftMove(player dtypes.Rect) int{
	var i int
	for i=0;i<len(coords.Platform);i++{
		if player.XHi<coords.Platform[i].XLo && player.YLo>coords.Platform[i].YHi && player.YHi<coords.Platform[i].YLo{
			log.Println("CollidesWithBlockOnLeftMove returns true.")
			return 1
		}
	}
	log.Println("CollidesWithBlockOnLeftMove returns false.")
	return 0
}
func GetPositionCollidesWithBlockOnLeft(player dtypes.Rect) dtypes.Rect{
	var i int
	for i=0;i<len(coords.Platform);i++{
		if player.XHi<coords.Platform[i].XLo && player.YLo>coords.Platform[i].YHi && player.YHi<coords.Platform[i].YLo{
			return dtypes.Rect{coords.Platform[i].XLo,player.YHi,coords.Platform[i].XLo+2*offsetX,player.YLo}
		}
	}
	return dtypes.Rect{}
}
func GetPositionCollidesWithBlockOnRight(player dtypes.Rect)dtypes.Rect{
	var i int
	for i=0;i<len(coords.Platform);i++{
		if player.XLo>coords.Platform[i].XHi && player.YLo>coords.Platform[i].YHi && player.YHi<coords.Platform[i].YLo{
			return dtypes.Rect{coords.Platform[i].XHi-2*offsetX,player.YHi,coords.Platform[i].XHi,player.YLo}
		}
	}
	return dtypes.Rect{}
}
func GetPositionCollidesWithBlockVer(player dtypes.Rect)dtypes.Rect{
	var i int
  for i=0;i<len(coords.Platform);i++{
      if(player.YLo>coords.Platform[i].YHi && player.XLo>coords.Platform[i].XHi && player.XHi<coords.Platform[i].XLo){ // foot of player collides with top of block
      	return dtypes.Rect{player.XHi,coords.Platform[i].YHi-2*offsetY,player.XLo,coords.Platform[i].YHi}
  	}
  }
  return dtypes.Rect{}
}
func GetBoundary(player dtypes.Position) dtypes.Rect {
  var temp dtypes.Rect
  temp.XHi = player.X-offsetX
  temp.YHi = player.Y-offsetY
  temp.XLo = player.X+offsetX
  temp.YLo = player.Y+offsetY
  return temp
}
func GetPosition(player dtypes.Rect) dtypes.Position{
	var temp dtypes.Position
	temp.X=(player.XHi+player.XLo)/2
	temp.Y=(player.YHi+player.YLo)/2
	return temp
}
func CollidesGem(player dtypes.Rect,id string) {
	for i:=0;i<len(coords.Gems);i++ {
		if player.XLo>=coords.Gems[i].Pos.XHi || player.XHi<=coords.Gems[i].Pos.XLo{
			UpdateHealth(coords.Gems[i].Gemtype,coords.Gems[i].Value,id)
				for j:=0;j<len(coords.Freepositions);j++{
					if coords.Freepositions[i].Available==true  {
						coords.Gems[i].XHi=coords.Freepositions[i].XHi
						coords.Gems[i].XLo=coords.Freepositions[i].XLo
						coords.Gems[i].YHi=coords.Freepositions[i].YHi
						coords.Gems[i].XLo=coords.Freepositions[i].YLo
						coords.Freepositions[i].Available=false
						break
					}
				}
			}
		}
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
