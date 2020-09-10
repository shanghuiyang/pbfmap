package pbfmap

import (
	"fmt"

	"github.com/qedus/osmpbf"
)

// Node ...
type Node struct {
	*osmpbf.Node
}

// Type ...
func (n *Node) Type() ElementType {
	return NodeType
}

// Length ...
func (n *Node) Length() float64 {
	return 0
}

// String ...
func (n *Node) String() string {
	s := fmt.Sprintf("%-7s = %v\n", "node_id", n.ID)
	s += fmt.Sprintf("%-7s = %v\n", "version", n.Info.Version)
	s += fmt.Sprintf("%-7s = %.7f\n", "lat", n.Lat)
	s += fmt.Sprintf("%-7s = %.7f\n", "lon", n.Lon)

	max := 0
	for k := range n.Tags {
		if l := len(k); l > max {
			max = len(k)
		}
	}
	format := fmt.Sprintf("%%-%vs = %%v\n", max)

	s += "tags:\n"
	for k, v := range n.Tags {
		s += fmt.Sprintf(format, k, v)
	}
	return s
}
