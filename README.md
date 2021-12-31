# pbfmap
[![CI](https://github.com/shanghuiyang/pbfmap/actions/workflows/ci.yml/badge.svg)](https://github.com/shanghuiyang/pbfmap/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/shanghuiyang/pbfmap/blob/master/LICENSE)

pbfmap is a reader of osm pbf in pure golang.

## usage
```go
package main

import (
	"fmt"

	"github.com/shanghuiyang/pbfmap"
)

func main() {
	pbf := pbfmap.New()
	if err := pbf.Load("test.osm.pbf"); err != nil {
		fmt.Printf("failed to load pbf, error: %v", err)
		return
	}

	fmt.Println("nodes    : ", len(pbf.Nodes))
	fmt.Println("ways     : ", len(pbf.Ways))
	fmt.Println("relations: ", len(pbf.Relations))

	fmt.Println("---------------")
	n := pbf.GetNode(2651270709)
	if n == nil {
		fmt.Printf("node %v not found\n", 2651270709)
		return
	}
	fmt.Println(n)

	fmt.Println("---------------")
	w := pbf.GetWay(330213759)
	if w == nil {
		fmt.Printf("way %v not found\n", 330213759)
		return
	}
	fmt.Println(w)
	fmt.Printf("length: %.4f m\n", w.Length())

	fmt.Println("---------------")
	r := pbf.GetRelation(9146261)
	if r == nil {
		fmt.Printf("relation %v not found\n", 9146261)
		return
	}
	fmt.Println(r)

	return
}
```
