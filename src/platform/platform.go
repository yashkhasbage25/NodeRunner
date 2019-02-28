package platform

import (
	"github.com/IITH-SBJoshi/concurrency-3/src/dtypes"
)

// Platform stores rectangles describing platforms
var Platform [8]dtypes.Rect

// Ladder stores rectangles describing ladders
var Ladder [6]dtypes.Rect

// Initialize initializes positions of platforms and ladders
func Initialize() {
	Platform[0] = dtypes.Rect{
		XHi: 0,
		YHi: 45,
		XLo: 840,
		YLo: 75,
	}
	Platform[1] = dtypes.Rect{
		XHi: 800,
		YHi: 125,
		XLo: 1200,
		YLo: 155,
	}
	Platform[2] = dtypes.Rect{
		XHi: 0,
		YHi: 200,
		XLo: 400,
		YLo: 230,
	}
	Platform[3] = dtypes.Rect{
		XHi: 600,
		YHi: 210,
		XLo: 1200,
		YLo: 240,
	}
	Platform[4] = dtypes.Rect{
		XHi: 0,
		YHi: 325,
		XLo: 800,
		YLo: 355,
	}
	Platform[5] = dtypes.Rect{
		XHi: 170,
		YHi: 450,
		XLo: 530,
		YLo: 480,
	}
	Platform[6] = dtypes.Rect{ // FIXME:
		XHi: 800,
		YHi: 440,
		XLo: 1200,
		YLo: 470,
	}
	Platform[7] = dtypes.Rect{ // FIXME:
		XHi: 0,
		YHi: 550,
		XLo: 1200,
		YLo: 50,
	}

	Ladder[0] = dtypes.Rect{
		XHi: 170,
		YHi: 450,
		XLo: 200,
		YLo: 550,
	}
	Ladder[1] = dtypes.Rect{
		XHi: 1170,
		YHi: 440,
		XLo: 1200,
		YLo: 550,
	}
	Ladder[2] = dtypes.Rect{
		XHi: 800,
		YHi: 210,
		XLo: 830,
		YLo: 440,
	}
	Ladder[3] = dtypes.Rect{
		XHi: 370,
		YHi: 45,
		XLo: 400,
		YLo: 200,
	}
	Ladder[4] = dtypes.Rect{
		XHi: 100,
		YHi: 200,
		XLo: 130,
		YLo: 325,
	}
	Ladder[5] = dtypes.Rect{
		XHi: 400,
		YHi: 325,
		XLo: 430,
		YLo: 450,
	}
}
