package main

import "fmt"

func main() {
	in := [][]int{{0, 0, 0}, {0, 1, 1}}
	r := floodFill(in, 1, 1, 1)
	fmt.Println(r)
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] != newColor {
		dfs(image, sr, sc, image[sr][sc], newColor)
	}
	return image
}

func dfs(image [][]int, sr int, sc int, oldColor, newColor int) {

	if image[sr][sc] != oldColor {
		return
	}
	image[sr][sc] = newColor

	if sr >= 1 {
		dfs(image, sr-1, sc, oldColor, newColor)
	}
	if sc >= 1 {
		dfs(image, sr, sc-1, oldColor, newColor)
	}
	if sr+1 < len(image) {
		dfs(image, sr+1, sc, oldColor, newColor)
	}
	if sc+1 < len(image[0]) {
		dfs(image, sr, sc+1, oldColor, newColor)
	}
}
