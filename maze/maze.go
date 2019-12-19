package main

import (
	"fmt"
	"os"
)

func main() {
	maze := mazeRead("maze/maze.in")
	step := walk(maze, point{0, 0}, point{6, 5})
	for i := range step {
		for j := range step[i] {
			fmt.Printf("%d ", step[i][j])
		}
		fmt.Println()
	}
}

func mazeRead(filename string) [][]int {
	file, e := os.Open(filename)
	if e != nil {
		panic(e)
	}
	var row, col int
	//从文件中读取第一行数据
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)

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
	{0, 1}, {1, 0}, {-1, 0}, {0, -1},
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {
	step := make([][]int, len(maze))
	for i := range step {
		step[i] = make([]int, len(maze[i]))
	}

	//队列
	queue := []point{start}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur == end {
			break
		}

		for _, i := range dirs {
			next := cur.add(i)
			//判断，1、不能等于起点
			if next == start {
				continue
			}
			//2、不能是墙
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			//3、不能在step中有值
			val, ok = next.at(step)
			if !ok || val != 0 {
				continue
			}

			//将next放入step中
			val, _ = cur.at(step)
			step[next.i][next.j] = val + 1
			//将next放入队列中
			queue = append(queue, next)
		}
	}

	return step
}
