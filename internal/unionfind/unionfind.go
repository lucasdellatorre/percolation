package unionfind

import "fmt"

type UnionFind struct {
	Id          []int
	Sz          []int
	BlockedGrid []bool
	N           int
}

func (u *UnionFind) isOpen(i int) bool {
	return !u.BlockedGrid[i]
}

func (u *UnionFind) Open(i int) {
	fmt.Println("Open", i)
	u.BlockedGrid[i] = false

	// Check and join with adjacent open sites
	if i-1 >= 0 && i%u.N != 0 && u.isOpen(i-1) { // left
		fmt.Printf("Union(%d, %d)\n", i, i-1)
		u.union(i, i-1)
	}
	if i+1 < len(u.BlockedGrid) && (i+1)%u.N != 0 && u.isOpen(i+1) { // right
		u.union(i, i+1)
		fmt.Printf("Union(%d, %d)\n", i, i+1)
	}
	if i-u.N >= 0 && u.isOpen(i-u.N) { // top
		fmt.Printf("Union(%d, %d)\n", i, i-u.N)
		u.union(i, i-u.N)
	}
	if i+u.N < len(u.BlockedGrid) && u.isOpen(i+u.N) { // bottom
		fmt.Printf("Union(%d, %d)\n", i, i+u.N)
		u.union(i, i+u.N)
	}
}

func (u *UnionFind) union(p int, q int) {
	i := u.root(p)
	j := u.root(q)

	if i == j {
		return
	}

	if u.Sz[i] < u.Sz[j] {
		u.Id[i] = j
		u.Sz[j] += u.Sz[i]

	} else {
		u.Id[j] = i
		u.Sz[i] += u.Sz[j]

	}
}

func (u *UnionFind) root(i int) int {
	for i != u.Id[i] {
		u.Id[i] = u.Id[u.Id[i]] // path compression
		i = u.Id[i]
	}
	return i
}

func (u *UnionFind) connection(p, q int) bool {
	return u.root(p) == u.root(q)
}

func (u *UnionFind) Percolates() bool {
	return u.connection(u.N*u.N, u.N*u.N+1)
}

func (u *UnionFind) percolationPath() bool { // Todo: make the path that percolates as blue
	return true
}

func NewUnionFind(n int) *UnionFind {
	id := make([]int, n*n+2)
	sz := make([]int, n*n+2)
	grid := make([]bool, n*n)

	initializeIdMatrix(id)
	initializeSizeMatrix(sz)
	initializeGridMatrix(grid)

	u := &UnionFind{
		Id:          id,
		Sz:          sz,
		BlockedGrid: grid,
		N:           n,
	}

	for i := 0; i < n; i++ {
		u.union(i, n*n)
	}

	for i := n * (n - 1); i < n*n; i++ {
		u.union(i, n*n+1)
	}

	return u
}

func initializeIdMatrix(matrix []int) {
	for i := range matrix {
		matrix[i] = i
	}
}

func initializeSizeMatrix(matrix []int) {
	for i := range matrix {
		matrix[i] = 1
	}
}

func initializeGridMatrix(matrix []bool) {
	for i := range matrix {
		matrix[i] = true
	}
}

func (u *UnionFind) printIdMatrix() {
	for i := 0; i < u.N*u.N; i++ {
		if i > 0 && i%u.N == 0 {
			fmt.Println()
		}
		fmt.Printf(" %d ", u.Id[i])
	}
}

func (u *UnionFind) printSizeMatrix() {
	for i := 0; i < u.N*u.N; i++ {
		if i > 0 && i%u.N == 0 {
			fmt.Println()
		}
		fmt.Printf(" %d ", u.Sz[i])
	}
}

func (u *UnionFind) printGridMatrix() {
	for i := 0; i < u.N*u.N; i++ {
		if i > 0 && i%u.N == 0 {
			fmt.Println()
		}
		fmt.Printf(" %t ", u.BlockedGrid[i])
	}
}

func (u *UnionFind) PrintUf() {
	u.printIdMatrix()
	fmt.Println()
	u.printSizeMatrix()
	fmt.Println()
	u.printGridMatrix()
	fmt.Println()
	fmt.Println("Percolates?", u.Percolates())
}