package handler
import (
	"fmt"
	"dtypes"
	"log"
	"coordinate"
)
const offsetX=20
const offsetY=15
/*type Position struct {
	X int;
	Y int;
}*/

/*var coords.platform[3] dtypes.Rect
var coords.ladder[3] dtypes.Rect*/
func OnPlatform(player dtypes.Rect) int{
  var i int
  log.Println("Executing OnPlatform")
  for i=0;i<len(coords.platform);i++ {
  	//log.Println(player.XHi,coords.platform[i].XHi, "---", player.YLo,"---",coords.platform[i].YHi, "---", player.XLo,coords.platform[i].XLo)
    if player.YLo==coords.platform[i].YHi && player.XLo>coords.platform[i].XHi && player.XHi<coords.platform[i].XLo{
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
  for i=0;i<len(coords.ladder);i++{
    if coords.ladder[i].YLo>=player.YLo && coords.ladder[i].YHi<=player.YLo {
       if center>=coords.ladder[i].XHi && center<=coords.ladder[i].XLo{
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
	  for i=0;i<len(coords.ladder);i++{
	    if coords.ladder[i].YLo>=player.YLo && coords.ladder[i].YHi<=player.YLo {
	       if center>=coords.ladder[i].XHi && center<=coords.ladder[i].XLo{
	  			log.Println("Executing",coords.ladder[i].YHi)
	          	return dtypes.Rect{player.XHi,coords.ladder[i].YHi-2*offset,player.XLo,coords.ladder[i].YHi}
	          }
	      }
  }
  return dtypes.Rect{}
}
func SetAccordingToLadderBottom(player dtypes.Rect) dtypes.Rect{
	var i int
	var center int=(player.XLo+player.XHi)/2
	  for i=0;i<len(coords.ladder);i++{
	    if coords.ladder[i].YLo>=player.YLo && coords.ladder[i].YHi<=player.YLo {
	       if center>=coords.ladder[i].XHi && center<=coords.ladder[i].XLo{
	          	//return dtypes.Rect{coords.ladder[i].XHi-2*offset,coords.ladder[i].YLo-2*offset,coords.ladder[i].XHi,coords.ladder[i].YLo}
	       		return dtypes.Rect{player.XHi,coords.ladder[i].YLo-2*offsetY,player.XLo,coords.ladder[i].YLo}
	          }
	      }
  }
   return dtypes.Rect{}
}
func FallsFromBlock(player dtypes.Rect) dtypes.Rect{ // originally was on coords.platform but not now
	var i int;
	for i=0;i<len(coords.platform);i++{
      if(player.XHi>coords.platform[i].XLo || player.XLo<coords.platform[i].XHi){ // foot of player collides with top of block
      	   	return dtypes.Rect{coords.platform[i].XLo,coords.platform[i].YLo-2*offsetY,coords.platform[i].XLo+2*offsetX,coords.platform[i].YHi}
  	}
  }
   return dtypes.Rect{}
}
func CollidesWithBlockVertically(player dtypes.Rect) int{ // falling vertically
  var i int
  for i=0;i<len(coords.platform);i++{
  	 //log.Println(player.YLo,"---",coords.platform[i].YHi)
      if(player.YLo>coords.platform[i].YHi && player.XLo>coords.platform[i].XHi && player.XHi<coords.platform[i].XLo){ // foot of player collides with top of block
	  log.Println("CollidesWithBlockVertically returns true.")
	  return 1
  	}
  }
  log.Println("CollidesWithBlockVertically returns false.")
  return 0
}
func CollidesWithBlockOnRightMove(player dtypes.Rect) int{
	var i int
	for i=0;i<len(coords.platform);i++{
		//log.Println(player.XLo,coords.platform[i].XHi, "---", player.YLo,"---",coords.platform[i].YHi, "---", player.YHi,coords.platform[i].YLo)
		if player.XLo>coords.platform[i].XHi && player.YLo>coords.platform[i].YHi && player.YHi<coords.platform[i].YLo{
			log.Println("CollidesWithBlockOnRightMove returns true.")
			return 1
		}
	}
	log.Println("CollidesWithBlockOnRightMove returns false.")
	return 0
}
func CollidesWithBlockOnLeftMove(player dtypes.Rect) int{
	var i int
	for i=0;i<len(coords.platform);i++{
		if player.XHi<coords.platform[i].XLo && player.YLo>coords.platform[i].YHi && player.YHi<coords.platform[i].YLo{
			log.Println("CollidesWithBlockOnLeftMove returns true.")
			return 1
		}
	}
	log.Println("CollidesWithBlockOnLeftMove returns false.")
	return 0
}
func GetPositionCollidesWithBlockOnLeft(player dtypes.Rect) dtypes.Rect{
	var i int
	for i=0;i<len(coords.platform);i++{
		if player.XHi<coords.platform[i].XLo && player.YLo>coords.platform[i].YHi && player.YHi<coords.platform[i].YLo{
			return dtypes.Rect{coords.platform[i].XLo,player.YHi,coords.platform[i].XLo+2*offsetX,player.YLo}
		}
	}
	return dtypes.Rect{}
}
func GetPositionCollidesWithBlockOnRight(player dtypes.Rect)dtypes.Rect{
	var i int
	for i=0;i<len(coords.platform);i++{
		if player.XLo>coords.platform[i].XHi && player.YLo>coords.platform[i].YHi && player.YHi<coords.platform[i].YLo{
			return dtypes.Rect{coords.platform[i].XHi-2*offsetX,player.YHi,coords.platform[i].XHi,player.YLo}
		}
	}
	return dtypes.Rect{}
}
func GetPositionCollidesWithBlockVer(player dtypes.Rect)dtypes.Rect{
	var i int
  for i=0;i<len(coords.platform);i++{
      if(player.YLo>coords.platform[i].YHi && player.XLo>coords.platform[i].XHi && player.XHi<coords.platform[i].XLo){ // foot of player collides with top of block
      	return dtypes.Rect{player.XHi,coords.platform[i].YHi-2*offsetY,player.XLo,coords.platform[i].YHi}
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
	for i:=0;i<len(coords.gems);i++ {
		if player.XLo>=coords.gems[i].pos.XHi || player.XHi<=coords.gems[i].pos.XLo{
			UpdateHealth(coords.gems[i].gemtype,coords.gems[i].value,id)
				for j:=0;j<len(coords.freepositions);j++{
					if coords.freepositions[i].available==true  {
						coords.gems[i].XHi=coords.freepositions[i].XHi
						coords.gems[i].XLo=coords.freepositions[i].XLo
						coords.gems[i].YHi=coords.freepositions[i].YHi
						coords.gems[i].XLo=coords.freepositions[i].YLo
						coords.freepositions[i].available=false
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
	//coords.platform[0]=dtypes.Rect{0,39,100,40}
	var p11 dtypes.Rect =GetBoundary(p1);
	//coords.platform[1]=dtypes.Rect{70,50,90,60}
	coords.platform[0]=dtypes.Rect{20,50,160,80}
	var p22 dtypes.Rect =GetBoundary(p2);
	//fmt.Println(coords.platform[0].XLo)
	//fmt.Println(OnPlatform(p11))
	//fmt.Println(CollidesWithBlockVertically(p11))
	//fmt.Println(CollidesWithBlockOnRightMove(p22))
	var p33 dtypes.Rect=GetPositionCollidesWithBlockVer(p22)
	//var p33 dtypes.Rect=GetPositionCollidesWithBlockOnRight(p22)

	/*	coords.platform[1]=dtypes.Rect{60,60,80,80}
	coords.platform[2]=dtypes.Rect{0,30,50,50}
	//coords.ladder[0]=dtypes.Rect{15,25,25,45}
	coords.ladder[0]=dtypes.Rect{35,12,45,42}
	//fmt.Println(AllignedWithLadder(p11))
	//fmt.Println(AllignedWithLadder(p22))
	log.Println(p33.XHi,p33.YHi,p33.XLo,p33.YLo)
	fmt.Println(p11.XHi,p11.YHi,p11.XLo,p11.YLo);
	/*

	//fmt.Println(on_coords.platform(p11))
	fmt.Println(CollidesWithBlock_vertically(p11))

}*/

