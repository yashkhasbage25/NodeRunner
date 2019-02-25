package platform

import (
	"github.com/IITH-SBJoshi/concurrency-3/src/dtypes"
)

var Platform [8]dtypes.Rect
var Ladder [6]dtypes.Rect

func Initialize() {
	Platform[0] = dtypes.Rect{0, 45, 840, 75}
	Platform[1] = dtypes.Rect{800, 125, 1200, 155}
	Platform[2] = dtypes.Rect{0, 200, 400, 230}
	Platform[3] = dtypes.Rect{600, 210, 1200, 240}
	Platform[4] = dtypes.Rect{0, 325, 800, 355}
	Platform[5] = dtypes.Rect{170, 450, 530.480}
	Platform[6] = dtypes.Rect{800, 440, 400, 470}
	Platform[7] = dtypes.Rect{0, 550, 1200, 50}

	Ladder[0] = dtypes.Rect{170, 450, 200, 550}
	Ladder[1] = dtypes.Rect{1170, 440, 1200, 550}
	Ladder[2] = dtypes.Rect{800, 210, 830, 440}
	Ladder[3] = dtypes.Rect{370, 45, 400, 200}
	Ladder[4] = dtypes.Rect{100, 200, 130, 325}
	Ladder[5] = dtypes.Rect{400, 325, 430, 550}
}
