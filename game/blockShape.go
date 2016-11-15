package game

type Shape struct {
	x int
	y int

	shape [][]bool
}

func createShape(x, y int, shape [][]bool) *Shape {
	return &Shape{x, y, shape}
}

func (s *Shape) getBlocks(w, h int) []*Block {
	var blocks []*Block

	for i := range s.shape {
		for j := range s.shape[i] {
			currentJ := len(s.shape[i]) - 1 - j
			if s.shape[i][currentJ] {
				blocks = append(blocks, CreateBlock(s.x+i*50, s.y+(h/2)*currentJ+j*20, w, h))
			}
		}
	}

	return blocks
}
