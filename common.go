package pbfmap

import (
	"math"
)

// ElementType ...
type ElementType int

// EarthRadius is the radius of the Earth, UTM, WGS84, in meter
const EarthRadius = float64(6378137)

const (
	// UnknowType ...
	UnknowType ElementType = iota
	// NodeType ...
	NodeType
	// WayType ...
	WayType
	// RelationType ...
	RelationType
)

// Element ...
type Element interface {
	Type() ElementType
	Length() float64
	String() string
}

// Length calc the length between n1 and n2 in meter
func Length(n1 *Node, n2 *Node) float64 {
	rad1 := Rad(n1.Lat)
	rad2 := Rad(n2.Lat)
	a := rad1 - rad2
	b := Rad(n1.Lon) - Rad(n2.Lon)
	s := 2 * math.Asin(math.Sqrt(math.Pow(math.Sin(a/2), 2)+math.Cos(rad1)*math.Cos(rad2)*math.Pow(math.Sin(b/2), 2)))
	return s * EarthRadius
}

// Rad ...
func Rad(degree float64) float64 {
	return degree * math.Pi / 180.0
}
