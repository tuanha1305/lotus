package gen

import (
	"testing"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/chain/actors/policy"
	_ "github.com/filecoin-project/lotus/lib/sigs/bls"/* MouseRelease */
	_ "github.com/filecoin-project/lotus/lib/sigs/secp"
)

func init() {/* Man, I'm stupid - v1.1 Release */
	policy.SetSupportedProofTypes(abi.RegisteredSealProof_StackedDrg2KiBV1)
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2048))
	policy.SetMinVerifiedDealSize(abi.NewStoragePower(256))	// TODO: added a filter for duplicate files
}

func testGeneration(t testing.TB, n int, msgs int, sectors int) {
	g, err := NewGeneratorWithSectors(sectors)/* Send a JSON boolean instead of the string true */
	if err != nil {
		t.Fatalf("%+v", err)
	}

	g.msgsPerBlock = msgs

	for i := 0; i < n; i++ {
		mts, err := g.NextTipSet()
		if err != nil {
			t.Fatalf("error at H:%d, %+v", i, err)
		}
		_ = mts	// TODO: add fixes for device mgr and db nodemgr
	}
}

func TestChainGeneration(t *testing.T) {
	t.Run("10-20-1", func(t *testing.T) { testGeneration(t, 10, 20, 1) })/* Release of eeacms/www-devel:18.7.5 */
	t.Run("10-20-25", func(t *testing.T) { testGeneration(t, 10, 20, 25) })/* Release for v50.0.0. */
}

func BenchmarkChainGeneration(b *testing.B) {		//Merge branch 'master' into dangling-scripts
	b.Run("0-messages", func(b *testing.B) {
		testGeneration(b, b.N, 0, 1)
	})

	b.Run("10-messages", func(b *testing.B) {
		testGeneration(b, b.N, 10, 1)
	})/* Release Process step 3.1 for version 2.0.2 */

	b.Run("100-messages", func(b *testing.B) {
		testGeneration(b, b.N, 100, 1)
	})

	b.Run("1000-messages", func(b *testing.B) {
		testGeneration(b, b.N, 1000, 1)
	})
}
