package main

import (
	"container/list"
	"fmt"
)

func main() {
	// rooms1 := [][]int{{2}, {}, {1}}
	// rooms2 := [][]int{{1}, {2}, {0}}
	// rooms3 := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	// rooms4 := [][]int{{2}, {0}, {1}}
	rooms5 := [][]int{{1}, {2}, {3}, {0}}
	res := canVisitAllRooms(rooms5)
	fmt.Println(res)
}
func canVisitAllRooms(rooms [][]int) bool {
	q := list.New()

	seen := make(map[int]bool)
	seen[0] = true
	q.PushFront(0)

	for q.Len() != 0 {
		el := q.Front()
		for _, rs := range rooms[el.Value.(int)] {
			if _, ok := seen[rs]; !ok {
				seen[rs] = true
				q.PushFront(rs)
			}
		}
		q.Remove(el)
	}

	return len(seen) == len(rooms)
}
