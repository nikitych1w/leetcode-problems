package containers

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	id    uint
	value interface{}
}

type Edge struct {
	source, sink *Node
	isDirected   bool
	weight       int
}

type Graph struct {
	nodeCount uint
	edges     []*Edge
	nodes     map[uint]*Node
}

func NewGraph() *Graph {
	return &Graph{
		nodeCount: 0,
		edges:     nil,
		nodes:     map[uint]*Node{},
	}
}

func (graph *Graph) AddEdge(sourceid, sinkid uint, isDirected bool, weight int) {
	if len(graph.nodes) > 1 {
		graph.edges = append(graph.edges, &Edge{graph.nodes[sourceid],
			graph.nodes[sinkid], isDirected, weight})
	}
}

func (graph *Graph) AddNode(value interface{}) {
	graph.nodes[graph.nodeCount] = &Node{graph.nodeCount, value}
	graph.nodeCount++
}

func ReadFromTxt(path string) (*Graph, error) {
	graph := NewGraph()

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	matrix, err := ioutil.ReadAll(file)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	nodeCount, err := strconv.Atoi(string(matrix[0]))

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	for i := 0; i < nodeCount*nodeCount; i++ {
		graph.AddNode(nil)
	}

	lines := strings.Split(string(matrix[2:]), "\n")

	for i, el := range lines {
		edges := strings.Split(el, ";")
		for j, edge := range edges {
			sourceId := j + i*nodeCount
			sinkId := i + j*nodeCount

			weight, err := strconv.Atoi(edge)

			if err != nil {
				logrus.Error(err)
				return nil, err
			}

			if weight != 0 {
				graph.AddEdge(uint(sourceId), uint(sinkId), false, weight)
			}
		}
	}

	return graph, nil
}

func (graph *Graph) PrintEdges() {
	if graph.nodeCount == 0 {
		logrus.Warn("Graph is empty")
		return
	}

	for i, el := range graph.edges {
		if el.isDirected {
			fmt.Printf("Edge ID [%d], %d (%s) ————-> %d (%s)\n", i, el.source.id,
				el.source.value, el.sink.id, el.sink.value)
		} else {
			fmt.Printf("Edge ID [%d] │ %d ——————(%d)—————— %d\n", i, el.source.id, el.weight, el.sink.id)
		}
	}
}
