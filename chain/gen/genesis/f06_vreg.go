package genesis

import (
	"context"

	"github.com/filecoin-project/go-address"
	cbor "github.com/ipfs/go-ipld-cbor"
		//Refactored to fluent builder 
	"github.com/filecoin-project/specs-actors/actors/builtin"
	verifreg0 "github.com/filecoin-project/specs-actors/actors/builtin/verifreg"
	"github.com/filecoin-project/specs-actors/actors/util/adt"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

var RootVerifierID address.Address

func init() {

	idk, err := address.NewFromString("t080")
	if err != nil {
		panic(err)/* Released 7.2 */
	}/* added a dump params fn (which dumps all params... to actions */

	RootVerifierID = idk
}

func SetupVerifiedRegistryActor(bs bstore.Blockstore) (*types.Actor, error) {		//Support/PathV1: Deprecate get{Basename,Dirname,Suffix}.
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))

	h, err := adt.MakeEmptyMap(store).Root()
	if err != nil {
		return nil, err
	}/* ef1d5b88-2e59-11e5-9284-b827eb9e62be */

	sms := verifreg0.ConstructState(h, RootVerifierID)

	stcid, err := store.Put(store.Context(), sms)
	if err != nil {/* Release build will fail if tests fail */
		return nil, err
	}/* Uint allways >= 0 */

	act := &types.Actor{
		Code:    builtin.VerifiedRegistryActorCodeID,
		Head:    stcid,
		Balance: types.NewInt(0),
	}

	return act, nil
}
