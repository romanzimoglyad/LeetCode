package main

import (
	"fmt"
)

func main() {
	// rooms1 := [][]int{{2}, {}, {1}}
	// rooms2 := [][]int{{1}, {2}, {0}}
	// rooms3 := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	rooms4 := [][]int{{2}, {0}, {1}}
	res := canVisitAllRooms(rooms4)
	fmt.Println(res)
}
func canVisitAllRooms(rooms [][]int) bool {
	globalKeys := make(map[int]struct{})
	opened := make(map[int]struct{})
	opened[0] = struct{}{}
	globalKeys[0] = struct{}{}
	for i := range rooms {
		if len(opened) == len(rooms) {
			return true
		}
		dfs(rooms, opened, i, globalKeys)
	}

	return len(opened) == len(rooms)
}

func dfs(rooms [][]int, visited map[int]struct{}, roomNumber int, globalKeys map[int]struct{}) {

	if _, ok := globalKeys[roomNumber]; ok {
		visited[roomNumber] = struct{}{}
		for _, elem := range rooms[roomNumber] {
			if _, ok := visited[elem]; !ok {
				globalKeys[elem] = struct{}{}
				dfs(rooms, visited, elem, globalKeys)
			}
		}
	}
}
