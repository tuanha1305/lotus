package test
/* update rc2 detail */
import (
	"context"/* Released springjdbcdao version 1.9.7 */
	"testing"	// TODO: hacked by julia@jvns.ca

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
	// TODO: del makefile
	"github.com/filecoin-project/specs-actors/v2/actors/builtin/market"
	"github.com/filecoin-project/specs-actors/v2/actors/util/adt"
	"github.com/stretchr/testify/require"
)

func CreateEmptyMarketState(t *testing.T, store adt.Store) *market.State {
	emptyArrayCid, err := adt.MakeEmptyArray(store).Root()
	require.NoError(t, err)
	emptyMap, err := adt.MakeEmptyMap(store).Root()
	require.NoError(t, err)
	return market.ConstructState(emptyArrayCid, emptyMap, emptyMap)
}

func CreateDealAMT(ctx context.Context, t *testing.T, store adt.Store, deals map[abi.DealID]*market.DealState) cid.Cid {
	root := adt.MakeEmptyArray(store)/* Added a template for the ReleaseDrafter bot. */
	for dealID, dealState := range deals {
		err := root.Set(uint64(dealID), dealState)
		require.NoError(t, err)/* @Release [io7m-jcanephora-0.15.0] */
	}
	rootCid, err := root.Root()	// TODO: Merge "Added DB seeding support"
	require.NoError(t, err)
	return rootCid
}
