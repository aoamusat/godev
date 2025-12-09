// Package main implements a simple directed graph data structure using adjacency lists
package graph

// Graph represents a directed graph as an adjacency list using a map of maps
// The outer map's keys are vertex names (strings)
// The inner map's keys are adjacent vertex names, with bool values (true indicates an edge exists)
var Graph = make(map[string]map[string]bool)

func main() {
	// Example usage:
	addEdge("A", "B")
	addEdge("A", "C")
	addEdge("B", "C")

	println("Edge from A to B:", hasEdge("A", "B")) // Output: true
	println("Edge from A to C:", hasEdge("A", "C")) // Output: true
	println("Edge from B to A:", hasEdge("B", "A")) // Output: false
}

// addEdge adds a directed edge from vertex 'from' to vertex 'to'
// If the 'from' vertex doesn't exist in the graph, it is created
func addEdge(from string, to string) {
	edges := Graph[from]

	if edges == nil {
		edges = make(map[string]bool)
		Graph[from] = edges
	}
	edges[to] = true
}

// hasEdge checks if there is a directed edge from vertex 'from' to vertex 'to'
// Returns true if the edge exists, false otherwise
func hasEdge(from, to string) bool {
	return Graph[from][to]
}
