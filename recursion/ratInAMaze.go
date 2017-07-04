package recursion

// Rat In a Maze Problem (Backtracking)
// A maze is given as N*N binary matrix. Source point is maze[0][0] and Destination is maze[N-1][N-1].
// A rat starts at source and has to reach destination by going only RIGHT and DOWN
// 0 in maze means, it's blocked so rat can't go to that shell, 1 means rat can take that shell in path

func ratInMaze(maze [][]int) [][]int {

	// Initialize solution maze; starting index/point i.e. (0,0) is initialzed to 1
	solutionMaze := [][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	done := false
	findPath(maze, &solutionMaze, 0, 0, &done) // 0,0 is the starting point in the maze
	return solutionMaze
}

func findPath(maze [][]int, solutionMaze *[][]int, row, col int, done *bool) {
	size := len(maze[0])

	//Accept case
	if row == size-1 && col == size-1 {
		(*solutionMaze)[size-1][size-1] = 1
		*done = true
		return
	}

	// Steps : first go right and then go down
	if *done != true && isSafeToGo(maze, row, col+1) { // right
		(*solutionMaze)[row][col+1] = 1
		findPath(maze, solutionMaze, row, col+1, done)
	}
	if *done != true && isSafeToGo(maze, row+1, col) { // down
		(*solutionMaze)[row+1][col] = 1
		findPath(maze, solutionMaze, row+1, col, done)
	}

	if *done != true { // backtracking case
		(*solutionMaze)[row][col] = 0
	}
}

// Checks if position at maze[row][col] is VALID
func isSafeToGo(maze [][]int, row, col int) bool {
	size := len(maze[0])
	if row < size && col < size && maze[row][col] == 1 {
		return true
	}
	return false
}
