package main

import "github.com/nikitych1w/leetcode-problems/containers"

func main() {
	graph, err := containers.ReadFromTxt("./data/graph1.txt")
	if err != nil {
		panic(err)
	}

	graph.PrintEdges()
}
