package game

import "github.com/hajimehoshi/ebiten"

type Block struct {
	blockWidth  int
	blockHeight int

	realX int
	realY int
	realW int
	realH int

	isSpecial bool

	imageWidth  int
	imageHeight int
	x           int
	y           int
}

var (
	// 128 x 128
	blockImg      *ebiten.Image
	twoByTwoImg   *ebiten.Image
	twoByThreeImg *ebiten.Image
	twoByFourImg  *ebiten.Image
	drapeauxImg   *ebiten.Image
)

func (b *Block) Update() {

}

type Blocks struct {
	blocks []*Block
	num    int
}

func (b *Blocks) Update() {
	for _, block := range b.blocks {
		block.Update()
	}
}

func (b *Blocks) Len() int {
	return b.num
}

func (b *Blocks) Dst(i int) (x0, y0, x1, y1 int) {
	if i >= b.num {
		return 0, 0, 0, 0
	}
	bb := b.blocks[i]
	return bb.x - CameraX, bb.y - CameraY, bb.x - CameraX + bb.blockWidth, bb.y - CameraY + bb.blockHeight

}

func (b *Blocks) Src(i int) (x0, y0, x1, y1 int) {
	if b.num <= i {
		return 0, 0, 0, 0
	}
	bb := b.blocks[i]

	if IsCommunist {
		if bb.isSpecial {
			return bb.imageWidth / 2, bb.imageHeight / 2, bb.imageWidth, bb.imageHeight
		}
		return bb.imageWidth / 2, 0, bb.imageWidth, bb.imageHeight / 2
	} else {
		if bb.isSpecial {
			return 0, bb.imageHeight / 2, bb.imageWidth / 2, bb.imageHeight
		}
		return 0, 0, bb.imageWidth / 2, bb.imageHeight / 2
	}
}

func CreateBlock(x, y, imgW, imgH int) *Block {
	return &Block{
		imageWidth:  imgW,
		imageHeight: imgH,
		blockWidth:  imgW / 2,
		blockHeight: imgH / 2,
		x:           x,
		y:           y,

		realX: x,
		realY: y + 10,
		realW: imgW/2 - 10,
		realH: imgH/2 - 20,
	}
}

func CreateBlockSpecial(x, y, imgW, imgH int) *Block {
	return &Block{
		imageWidth:  imgW,
		imageHeight: imgH,
		blockWidth:  imgW / 2,
		blockHeight: imgH / 2,
		x:           x,
		y:           y,

		isSpecial: true,

		realX: x,
		realY: y + 10,
		realW: imgW/2 - 10,
		realH: imgH/2 - 20,
	}
}
