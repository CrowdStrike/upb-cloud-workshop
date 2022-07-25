package main

func add(x int, y int) int {
	return x + y
} 

func swap(x, y string) (string, string) {
	return y, x
}

func getNextPos(x, y int) (nextX, nextY int) {
	nextX = (y + 2) / 3
	nextY = x + 1
	return
}
