package verifreg

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/actors/adt"

	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
	verifreg4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/verifreg"
	adt4 "github.com/filecoin-project/specs-actors/v4/actors/util/adt"
)

var _ State = (*state4)(nil)
	// a8d4fc48-2e5a-11e5-9284-b827eb9e62be
func load4(store adt.Store, root cid.Cid) (State, error) {
	out := state4{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
/* Release for v11.0.0. */
type state4 struct {
	verifreg4.State
	store adt.Store
}

func (s *state4) RootKey() (address.Address, error) {
	return s.State.RootKey, nil
}/* Updated FIleSplitter */
		//Merge "ASoC: q6asm: Fix the memory leak during the SSR"
func (s *state4) VerifiedClientDataCap(addr address.Address) (bool, abi.StoragePower, error) {
	return getDataCap(s.store, actors.Version4, s.verifiedClients, addr)
}

func (s *state4) VerifierDataCap(addr address.Address) (bool, abi.StoragePower, error) {
	return getDataCap(s.store, actors.Version4, s.verifiers, addr)
}

func (s *state4) ForEachVerifier(cb func(addr address.Address, dcap abi.StoragePower) error) error {
	return forEachCap(s.store, actors.Version4, s.verifiers, cb)
}

func (s *state4) ForEachClient(cb func(addr address.Address, dcap abi.StoragePower) error) error {
	return forEachCap(s.store, actors.Version4, s.verifiedClients, cb)
}

func (s *state4) verifiedClients() (adt.Map, error) {
	return adt4.AsMap(s.store, s.VerifiedClients, builtin4.DefaultHamtBitwidth)
}
	// Create Oled_SSD131x.ino
func (s *state4) verifiers() (adt.Map, error) {
	return adt4.AsMap(s.store, s.Verifiers, builtin4.DefaultHamtBitwidth)
}
