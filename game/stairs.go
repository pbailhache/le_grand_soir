package game

var (
	stairsUpMatrix   = make([][]bool, 4)
	stairsDownMatrix = make([][]bool, 4)

	stairsUpAndColMatrix = make([][]bool, 5)

	bigStairsMatrix = make([][]bool, 9)
)

func createAllShapeMatrixes() {
	for i := range stairsUpMatrix {
		stairsUpMatrix[i] = make([]bool, 4)
	}
	stairsUpMatrix[0][3] = true

	stairsUpMatrix[1][2] = true
	stairsUpMatrix[1][3] = true

	stairsUpMatrix[2][1] = true
	stairsUpMatrix[2][2] = true
	stairsUpMatrix[2][3] = true

	stairsUpMatrix[3][0] = true
	stairsUpMatrix[3][1] = true
	stairsUpMatrix[3][2] = true
	stairsUpMatrix[3][3] = true

	for i := range stairsDownMatrix {
		stairsDownMatrix[i] = make([]bool, 4)
	}
	stairsDownMatrix[3][3] = true

	stairsDownMatrix[2][2] = true
	stairsDownMatrix[2][3] = true

	stairsDownMatrix[1][1] = true
	stairsDownMatrix[1][2] = true
	stairsDownMatrix[1][3] = true

	stairsDownMatrix[0][0] = true
	stairsDownMatrix[0][1] = true
	stairsDownMatrix[0][2] = true
	stairsDownMatrix[0][3] = true

	for i := range stairsUpAndColMatrix {
		stairsUpAndColMatrix[i] = make([]bool, 4)
	}
	stairsUpAndColMatrix[0][3] = true

	stairsUpAndColMatrix[1][2] = true
	stairsUpAndColMatrix[1][3] = true

	stairsUpAndColMatrix[2][1] = true
	stairsUpAndColMatrix[2][2] = true
	stairsUpAndColMatrix[2][3] = true

	stairsUpAndColMatrix[3][0] = true
	stairsUpAndColMatrix[3][1] = true
	stairsUpAndColMatrix[3][2] = true
	stairsUpAndColMatrix[3][3] = true

	stairsUpAndColMatrix[4][0] = true
	stairsUpAndColMatrix[4][1] = true
	stairsUpAndColMatrix[4][2] = true
	stairsUpAndColMatrix[4][3] = true

	for i := range bigStairsMatrix {
		bigStairsMatrix[i] = make([]bool, 8)
	}
	bigStairsMatrix[0][7] = true

	bigStairsMatrix[1][6] = true
	bigStairsMatrix[1][7] = true

	bigStairsMatrix[2][5] = true
	bigStairsMatrix[2][6] = true
	bigStairsMatrix[2][7] = true

	bigStairsMatrix[3][4] = true
	bigStairsMatrix[3][5] = true
	bigStairsMatrix[3][6] = true
	bigStairsMatrix[3][7] = true

	bigStairsMatrix[4][3] = true
	bigStairsMatrix[4][4] = true
	bigStairsMatrix[4][5] = true
	bigStairsMatrix[4][6] = true
	bigStairsMatrix[4][7] = true

	bigStairsMatrix[5][2] = true
	bigStairsMatrix[5][3] = true
	bigStairsMatrix[5][4] = true
	bigStairsMatrix[5][5] = true
	bigStairsMatrix[5][6] = true
	bigStairsMatrix[5][7] = true

	bigStairsMatrix[6][1] = true
	bigStairsMatrix[6][2] = true
	bigStairsMatrix[6][3] = true
	bigStairsMatrix[6][4] = true
	bigStairsMatrix[6][5] = true
	bigStairsMatrix[6][6] = true
	bigStairsMatrix[6][7] = true

	bigStairsMatrix[7][0] = true
	bigStairsMatrix[7][1] = true
	bigStairsMatrix[7][2] = true
	bigStairsMatrix[7][3] = true
	bigStairsMatrix[7][4] = true
	bigStairsMatrix[7][5] = true
	bigStairsMatrix[7][6] = true
	bigStairsMatrix[7][7] = true

	bigStairsMatrix[8][0] = true
	bigStairsMatrix[8][1] = true
	bigStairsMatrix[8][2] = true
	bigStairsMatrix[8][3] = true
	bigStairsMatrix[8][4] = true
	bigStairsMatrix[8][5] = true
	bigStairsMatrix[8][6] = true
	bigStairsMatrix[8][7] = true
}
