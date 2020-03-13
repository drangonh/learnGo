package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	var maze [][]int
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	//func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
	//Fscanf 扫描从 r 中读取的文本，并将连续由空格分隔的值存储为连续的实参，其格式由 format 决定。它返回成功解析的条目数
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze = make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

type point struct {
	i, j int
}

var dirs = []point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func (p point) add(q point) point {
	return point{p.i + q.i, p.j + q.j}
}

func (p point) at(grid [][]int) (int, bool) {
	//行不能越界
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	//列不能越界
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

func walk(maze [][]int, start point, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			//跳出所有循环
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			// maze at next is 0
			//steps at next is 0
			//start != end
			val, ok := next.at(maze)
			if !ok || val == 1 {
				//跳出当前循环
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			//fmt.Println(cur, steps)
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1
			Q = append(Q, next)
		}
	}

	return steps
}

func main() {
	maze := readMaze("maze/maze.in")
	//fmt.Println(maze)

	steps := walk(maze, point{0, 0}, point{len(maze), len(maze[0])})

	var max int
	for _, row := range steps {
		for _, col := range row {
			fmt.Printf("%3d", col)
			if col > max {
				max = col
			}
		}
		fmt.Println()
	}

	//var walkStep []point
	//var lastStep point
	//x := len(steps) - 1
	//y := len(steps[0]) - 1
	//
	//walkStep = append([]point{}, point{len(steps) - 1, len(steps[0]) - 1})
	//
	//for i := len(steps) - 1; i >= 0; i-- {
	//
	//	for j := len(steps[i]) - 1; j >= 0; j-- {
	//
	//		if math.Abs(x-) == 1 || y-j == 1 {
	//
	//			if max-steps[i][j] == 1 {
	//				x = i
	//				y = j
	//				lastStep = point{i, j}
	//				walkStep = append(
	//					[]point{lastStep,},
	//					walkStep[0:]...
	//				)
	//
	//			}
	//		}
	//
	//	}
	//}
	//
	//fmt.Println(walkStep)
}
