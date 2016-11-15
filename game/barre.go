package game

import "github.com/hajimehoshi/ebiten"

type Barre struct {
	x int
	y int

	w int
	h int
}

var (
	bgBarreImg    *ebiten.Image
	innerBarreImg *ebiten.Image
)

func (b *Barre) Update() {
	if CommunismLevel > 1000*2 {
		IsCommunist = !IsCommunist
		CommunismLevel = 0
	}
}

func (b *Barre) Len() int {
	return 1
}

func (b *Barre) Dst(i int) (x0, y0, x1, y1 int) {
	return b.x, b.y, b.x + (b.w * (CommunismLevel / 10 / 2) / 100), b.y + b.h
}

func (b *Barre) Src(i int) (x0, y0, x1, y1 int) {
	return 0, 0, b.w, b.h
}
