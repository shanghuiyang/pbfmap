package pbfmap

import (
	"errors"
	"io"
	"os"
	"runtime"

	"github.com/qedus/osmpbf"
)

// PbfMap ...
type PbfMap struct {
	Nodes     map[int64]*Node
	Ways      map[int64]*Way
	Relations map[int64]*Relation
}

// New ...
func New() *PbfMap {
	return &PbfMap{
		Nodes:     map[int64]*Node{},
		Ways:      map[int64]*Way{},
		Relations: map[int64]*Relation{},
	}
}

// Load ...
func (pbf *PbfMap) Load(file string) error {
	pbf.Nodes = map[int64]*Node{}
	pbf.Ways = map[int64]*Way{}
	pbf.Relations = map[int64]*Relation{}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
		}
	}()

	d := osmpbf.NewDecoder(f)
	d.SetBufferSize(osmpbf.MaxBlobSize)
	if err := d.Start(runtime.GOMAXPROCS(-1)); err != nil {
		return err
	}

	for {
		v, err := d.Decode()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch v := v.(type) {
		case *osmpbf.Node:
			n := &Node{Node: v}
			pbf.Nodes[n.ID] = n
		case *osmpbf.Way:
			w := &Way{Way: v, PbfMap: pbf}
			pbf.Ways[w.ID] = w
		case *osmpbf.Relation:
			r := &Relation{Relation: v, PbfMap: pbf}
			pbf.Relations[r.ID] = r
		default:
			return errors.New("unknown type")
		}
	}
	return nil
}

// GetNode ...
func (pbf *PbfMap) GetNode(id int64) *Node {
	n, ok := pbf.Nodes[id]
	if !ok {
		return nil
	}
	return n
}

// GetWay ...
func (pbf *PbfMap) GetWay(id int64) *Way {
	w, ok := pbf.Ways[id]
	if !ok {
		return nil
	}
	return w
}

// GetRelation ...
func (pbf *PbfMap) GetRelation(id int64) *Relation {
	r, ok := pbf.Relations[id]
	if !ok {
		return nil
	}
	return r
}
