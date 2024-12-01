package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Graph structure
type Graph struct {
	vertices map[string]map[string]int // adjacency list for weighted graph
	directed bool
}

// NewGraph returns a new Graph
func NewGraph(directed bool) *Graph {
	return &Graph{
		vertices: make(map[string]map[string]int),
		directed: directed,
	}
}

// AddVertex adds a vertex to the graph
func (g *Graph) AddVertex(vertex string) {
	if _, exists := g.vertices[vertex]; !exists {
		g.vertices[vertex] = make(map[string]int)
	}
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(v1, v2 string, weight int) {
	g.AddVertex(v1)
	g.AddVertex(v2)
	g.vertices[v1][v2] = weight
	if !g.directed {
		g.vertices[v2][v1] = weight
	}
}

// BFS performs Breadth-First Search on the graph
func (g *Graph) BFS(start string) []string {
	visited := make(map[string]bool)
	queue := []string{start}
	visited[start] = true
	result := []string{}

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		result = append(result, v)

		for neighbor := range g.vertices[v] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}

	return result
}

// DFS performs Depth-First Search on the graph
func (g *Graph) DFS(start string) []string {
	visited := make(map[string]bool)
	result := []string{}
	g.dfsHelper(start, visited, &result)
	return result
}

func (g *Graph) dfsHelper(v string, visited map[string]bool, result *[]string) {
	visited[v] = true
	*result = append(*result, v)

	for neighbor := range g.vertices[v] {
		if !visited[neighbor] {
			g.dfsHelper(neighbor, visited, result)
		}
	}
}

// Dijkstra finds the shortest path using Dijkstra's algorithm
func (g *Graph) Dijkstra(start string) map[string]int {
	distances := make(map[string]int)
	for vertex := range g.vertices {
		distances[vertex] = math.MaxInt
	}
	distances[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{value: start, priority: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)
		currentVertex := current.value

		for neighbor, weight := range g.vertices[currentVertex] {
			newDist := distances[currentVertex] + weight
			if newDist < distances[neighbor] {
				distances[neighbor] = newDist
				heap.Push(pq, &Item{value: neighbor, priority: newDist})
			}
		}
	}

	return distances
}

// HasCycle detects a cycle in the graph (only for directed graphs)
func (g *Graph) HasCycle() bool {
	visited := make(map[string]bool)
	recStack := make(map[string]bool)

	for vertex := range g.vertices {
		if !visited[vertex] {
			if g.cycleHelper(vertex, visited, recStack) {
				return true
			}
		}
	}

	return false
}

func (g *Graph) cycleHelper(v string, visited, recStack map[string]bool) bool {
	visited[v] = true
	recStack[v] = true

	for neighbor := range g.vertices[v] {
		if !visited[neighbor] && g.cycleHelper(neighbor, visited, recStack) {
			return true
		} else if recStack[neighbor] {
			return true
		}
	}

	recStack[v] = false
	return false
}

// Priority Queue implementation for Dijkstra's algorithm
type Item struct {
	value    string
	priority int
	index    int
}

// PriorityQueue implements a min-heap for Dijkstra's algorithm
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func main() {
	graph := NewGraph(false) // undirected graph

	// Add edges
	graph.AddEdge("A", "B", 1)
	graph.AddEdge("A", "C", 4)
	graph.AddEdge("B", "C", 2)
	graph.AddEdge("B", "D", 5)
	graph.AddEdge("C", "D", 1)

	// BFS
	fmt.Println("BFS starting from A:", graph.BFS("A"))

	// DFS
	fmt.Println("DFS starting from A:", graph.DFS("A"))

	// Dijkstra
	fmt.Println("Shortest paths from A using Dijkstra:", graph.dijkstra("A"))

	// Detect cycle (only works for directed graphs)
	directedGraph := NewGraph(true)
	directedGraph.AddEdge("A", "B", 1)
	directedGraph.AddEdge("B", "C", 1)
	directedGraph.AddEdge("C", "A", 1)
	fmt.Println("Cycle detected in directed graph:", directedGraph.HasCycle())
}

func (g *Graph) dijkstra(start string) map[string]int {
	distances := make(map[string]int)
	for node := range g.vertices {
		distances[node] = math.MaxInt
	}
	distances[start] = 0
	visited := make(map[string]bool)

		for len(visited) < len(g.vertices) {
			currenctNode := getClosestNode(distances, visited)
			visited[currenctNode] = true
			for node , dist:= range g.vertices[currenctNode]{
				if !visited[node]{
					minNodedist := distances[currenctNode] + dist
					if minNodedist < distances[node]{
						distances[node] = minNodedist
					}
				}
			}
		}
		return distances
	}

	func getClosestNode(distances map[string]int, visited map[string]bool) string {
		minNode := math.MaxInt
		var closeNode string
		for node, dist := range distances {
			if !visited[node] && dist < minNode {
				minNode = dist
				closeNode = node
			}
		}
		return closeNode
	}
// func (g *Graph) dijkstra(start string) map[string]int {
// 	distance := make(map[string]int)
// 	for node := range g.vertices {
// 		distance[node] = math.MaxInt
// 	}
// 	distance[start] = 0
// 	visited := make(map[string]bool)
// 	for len(visited) < len(distance) {
// 		currentNode := closestNode(visited, distance)
// 		visited[currentNode] = true
// 		for node, dist := range g.vertices[currentNode] {
// 			minDist := distance[currentNode] + dist
// 			if minDist < distance[node] {
// 				distance[node] = minDist
// 			}
// 		}

// 	}
// 	return distance
// }
// func closestNode(visited map[string]bool, distance map[string]int) string {
// 	minNode := math.MaxInt
// 	var closeNode string
// 	for node, dist := range distance {
// 		if !visited[node] && dist < minNode {
// 			minNode = dist
// 			closeNode = node
// 		}
// 	}
// 	return closeNode
// }
