package pbfmap

import (
	"fmt"

	"github.com/qedus/osmpbf"
)

var memberType2Str = map[osmpbf.MemberType]string{
	osmpbf.NodeType:     "node",
	osmpbf.WayType:      "way",
	osmpbf.RelationType: "relation",
}

// Relation ...
type Relation struct {
	*osmpbf.Relation
	*PbfMap
}

// GetMember ...
func (r *Relation) GetMember(id int64) *osmpbf.Member {
	for _, m := range r.Members {
		if m.ID == id {
			return &m
		}
	}
	return nil
}

// Type ...
func (r *Relation) Type() ElementType {
	return RelationType
}

// Length ...
func (r *Relation) Length() float64 {
	return 0
}

// String ...
func (r *Relation) String() string {
	s := fmt.Sprintf("%-11s = %v\n", "relation_id", r.ID)
	s += fmt.Sprintf("%-11s = %v\n", "version", r.Info.Version)

	s += "members:\n"
	for i, m := range r.Members {
		s += fmt.Sprintf("%d\tmemeber_id = %v\t type = %v\t role = %v\n", i+1, m.ID, memberType2Str[m.Type], m.Role)
	}

	max := 0
	for k := range r.Tags {
		if l := len(k); l > max {
			max = len(k)
		}
	}
	format := fmt.Sprintf("%%-%vs = %%v\n", max)

	s += "tags:\n"
	for k, v := range r.Tags {
		s += fmt.Sprintf(format, k, v)
	}
	return s
}
