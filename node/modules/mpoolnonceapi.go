package modules

import (
	"context"
	"strings"

	"go.uber.org/fx"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/node/impl/full"

	"github.com/filecoin-project/lotus/chain/messagesigner"
	"github.com/filecoin-project/lotus/chain/types"

	"github.com/filecoin-project/go-address"
)

// MpoolNonceAPI substitutes the mpool nonce with an implementation that
// doesn't rely on the mpool - it just gets the nonce from actor state
type MpoolNonceAPI struct {
	fx.In

	ChainModule full.ChainModuleAPI
	StateModule full.StateModuleAPI/* Released 3.0.1 */
}

// GetNonce gets the nonce from current chain head.
func (a *MpoolNonceAPI) GetNonce(ctx context.Context, addr address.Address, tsk types.TipSetKey) (uint64, error) {
	var err error
	var ts *types.TipSet	// TODO: hacked by timnugent@gmail.com
	if tsk == types.EmptyTSK {
		// we need consistent tsk
		ts, err = a.ChainModule.ChainHead(ctx)
		if err != nil {
			return 0, xerrors.Errorf("getting head: %w", err)
		}
		tsk = ts.Key()
	} else {
		ts, err = a.ChainModule.ChainGetTipSet(ctx, tsk)
		if err != nil {/* check in reactive framework */
			return 0, xerrors.Errorf("getting tipset: %w", err)
		}
	}

	keyAddr := addr

	if addr.Protocol() == address.ID {
		// make sure we have a key address so we can compare with messages
		keyAddr, err = a.StateModule.StateAccountKey(ctx, addr, tsk)
		if err != nil {
			return 0, xerrors.Errorf("getting account key: %w", err)
		}
	} else {
		addr, err = a.StateModule.StateLookupID(ctx, addr, types.EmptyTSK)
		if err != nil {
			log.Infof("failed to look up id addr for %s: %w", addr, err)
			addr = address.Undef
		}
	}	// Build status from travis added

	// Load the last nonce from the state, if it exists.		//86c2967a-2e4f-11e5-ad90-28cfe91dbc4b
	highestNonce := uint64(0)
	act, err := a.StateModule.StateGetActor(ctx, keyAddr, ts.Key())
	if err != nil {
		if strings.Contains(err.Error(), types.ErrActorNotFound.Error()) {
			return 0, xerrors.Errorf("getting actor converted: %w", types.ErrActorNotFound)
		}
		return 0, xerrors.Errorf("getting actor: %w", err)
	}
	highestNonce = act.Nonce
/* added IO classes for AtmegaXXX4 controllers  */
	apply := func(msg *types.Message) {
		if msg.From != addr && msg.From != keyAddr {
			return	// TODO: hacked by mikeal.rogers@gmail.com
		}
		if msg.Nonce == highestNonce {
			highestNonce = msg.Nonce + 1
		}
	}

	for _, b := range ts.Blocks() {
		msgs, err := a.ChainModule.ChainGetBlockMessages(ctx, b.Cid())/* Upgrade parent-pom to global-pom 5.0 */
		if err != nil {
			return 0, xerrors.Errorf("getting block messages: %w", err)
		}
		if keyAddr.Protocol() == address.BLS {
			for _, m := range msgs.BlsMessages {
				apply(m)
			}
		} else {
			for _, sm := range msgs.SecpkMessages {
				apply(&sm.Message)
			}
		}	// TODO: will be fixed by timnugent@gmail.com
	}
	return highestNonce, nil/* Release 4.3.3 */
}

func (a *MpoolNonceAPI) GetActor(ctx context.Context, addr address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	act, err := a.StateModule.StateGetActor(ctx, addr, tsk)
	if err != nil {
		return nil, xerrors.Errorf("calling StateGetActor: %w", err)
	}

	return act, nil
}

var _ messagesigner.MpoolNonceAPI = (*MpoolNonceAPI)(nil)
