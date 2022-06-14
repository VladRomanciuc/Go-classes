package main

import "fmt"

/*
created a visited map that represents the set of all indices in a graph that have been visited.

Complete the visit function that should mark a node as visited by creating its key.
*/


func visit(m map[int]struct{}, vertex int) {
	m[vertex] = struct{}{}
}

func isVisited(m map[int]struct{}, vertex int) bool {
	_, ok := m[vertex]
	return ok
}

func main() {
	visited := make(map[int]struct{})

	visit(visited, 1)
	visit(visited, 1)
	visit(visited, 1)
	visit(visited, 2)
	visit(visited, 1023)
	visit(visited, 4)
	visit(visited, 10)

	fmt.Println(isVisited(visited, 5))
	fmt.Println(isVisited(visited, 1023))

	fmt.Println(visited)
}
