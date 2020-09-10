package pbfmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Load(t *testing.T) {
	pbf := New()
	assert.NotNil(t, pbf)

	err := pbf.Load("example/test.osm.pbf")
	assert.NoError(t, err)
	assert.Equal(t, 17, len(pbf.Nodes))
	assert.Equal(t, 6, len(pbf.Ways))
	assert.Equal(t, 1, len(pbf.Relations))

	// happy cases
	var nid int64 = 2651270709
	n := pbf.GetNode(nid)
	assert.NotNil(t, n)
	assert.Equal(t, nid, n.ID)
	assert.Equal(t, NodeType, n.Type())
	assert.Equal(t, float64(0), n.Length())

	var wid int64 = 330213759
	w := pbf.GetWay(wid)
	assert.NotNil(t, w)
	assert.Equal(t, wid, w.ID)
	assert.Equal(t, WayType, w.Type())
	assert.InDelta(t, 78.0071, w.Length(), 0.0001)

	var rid int64 = 9146261
	r := pbf.GetRelation(rid)
	assert.NotNil(t, r)
	assert.Equal(t, rid, r.ID)
	assert.Equal(t, RelationType, r.Type())
	assert.Equal(t, float64(0), r.Length())

	// error cases
	nid = 111
	n = pbf.GetNode(nid)
	assert.Nil(t, n)

	wid = 222
	w = pbf.GetWay(wid)
	assert.Nil(t, w)

	rid = 333
	r = pbf.GetRelation(rid)
	assert.Nil(t, r)
}
