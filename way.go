package pbfmap

import (
	"fmt"

	"github.com/qedus/osmpbf"
)

// Way ...
type Way struct {
	*osmpbf.Way
	*PbfMap
	length float64
}

// GetNode ...
func (w *Way) GetNode(id int64) *Node {
	if n, ok := w.Nodes[id]; ok {
		return n
	}
	return nil
}

// Type ...
func (w *Way) Type() ElementType {
	return WayType
}

// Length ...
func (w *Way) Length() float64 {
	if w.length > 0 {
		return w.length
	}
	nodes := []*Node{}
	for _, id := range w.NodeIDs {
		n := w.GetNode(id)
		if n == nil {
			// invalid way feature
			w.length = 0
			return w.length
		}
		nodes = append(nodes, n)
	}
	for i := 0; i < len(nodes)-1; i++ {
		w.length += Length(nodes[i], nodes[i+1])
	}
	return w.length
}

// String ...
func (w *Way) String() string {
	s := fmt.Sprintf("%-7s = %v\n", "way_id", w.ID)
	s += fmt.Sprintf("%-7s = %v\n", "version", w.Info.Version)

	s += "nodes:\n"
	for i, nid := range w.NodeIDs {
		n := w.GetNode(nid)
		if n == nil {
			continue
		}
		s += fmt.Sprintf("%d\tnode_id = %v\t lat = %.7f\t lon = %.7f\n", i+1, n.ID, n.Lat, n.Lon)
	}

	max := 0
	for k := range w.Tags {
		if l := len(k); l > max {
			max = len(k)
		}
	}
	format := fmt.Sprintf("%%-%vs = %%v\n", max)

	s += "tags:\n"
	for k, v := range w.Tags {
		s += fmt.Sprintf(format, k, v)
	}
	return s
}
