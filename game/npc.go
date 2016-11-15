package game

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

type Npc struct {
	npcWidth  int
	npcHeight int
	x         int
	y         int

	sens bool

	realNpcWidth  int
	realNpcHeight int
	realNpcX      int
	realNpcY      int

	spriteWidth  int
	spriteHeight int
	imageWidth   int
	imageHeight  int

	currentAnimX int
	currentAnimY int
	nbAnimX      int
	nbAnimY      int

	up        bool
	jumping   bool
	falling   bool
	jumpSpeed int
	vX        int
	vY        int

	maxMovX     int
	currentMovX int
}

var (
	porcCapitalisteImg *ebiten.Image
	camaradeImg        *ebiten.Image
)

func (n *Npc) Update() {
	n.currentAnimX = (n.currentAnimX + 1) % n.nbAnimX
	n.y += n.vY
	n.x += n.vX

	n.currentMovX += int(math.Abs(float64(n.vX)))
	if n.currentMovX >= n.maxMovX {
		n.Flip()

		n.currentMovX = 0
	}

	n.vY += Gravity

	if n.y < CameraY {
		n.y = CameraY
	} else if n.y+n.npcHeight >= WorldBaseline {
		n.y = WorldBaseline - n.npcHeight
	}
	n.realNpcX = n.x + (n.npcWidth)/3
	n.realNpcY = n.y
}

func (n *Npc) Flip() {
	n.vX = n.vX * -1
	n.sens = !n.sens
}

func (n *Npc) Jump() {
	if n.jumping {
		n.falling = false
		n.jumpSpeed += Gravity
		if n.jumpSpeed > 0 {
			n.jumping = false
			n.falling = true
		}
		n.vY = n.jumpSpeed
	}

	if n.falling && n.jumpSpeed < BaseJumpSpeed+1 {
		n.jumpSpeed += Gravity
		n.vY = n.jumpSpeed
	}

	if n.jumpSpeed >= BaseJumpSpeed {
		n.falling = false
		n.jumpSpeed = -BaseJumpSpeed
	}
}

func (n *Npc) Retraite() {
	n.x = RetraiteX
}

func (n *Npc) ManageCollisionBlock(block *Block) {
	nextX := n.realNpcX + n.vX
	nextY := n.realNpcY

	if (block.realX >= nextX+n.realNpcWidth) || // trop à droite
		(block.realX+block.realW <= nextX) || // trop à gauche
		(block.realY >= nextY+n.realNpcHeight) || // trop en bas
		(block.realY+block.realH <= nextY) { // trop en haut
	} else {
		n.Flip()
	}

	nextX = n.realNpcX
	nextY = n.realNpcY + n.vY

	if (block.realX >= nextX+n.realNpcWidth) || // trop à droite
		(block.realX+block.realW <= nextX) || // trop à gauche
		(block.realY >= nextY+n.realNpcHeight) || // trop en bas
		(block.realY+block.realH <= nextY) { // trop en haut
	} else {
		n.vY = 0
	}
}

func (n *Npc) ManageCollisionPlayer(player *Player) {
	nextX := n.realNpcX + n.vX
	nextY := n.realNpcY

	if (player.realPlayerX >= nextX+n.realNpcWidth) || // trop à droite
		(player.realPlayerX+player.realPlayerWidth <= nextX) || // trop à gauche
		(player.realPlayerY >= nextY+n.realNpcHeight) || // trop en bas
		(player.realPlayerY+player.realPlayerHeight <= nextY) { // trop en haut
	} else {
		//n.vX = 0
		if player.realPlayerY+player.realPlayerHeight <= n.realNpcY+n.realNpcHeight/FractionTopHitbox {
			n.Retraite()
			player.miniJumping = true
			player.Minijump()
		} else {
			reset()
		}
	}

	nextX = n.realNpcX
	nextY = n.realNpcY + n.vY

	if (player.realPlayerX >= nextX+n.realNpcWidth) || // trop à droite
		(player.realPlayerX+player.realPlayerWidth <= nextX) || // trop à gauche
		(player.realPlayerY >= nextY+n.realNpcHeight) || // trop en bas
		(player.realPlayerY+player.realPlayerHeight <= nextY) { // trop en haut
	} else {
		//n.vY = 0
		if player.realPlayerY+player.realPlayerHeight <= n.realNpcY+n.realNpcHeight/FractionTopHitbox {
			n.Retraite()
			player.miniJumping = true
			player.Minijump()
		} else {
			reset()
		}
	}

}

func (n *Npc) Len() int {
	return 1
}

func (n *Npc) Dst(i int) (x0, y0, x1, y1 int) {
	return n.x - CameraX, n.y - CameraY, n.x - CameraX + n.npcWidth, n.y - CameraY + n.npcHeight
}

func (n *Npc) Src(i int) (x0, y0, x1, y1 int) {
	if n.sens {
		return n.currentAnimX*n.spriteWidth + n.spriteWidth,
			n.currentAnimY * n.spriteHeight,
			n.currentAnimX * n.spriteWidth,
			n.currentAnimY*n.spriteHeight + n.spriteHeight
	}

	return n.currentAnimX * n.spriteWidth,
		n.currentAnimY * n.spriteHeight,
		n.currentAnimX*n.spriteWidth + n.spriteWidth,
		n.currentAnimY*n.spriteHeight + n.spriteHeight
}

type Npcs struct {
	elems []*Npc
	num   int
}

func (n *Npcs) Update() {
	for _, elem := range n.elems {
		elem.Update()
	}
}

func (n *Npcs) Len() int {
	return n.num
}

func (n *Npcs) Dst(i int) (x0, y0, x1, y1 int) {
	if i >= n.num {
		return 0, 0, 0, 0
	}
	return n.elems[i].Dst(i)

}

func (n *Npcs) Src(i int) (x0, y0, x1, y1 int) {
	if n.num <= i {
		return 0, 0, 0, 0
	}
	return n.elems[i].Src(i)
}
