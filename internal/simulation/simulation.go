package simulation

import (
	"github.com/lucasdellatorre/percolation/internal/unionfind"
	"github.com/lucasdellatorre/percolation/internal/util"
)

func Run(n int) {
	u := unionfind.NewUnionFind(n)

	randomNumbers := util.GenerateUniqueRandomNumbers(n * n)

	for i := 0; !u.Percolates(); i++ {
		u.Open(randomNumbers[i])
		u.PrintUf()
	}
}
