package algo

func SolveMaze(maze [][]int) []struct{ x, y int } {
	start := struct{ x, y int }{0, 0}
	end := struct{ x, y int }{len(maze) - 1, len(maze[0]) - 1}

	stack := []struct {
		position struct{ x, y int }
		path     []struct{ x, y int }
	}{{start, []struct{ x, y int }{start}}}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		x, y := current.position.x, current.position.y

		if x == end.x && y == end.y {
			return current.path
		}

		for _, move := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			dx, dy := move[0], move[1]
			nx, ny := x+dx, y+dy

			if 0 <= nx && nx < len(maze) && 0 <= ny && ny < len(maze[0]) && maze[nx][ny] == 0 {
				newPath := make([]struct{ x, y int }, len(current.path))
				copy(newPath, current.path)
				stack = append(stack, struct {
					position struct{ x, y int }
					path     []struct{ x, y int }
				}{struct{ x, y int }{nx, ny}, append(newPath, struct{ x, y int }{nx, ny})})
				maze[nx][ny] = -1
			}
		}
	}

	return nil
}
