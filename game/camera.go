package game

import "github.com/hajimehoshi/ebiten"

var (
	CameraX int
	CameraY int
)

type Camera struct {
	x int
	y int
}

var (
	cameraImg *ebiten.Image
)

func (c *Camera) Update() {
	CameraX = c.x
	CameraY = c.y
	//fmt.Printf("X = %d, Y = %d, X+W = %d, Y+H = %d, worldW = %d, worldH = %d\n", c.x, c.y, c.x+ScreenWidth, c.y+ScreenHeight, WorldWidth, WorldHeight)
	if c.x < 0 {
		c.x = 0
	} else if WorldWidth <= c.x+ScreenWidth {
		c.x = WorldWidth - ScreenWidth
	}
	if c.y < 0 {
		c.y = 0
	} else if WorldHeight <= c.y+ScreenHeight {
		c.y = 2*(WorldHeight-ScreenHeight) - c.y
	}
}

func (c *Camera) Len() int {
	return 1
}

func (c *Camera) Dst(i int) (x0, y0, x1, y1 int) {
	return 0, 0, ScreenWidth, ScreenHeight
}

func (c *Camera) Src(i int) (x0, y0, x1, y1 int) {
	return c.x, c.y, c.x + ScreenWidth, c.y + ScreenHeight
}
