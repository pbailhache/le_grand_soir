package game

import "github.com/hajimehoshi/ebiten"

type Player struct {
	playerWidth  int
	playerHeight int
	x            int
	y            int

	sens bool

	realPlayerWidth  int
	realPlayerHeight int
	realPlayerX      int
	realPlayerY      int

	spriteWidth  int
	spriteHeight int
	imageWidth   int
	imageHeight  int

	currentAnimX int
	currentAnimY int
	nbAnimX      int
	nbAnimY      int

	up      bool
	jumping bool
	falling bool

	miniJumping   bool
	miniFalling   bool
	miniJumpSpeed int

	jumpSpeed int
	vX        int
	vY        int
}

var (
	playerImg *ebiten.Image
)

func (p *Player) Update() {
	p.y += p.vY
	p.x += p.vX

	p.vY += Gravity

	if p.x < CameraX {
		p.x = CameraX
	} else if CameraX+ScreenWidth <= p.x+p.playerWidth {
		p.x = CameraX + ScreenWidth - p.playerWidth
	}
	if p.y < CameraY {
		// p.y = CameraY
	} else if p.y+p.playerHeight >= WorldBaseline {
		p.y = WorldBaseline - p.playerHeight
	}
	if p.x == 1180 && p.y == 461 {
		p.x = 0
		p.y = WorldBaseline - 51
	}

	p.realPlayerX = p.x + (p.playerWidth)/3
	p.realPlayerY = p.y
}

func (p *Player) Flip() {
	p.sens = !p.sens
}

func (p *Player) Minijump() {
	if p.miniJumping {
		p.miniFalling = false
		p.miniJumpSpeed += Gravity
		if p.miniJumpSpeed > 0 {
			p.jumping = false
			p.miniFalling = true
		}
		p.vY = p.miniJumpSpeed
	}

	if p.miniFalling && p.miniJumpSpeed < BaseMiniJumpSpeed+1 {
		p.miniJumpSpeed += Gravity
		p.vY = p.miniJumpSpeed
	}

	if p.miniJumpSpeed >= BaseMiniJumpSpeed {
		p.miniFalling = false
		p.miniJumpSpeed = -BaseMiniJumpSpeed
	}
}

func (p *Player) Jump() {
	if p.jumping {
		p.falling = false
		p.jumpSpeed += Gravity
		if p.jumpSpeed > 0 {
			p.jumping = false
			p.falling = true
		}
		p.vY = p.jumpSpeed
	}

	if p.falling && p.jumpSpeed < BaseJumpSpeed+1 {
		p.jumpSpeed += Gravity
		p.vY = p.jumpSpeed
	}

	if p.jumpSpeed >= BaseJumpSpeed {
		p.falling = false
		p.jumpSpeed = -BaseJumpSpeed
	}
}

func (p *Player) ManageCollision(block *Block) {

	nextX := p.realPlayerX + p.vX
	nextY := p.realPlayerY

	if (block.realX >= nextX+p.realPlayerWidth) || // trop à droite
		(block.realX+block.realW <= nextX) || // trop à gauche
		(block.realY >= nextY+p.realPlayerHeight) || // trop en bas
		(block.realY+block.realH <= nextY) { // trop en haut
	} else {
		p.vX = 0
	}

	//fmt.Printf("Box Player (%d,%d,%d,%d) Box block (%d,%d,%d,%d)\n", p.realPlayerX, p.realPlayerY, p.realPlayerX+p.realPlayerWidth, p.realPlayerY+p.realPlayerHeight, block.realX, block.realY, block.realX+block.realW, block.realY+block.realH)

	nextX = p.realPlayerX
	nextY = p.realPlayerY + p.vY

	if (block.realX >= nextX+p.realPlayerWidth) || // trop à droite
		(block.realX+block.realW <= nextX) || // trop à gauche
		(block.realY >= nextY+p.realPlayerHeight) || // trop en bas
		(block.realY+block.realH <= nextY) { // trop en haut
	} else {
		p.vY = 0
	}

}

func (p *Player) Len() int {
	return 1
}

func (p *Player) Dst(i int) (x0, y0, x1, y1 int) {
	//TODO ajouter la hauteur poru que le perso soit "au sol" au début
	return p.x - CameraX, p.y - CameraY, p.x - CameraX + p.playerWidth, p.y - CameraY + p.playerHeight
	//return ScreenWidth/2 - p.playerWidth/2, WorldBaseline + ScreenHeight/2 - p.playerHeight/2, ScreenWidth/2 + p.playerWidth/2, WorldBaseline + ScreenHeight/2 + p.playerHeight/2
}

func (p *Player) Src(i int) (x0, y0, x1, y1 int) {
	//fmt.Printf("Current AnimX = %d, TotalAnimX = %d\n", p.currentAnimX, p.nbAnimX)
	if p.sens {
		return p.currentAnimX*p.spriteWidth + p.spriteWidth,
			p.currentAnimY * p.spriteHeight,
			p.currentAnimX * p.spriteWidth,
			p.currentAnimY*p.spriteHeight + p.spriteHeight
	}

	return p.currentAnimX * p.spriteWidth,
		p.currentAnimY * p.spriteHeight,
		p.currentAnimX*p.spriteWidth + p.spriteWidth,
		p.currentAnimY*p.spriteHeight + p.spriteHeight
}
