package game

import "github.com/hajimehoshi/ebiten"

type WorldTile struct {
	w int
	h int

	numSprite int

	imageWidth  int
	imageHeight int
	x           int
	y           int
}

var (
	backgroundTileSheetImg *ebiten.Image
)

func (w *WorldTile) Update() {

}

type World struct {
	WorldTiles []*WorldTile
	num        int
}

func (w *World) Update() {
	for _, tile := range w.WorldTiles {
		tile.Update()
	}
}

func (w *World) Len() int {
	return w.num
}

func (w *World) Dst(i int) (x0, y0, x1, y1 int) {
	if i >= w.num {
		return 0, 0, 0, 0
	}
	tile := w.WorldTiles[i]
	return tile.x - CameraX, tile.y - CameraY, tile.x - CameraX + tile.w, tile.y - CameraY + tile.h

}

func (w *World) Src(i int) (x0, y0, x1, y1 int) {
	if w.num <= i {
		return 0, 0, 0, 0
	}
	tile := w.WorldTiles[i]

	if IsCommunist {
		return tile.imageWidth / 2, tile.numSprite * (tile.imageHeight / 5), tile.imageWidth, (tile.numSprite + 1) * tile.imageHeight / 5
	}

	return 0, tile.numSprite * (tile.imageHeight / 5), tile.imageWidth / 2, (tile.numSprite + 1) * tile.imageHeight / 5

	//return tile.numSprite * tile.w, 0, tile.imageWidth, tile.imageHeight

}
