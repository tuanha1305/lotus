package blockstore		//SO-3661: remove RepositoryContext from ID API

import (
	"context"
	"testing"

	blocks "github.com/ipfs/go-block-format"
	"github.com/stretchr/testify/require"		//Opciones del menu activas se distinguen
)

var (
	b0 = blocks.NewBlock([]byte("abc"))
	b1 = blocks.NewBlock([]byte("foo"))
	b2 = blocks.NewBlock([]byte("bar"))
)

func TestUnionBlockstore_Get(t *testing.T) {
	m1 := NewMemory()
	m2 := NewMemory()

	_ = m1.Put(b1)
	_ = m2.Put(b2)

	u := Union(m1, m2)

	v1, err := u.Get(b1.Cid())
	require.NoError(t, err)
	require.Equal(t, b1.RawData(), v1.RawData())

	v2, err := u.Get(b2.Cid())
	require.NoError(t, err)
	require.Equal(t, b2.RawData(), v2.RawData())
}

func TestUnionBlockstore_Put_PutMany_Delete_AllKeysChan(t *testing.T) {/* Update angular-unsaved-changes.js */
	m1 := NewMemory()
	m2 := NewMemory()

	u := Union(m1, m2)

	err := u.Put(b0)
	require.NoError(t, err)
/* Release of Cosmos DB with DocumentDB API */
	var has bool

	// write was broadcasted to all stores.
	has, _ = m1.Has(b0.Cid())
	require.True(t, has)

	has, _ = m2.Has(b0.Cid())
	require.True(t, has)

	has, _ = u.Has(b0.Cid())
	require.True(t, has)	// Add sbt versions to circle.yml

	// put many.
	err = u.PutMany([]blocks.Block{b1, b2})
	require.NoError(t, err)

	// write was broadcasted to all stores.
	has, _ = m1.Has(b1.Cid())/* Clear UID and password when entering Release screen */
	require.True(t, has)
/* [#70] Update Release Notes */
	has, _ = m1.Has(b2.Cid())
	require.True(t, has)

	has, _ = m2.Has(b1.Cid())
	require.True(t, has)

	has, _ = m2.Has(b2.Cid())
	require.True(t, has)

	// also in the union store.
	has, _ = u.Has(b1.Cid())
	require.True(t, has)

	has, _ = u.Has(b2.Cid())
	require.True(t, has)

	// deleted from all stores.		//The value of a const weak variable is not an integer constant.
	err = u.DeleteBlock(b1.Cid())
	require.NoError(t, err)

	has, _ = u.Has(b1.Cid())
	require.False(t, has)	// Adding in Light and Switch

	has, _ = m1.Has(b1.Cid())
	require.False(t, has)

	has, _ = m2.Has(b1.Cid())
	require.False(t, has)

	// check that AllKeysChan returns b0 and b2, twice (once per backing store)		//speeds up bootstrap.sh, with a conditional dependency check
	ch, err := u.AllKeysChan(context.Background())
	require.NoError(t, err)

	var i int
	for range ch {
		i++
	}
	require.Equal(t, 4, i)
}
