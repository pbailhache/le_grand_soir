package game

import (
	"fmt"
	"image"
	_ "image/png"
	"io/ioutil"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"

	"github.com/golang/freetype/truetype"

	"image/color"

	"math/rand"

	"time"

	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var player *Player
var camera *Camera
var barre *Barre

var toCollide []*Block

var blocks = &Blocks{make([]*Block, 37), 37}
var stairsUp = &Blocks{make([]*Block, 10), 10}
var stairsDown = &Blocks{make([]*Block, 10), 10}
var stairsUpAndCol = &Blocks{make([]*Block, 14), 14}
var bigStairs = &Blocks{make([]*Block, 44), 44}

var drapeaux = &Blocks{make([]*Block, 2), 2}

var twoByTwo = &Blocks{make([]*Block, 3), 3}
var twoByThree = &Blocks{make([]*Block, 1), 1}
var twoByFour = &Blocks{make([]*Block, 2), 2}

var nbEnem = 20

var porcs = &Npcs{make([]*Npc, nbEnem), nbEnem}
var camarades = &Npcs{make([]*Npc, nbEnem), nbEnem}

var worldArray [36]int
var world = &World{make([]*WorldTile, 36), 36}

var textImage *ebiten.Image
var isWin bool

func update(screen *ebiten.Image) error {
	frameStart := time.Now()

	CommunismLevel++

	if IsCommunist && player.x <= 200 {
		win("Victoire du Communisme")
	} else if !IsCommunist && player.x >= 8800 {
		win("Echec de la société")
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if !player.sens {
			player.Flip()
		}
		player.vX = -PlayerSpeed
		player.currentAnimX = (player.currentAnimX + 1) % player.nbAnimX
		if IsCommunist && player.x <= camera.x+ScreenWidth-ScreenWidth/4 {
			camera.x -= PlayerSpeed
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if player.sens {
			player.Flip()
		}
		player.vX = PlayerSpeed
		player.currentAnimX = (player.currentAnimX + 1) % player.nbAnimX

		if !IsCommunist && player.x >= camera.x+ScreenWidth/4 {
			camera.x += PlayerSpeed
		}
	}

	if !ebiten.IsKeyPressed(ebiten.KeyRight) && !ebiten.IsKeyPressed(ebiten.KeyLeft) {
		player.vX = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) && !player.falling {
		player.jumping = true
	}
	player.Jump()

	for _, block := range toCollide {
		player.ManageCollision(block)
	}

	for _, porc := range porcs.elems {
		for _, block := range toCollide {
			porc.ManageCollisionBlock(block)
		}
		if IsCommunist {
			porc.ManageCollisionPlayer(player)
		}
	}

	for _, camarade := range camarades.elems {
		for _, block := range toCollide {
			camarade.ManageCollisionBlock(block)
		}
		if !IsCommunist {
			camarade.ManageCollisionPlayer(player)
		}
	}

	player.Update()
	porcs.Update()
	camarades.Update()
	barre.Update()
	camera.Update()
	if ebiten.IsRunningSlowly() {
		return nil
	}

	cameraPos := &ebiten.DrawImageOptions{
		ImageParts: camera,
	}
	if err := screen.DrawImage(cameraImg, cameraPos); err != nil {
		return err
	}

	worldDrawOptions := &ebiten.DrawImageOptions{
		ImageParts: world,
	}
	if err := screen.DrawImage(backgroundTileSheetImg, worldDrawOptions); err != nil {
		return err
	}

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(ScreenWidth-BarreW-10, 10)
	if err := screen.DrawImage(bgBarreImg, opts); err != nil {
		return err
	}

	blocksPos := &ebiten.DrawImageOptions{
		ImageParts: blocks,
	}
	if err := screen.DrawImage(blockImg, blocksPos); err != nil {
		return err
	}
	stairUpPos := &ebiten.DrawImageOptions{
		ImageParts: stairsUp,
	}
	if err := screen.DrawImage(blockImg, stairUpPos); err != nil {
		return err
	}
	stairDownPos := &ebiten.DrawImageOptions{
		ImageParts: stairsDown,
	}
	if err := screen.DrawImage(blockImg, stairDownPos); err != nil {
		return err
	}
	stairsUpAndColPos := &ebiten.DrawImageOptions{
		ImageParts: stairsUpAndCol,
	}
	if err := screen.DrawImage(blockImg, stairsUpAndColPos); err != nil {
		return err
	}

	rdapeauxPos := &ebiten.DrawImageOptions{
		ImageParts: drapeaux,
	}
	if err := screen.DrawImage(drapeauxImg, rdapeauxPos); err != nil {
		return err
	}

	twoByTwoPos := &ebiten.DrawImageOptions{
		ImageParts: twoByTwo,
	}
	if err := screen.DrawImage(twoByTwoImg, twoByTwoPos); err != nil {
		return err
	}
	bigStairsPos := &ebiten.DrawImageOptions{
		ImageParts: bigStairs,
	}
	if err := screen.DrawImage(blockImg, bigStairsPos); err != nil {
		return err
	}

	twoByThreePos := &ebiten.DrawImageOptions{
		ImageParts: twoByThree,
	}
	if err := screen.DrawImage(twoByThreeImg, twoByThreePos); err != nil {
		return err
	}

	twoByFourPos := &ebiten.DrawImageOptions{
		ImageParts: twoByFour,
	}
	if err := screen.DrawImage(twoByFourImg, twoByFourPos); err != nil {
		return err
	}

	porcsPos := &ebiten.DrawImageOptions{
		ImageParts: porcs,
	}
	if err := screen.DrawImage(porcCapitalisteImg, porcsPos); err != nil {
		return err
	}

	camaradesPos := &ebiten.DrawImageOptions{
		ImageParts: camarades,
	}
	if err := screen.DrawImage(camaradeImg, camaradesPos); err != nil {
		return err
	}

	playerPos := &ebiten.DrawImageOptions{
		ImageParts: player,
	}
	if err := screen.DrawImage(playerImg, playerPos); err != nil {
		return err
	}

	innerBarrePos := &ebiten.DrawImageOptions{
		ImageParts: barre,
	}
	if err := screen.DrawImage(innerBarreImg, innerBarrePos); err != nil {
		return err
	}

	if isWin {
		if err := screen.DrawImage(textImage, &ebiten.DrawImageOptions{}); err != nil {
			return err
		}
	}

	time.Sleep(time.Second/60 - time.Since(frameStart))

	return nil
}

func reset() {
	IsCommunist = false
	CommunismLevel = 0
	camera = &Camera{
		x: 0,   //player.x + player.playerWidth/2 + ScreenWidth/2,
		y: 108, //player.playerHeight/2 + (ScreenHeight/2 + WorldBaseline + player.playerHeight),
	}
	backgroundTileSheetImgImgWidth, backgroundTileSheetImgImgHeight := backgroundTileSheetImg.Size()
	for i, sprite := range worldArray {
		world.WorldTiles[i] = &WorldTile{
			imageWidth:  backgroundTileSheetImgImgWidth,
			imageHeight: backgroundTileSheetImgImgHeight,

			w: 400,
			h: 720,

			numSprite: sprite,

			x: 250 * i,
			y: 0,
		}
	}

	barre = &Barre{
		x: ScreenWidth - BarreW - 10 + 3,
		y: 10 + 3,
		w: BarreW - 6,
		h: BarreH - 6,
	}

	w, h := playerImg.Size()
	x, y := 0, WorldBaseline-51
	playerWidth, playerHeight := 100, 102
	player = &Player{
		imageWidth:  w,
		imageHeight: h,

		x: x,
		y: y,

		// Le sens du perso, false vers la DROITE
		// le true sens, vers la GAUCHE
		sens: false,

		currentAnimX: 0,
		currentAnimY: 0,
		nbAnimX:      7,
		nbAnimY:      2,

		spriteWidth:  100,
		spriteHeight: 102,

		playerWidth:      playerWidth,
		playerHeight:     playerHeight,
		realPlayerX:      x + playerWidth/3,
		realPlayerY:      y + 10,
		realPlayerWidth:  playerWidth / 3,
		realPlayerHeight: playerHeight,

		jumpSpeed:     -BaseJumpSpeed,
		miniJumpSpeed: -BaseMiniJumpSpeed,
		miniFalling:   true,
		miniJumping:   false,
		falling:       true,
		jumping:       false,
	}
	generateNPC()

}

func generateNPC() {
	camaradeImgWidth, camaradeImgHeight := camaradeImg.Size()
	for i := range camarades.elems {

		randomSens := rand.Intn(1)
		vX := PlayerSpeed
		if randomSens == 0 {
			vX = vX * -1
		}

		var randomX int
		if IsCommunist {
			randomX = player.x + 100 + 700 + rand.Intn(WorldWidth-player.x+100+700)
		} else {
			randomX = rand.Intn(WorldWidth - player.x - 700)
		}
		y := 1500

		camarades.elems[i] = &Npc{
			imageWidth:  camaradeImgWidth,
			imageHeight: camaradeImgHeight,

			x: randomX,
			y: y,

			// Le sens du perso, false vers la DROITE
			// le true sens, vers la GAUCHE
			sens: randomSens == 0,

			currentAnimX: 0,
			currentAnimY: 0,
			nbAnimX:      7,
			nbAnimY:      2,

			realNpcX:     randomX + 100/3,
			realNpcY:     y + 10,
			spriteWidth:  100,
			spriteHeight: 102,

			npcWidth:      100,
			npcHeight:     102,
			realNpcWidth:  100 / 3,
			realNpcHeight: 102,

			jumpSpeed: -BaseJumpSpeed,
			falling:   true,
			jumping:   false,
			vX:        vX,
			maxMovX:   200,
		}

	}

	porcCapitalisteImgWidth, porcCapitalisteImgHeight := porcCapitalisteImg.Size()
	for i := range porcs.elems {

		randomSens := rand.Intn(1)
		vX := PlayerSpeed
		if randomSens == 0 {
			vX = vX * -1
		}

		var randomX int
		if IsCommunist {
			randomX = player.x + 100 + 700 + rand.Intn(WorldWidth-player.x+100+700)
		} else {
			randomX = rand.Intn(WorldWidth - player.x - 700)
		}
		y := 1500

		porcs.elems[i] = &Npc{
			imageWidth:  porcCapitalisteImgWidth,
			imageHeight: porcCapitalisteImgHeight,

			x: randomX,
			y: y,

			// Le sens du perso, false vers la DROITE
			// le true sens, vers la GAUCHE
			sens: randomSens == 0,

			currentAnimX: 0,
			currentAnimY: 0,
			nbAnimX:      7,
			nbAnimY:      2,

			realNpcX:     randomX + 100/3,
			realNpcY:     y + 10,
			spriteWidth:  100,
			spriteHeight: 102,

			npcWidth:      100,
			npcHeight:     102,
			realNpcWidth:  100 / 3,
			realNpcHeight: 102,

			jumpSpeed: -BaseJumpSpeed,
			falling:   true,
			jumping:   false,
			vX:        vX,
			maxMovX:   200,
		}

	}
}

func Run() {
	rand.Seed(time.Now().Unix())
	loadImages()
	createAllShapeMatrixes()

	// 0 tile basic blanche entreprise
	worldArray = [36]int{0, 3, 0, 0, 2, 0, 4, 3, 3, 3, 3, 1, 0, 0, 3, 0, 0, 2, 0, 1, 0, 0, 0, 4, 3, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 4}

	reset()
	w, h := blockImg.Size()
	y := WorldBaseline - 200
	blocks.blocks[0] = CreateBlockSpecial(650, y, w, h)
	blocks.blocks[1] = CreateBlock(850, y, w, h)
	blocks.blocks[2] = CreateBlockSpecial(900, y, w, h)
	blocks.blocks[3] = CreateBlock(950, y, w, h)
	blocks.blocks[4] = CreateBlockSpecial(1000, y, w, h)
	blocks.blocks[5] = CreateBlock(3250, y, w, h)
	blocks.blocks[6] = CreateBlock(3350, y, w, h)
	blocks.blocks[7] = CreateBlock(4250, y, w, h)
	blocks.blocks[8] = CreateBlockSpecial(4300, y, w, h)
	blocks.blocks[9] = CreateBlockSpecial(4500, y, w, h)
	blocks.blocks[10] = CreateBlockSpecial(4650, y, w, h)
	blocks.blocks[11] = CreateBlockSpecial(4800, y, w, h)
	blocks.blocks[12] = CreateBlock(5000, y, w, h)
	blocks.blocks[13] = CreateBlock(5450, y, w, h)
	blocks.blocks[14] = CreateBlock(5500, y, w, h)
	blocks.blocks[15] = CreateBlock(7100, y, w, h)
	blocks.blocks[16] = CreateBlock(7150, y, w, h)
	blocks.blocks[17] = CreateBlockSpecial(7200, y, w, h)
	blocks.blocks[18] = CreateBlock(7250, y, w, h)

	y = WorldBaseline - 250
	blocks.blocks[19] = CreateBlockSpecial(4000, y, w, h)

	y = WorldBaseline - 400
	blocks.blocks[20] = CreateBlockSpecial(950, y, w, h)
	blocks.blocks[21] = CreateBlock(3400, y, w, h)
	blocks.blocks[22] = CreateBlock(3450, y, w, h)
	blocks.blocks[23] = CreateBlock(3500, y, w, h)
	blocks.blocks[24] = CreateBlock(3550, y, w, h)
	blocks.blocks[25] = CreateBlock(3600, y, w, h)
	blocks.blocks[26] = CreateBlock(3650, y, w, h)
	blocks.blocks[27] = CreateBlock(3850, y, w, h)
	blocks.blocks[28] = CreateBlock(3900, y, w, h)
	blocks.blocks[29] = CreateBlock(3950, y, w, h)
	blocks.blocks[30] = CreateBlockSpecial(4650, y, w, h)
	blocks.blocks[31] = CreateBlock(5150, y, w, h)
	blocks.blocks[32] = CreateBlock(5200, y, w, h)
	blocks.blocks[33] = CreateBlock(5400, y, w, h)
	blocks.blocks[34] = CreateBlockSpecial(5450, y, w, h)
	blocks.blocks[35] = CreateBlockSpecial(5500, y, w, h)
	blocks.blocks[36] = CreateBlock(5550, y, w, h)

	w, h = twoByTwoImg.Size()
	y = WorldBaseline - h + 20
	twoByTwo.blocks[0] = CreateBlock(1200, y, w, h*2)
	twoByTwo.blocks[1] = CreateBlock(6900, y, w, h*2)
	twoByTwo.blocks[2] = CreateBlock(7600, y, w, h*2)

	w, h = twoByThreeImg.Size()
	y = WorldBaseline - h + 20
	twoByThree.blocks[0] = CreateBlock(1600, y, w, h*2)

	w, h = twoByFourImg.Size()
	y = WorldBaseline - h + 20
	twoByFour.blocks[0] = CreateBlock(1950, y, w, h*2)
	twoByFour.blocks[1] = CreateBlock(2400, y, w, h*2)

	w, h = drapeauxImg.Size()
	y = WorldBaseline - h
	drapeaux.blocks[0] = CreateBlock(200, y, w, h*2)
	drapeaux.blocks[1] = CreateBlock(8800, y, w, h*2)

	//Ajout des formes de blocks

	w, h = blockImg.Size()

	stairsUpShape := createShape(5750, WorldBaseline-h*2, stairsUpMatrix)
	for i, block := range stairsUpShape.getBlocks(w, h) {
		stairsUp.blocks[i] = block
	}
	stairsUpAndColShape := createShape(6150, WorldBaseline-h*2, stairsUpAndColMatrix)
	for i, block := range stairsUpAndColShape.getBlocks(w, h) {
		stairsUpAndCol.blocks[i] = block
	}
	stairsDownShape1 := createShape(6550, WorldBaseline-h*2, stairsDownMatrix)
	// stairsDownShape2 := createShape(6640, WorldBaseline-h*4, stairsDownMatrix)
	// tmp := append(stairsDownShape1.getBlocks(w, h), stairsDownShape2.getBlocks(w, h)...)
	for i, block := range stairsDownShape1.getBlocks(w, h) {
		stairsDown.blocks[i] = block
	}

	bigStairsShape := createShape(7700, WorldBaseline-h*4, bigStairsMatrix)
	for i, block := range bigStairsShape.getBlocks(w, h) {

		bigStairs.blocks[i] = block
	}

	//Box dee collisions
	toCollide = blocks.blocks
	toCollide = append(toCollide, stairsUp.blocks...)
	toCollide = append(toCollide, stairsUpAndCol.blocks...)
	toCollide = append(toCollide, stairsDown.blocks...)
	toCollide = append(toCollide, bigStairs.blocks...)
	toCollide = append(toCollide, twoByTwo.blocks...)
	toCollide = append(toCollide, twoByThree.blocks...)
	toCollide = append(toCollide, twoByFour.blocks...)

	if err := ebiten.Run(update, ScreenWidth, ScreenHeight, 1, "Le Grand Soir"); err != nil {
		log.Fatal(err)
	}
}

func win(text string) {
	fmt.Printf("\n%s\n", text)
	os.Exit(0)
}

func parseFont() {
	f, err := ebitenutil.OpenFile("resources/font/mplus-1p-regular.ttf")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = f.Close()
	}()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	tt, err := truetype.Parse(b)
	if err != nil {
		log.Fatal(err)
	}
	w, h := textImage.Size()
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	const size = 24
	const dpi = 72
	d := &font.Drawer{
		Dst: dst,
		Src: image.White,
		Face: truetype.NewFace(tt, &truetype.Options{
			Size:    size,
			DPI:     dpi,
			Hinting: font.HintingFull,
		}),
	}
	y := size
	d.Dot = fixed.P(ScreenWidth/2-size, ScreenHeight/2)
	d.DrawString("ddddddddddddddddddddddddddddddddddddddddddddddddddddd")
	y += size

	textImage.ReplacePixels(dst.Pix)
}

func loadImages() {
	var err error
	blockImg, _, err = ebitenutil.NewImageFromFile("resources/images/block.png", ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}

	textImage, err = ebiten.NewImage(ScreenWidth, ScreenHeight, ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}

	twoByTwoImg, _, err = ebitenutil.NewImageFromFile("resources/images/block_2x2.png", ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}

	drapeauxImg, _, err = ebitenutil.NewImageFromFile("resources/images/flag.png", ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}

	twoByThreeImg, _, err = ebitenutil.NewImageFromFile("resources/images/block_2x3.png", ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}

	twoByFourImg, _, err = ebitenutil.NewImageFromFile("resources/images/block_2x4.png", ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}

	playerImg, _, err = ebitenutil.NewImageFromFile("resources/images/hero.png", ebiten.FilterLinear)
	if err != nil {
		log.Fatal(err)
	}
	cameraImg, err = ebiten.NewImage(ScreenWidth, ScreenHeight, ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	cameraImg.Fill(color.Black)

	bgBarreImg, err = ebiten.NewImage(BarreW, BarreH, ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	bgBarreImg.Fill(color.Black)

	innerBarreImg, err = ebiten.NewImage(BarreW-4, BarreH-4, ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	innerBarreImg.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})

	porcCapitalisteImg, _, err = ebitenutil.NewImageFromFile("resources/images/boss.png", ebiten.FilterLinear)
	if err != nil {
		log.Fatal(err)
	}

	camaradeImg, _, err = ebitenutil.NewImageFromFile("resources/images/worker.png", ebiten.FilterLinear)
	if err != nil {
		log.Fatal(err)
	}

	backgroundTileSheetImg, _, err = ebitenutil.NewImageFromFile("resources/images/world.png", ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
}
