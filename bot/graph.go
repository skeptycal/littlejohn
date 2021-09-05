package bot

import (
	"fmt"
	"strings"
	"sync"
)

const (
	graphSep string = " -> "
	newline  byte   = '\n'
)

type (
	// Graph is the graph of rules allowing users to define
	// how rules interact with each other, which has precedence,
	// and which have overriding capabilities.
	Graph struct {
		nodes []*Node
		edges map[Node][]*Node
		lock  sync.RWMutex
	}
)

// AddNode adds a node to the graph
func (g *Graph) AddNode(n *Node) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.nodes = append(g.nodes, n)
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(n1, n2 *Node) {
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
}

func (g *Graph) String() string {
	g.lock.RLock()
	defer g.lock.RUnlock()

	sb := strings.Builder{}
	defer sb.Reset()

	for i := 0; i < len(g.nodes); i++ {
		sb.WriteString(g.nodes[i].String() + graphSep)
		near := g.edges[*g.nodes[i]]
		for j := 0; j < len(near); j++ {
			sb.WriteString(near[j].String() + " ")
		}
		sb.WriteByte(newline)
	}

	return sb.String()
}

type Node struct {
	rule *Rule
}

// temporary type
type Rule string

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.rule)
}
