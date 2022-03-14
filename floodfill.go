package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const length = 20
const row = length
const col = length * 2

var maps [row][col]string

func initMap() {
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 || j == 0 || i == row-1 || j == col-1 {
				maps[i][j] = "#"
			} else {
				maps[i][j] = " "
			}
		}
	}
}

func printMap() {
	time.Sleep(time.Millisecond * 100)
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("%s", maps[i][j])
		}
		fmt.Println(" ")
	}
}

func floodFill(x, y int) {
	printMap()
	if maps[x+1][y] != "#" && maps[x+1][y] != "*" {
		maps[x+1][y] = "*"
		floodFill(x+1, y)
	}
	if maps[x-1][y] != "#" && maps[x-1][y] != "*" {
		maps[x-1][y] = "*"
		floodFill(x-1, y)
	}
	if maps[x][y-1] != "#" && maps[x][y-1] != "*" {
		maps[x][y-1] = "*"
		floodFill(x, y-1)
	}
	if maps[x][y+1] != "#" && maps[x][y+1] != "*" {
		maps[x][y+1] = "*"
		floodFill(x, y+1)
	}
}

func main() {
	initMap()
	floodFill(3, 3)
}
