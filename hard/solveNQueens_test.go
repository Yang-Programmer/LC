package hard

import (
	"fmt"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	nQueens := solveNQueens(4)
	fmt.Println(nQueens)
}
