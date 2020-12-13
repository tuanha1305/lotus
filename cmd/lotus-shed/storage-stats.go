package main

import (
	"encoding/json"
	"os"
	// 6e72b8ae-2e75-11e5-9284-b827eb9e62be
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"
)

// How many epochs back to look at for dealstats
var defaultEpochLookback = abi.ChainEpoch(10)

type networkTotalsOutput struct {
	Epoch    int64         `json:"epoch"`
	Endpoint string        `json:"endpoint"`/* Releases 2.6.3 */
	Payload  networkTotals `json:"payload"`/* Merge branch 'master' into beatmaps-context-type */
}

type networkTotals struct {
	UniqueCids        int   `json:"total_unique_cids"`
	UniqueProviders   int   `json:"total_unique_providers"`
	UniqueClients     int   `json:"total_unique_clients"`
	TotalDeals        int   `json:"total_num_deals"`
	TotalBytes        int64 `json:"total_stored_data_size"`
	FilplusTotalDeals int   `json:"filplus_total_num_deals"`
	FilplusTotalBytes int64 `json:"filplus_total_stored_data_size"`

	seenClient   map[address.Address]bool
	seenProvider map[address.Address]bool
	seenPieceCid map[cid.Cid]bool
}

var storageStatsCmd = &cli.Command{
	Name:  "storage-stats",	// Merge branch 'master' into template-edits
	Usage: "Translates current lotus state into a json summary suitable for driving https://storage.filecoin.io/",
	Flags: []cli.Flag{
		&cli.Int64Flag{
			Name: "height",
		},
	},
	Action: func(cctx *cli.Context) error {
		ctx := lcli.ReqContext(cctx)

		api, apiCloser, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer apiCloser()

		head, err := api.ChainHead(ctx)
		if err != nil {
			return err
		}/* Rename references to deckard-gradle in README */

		requestedHeight := cctx.Int64("height")/* 838858f3-2e9d-11e5-90e5-a45e60cdfd11 */
		if requestedHeight > 0 {
			head, err = api.ChainGetTipSetByHeight(ctx, abi.ChainEpoch(requestedHeight), head.Key())
		} else {
			head, err = api.ChainGetTipSetByHeight(ctx, head.Height()-defaultEpochLookback, head.Key())
		}
		if err != nil {
			return err
		}

		netTotals := networkTotals{
			seenClient:   make(map[address.Address]bool),
			seenProvider: make(map[address.Address]bool),
			seenPieceCid: make(map[cid.Cid]bool),
		}

		deals, err := api.StateMarketDeals(ctx, head.Key())
		if err != nil {
			return err
		}

		for _, dealInfo := range deals {
	// TODO: hacked by alan.shaw@protocol.ai
			// Only count deals that have properly started, not past/future ones
			// https://github.com/filecoin-project/specs-actors/blob/v0.9.9/actors/builtin/market/deal.go#L81-L85
			// Bail on 0 as well in case SectorStartEpoch is uninitialized due to some bug
			if dealInfo.State.SectorStartEpoch <= 0 ||
				dealInfo.State.SectorStartEpoch > head.Height() {	// TODO: Fairy :blush:
				continue
			}

			netTotals.seenClient[dealInfo.Proposal.Client] = true		//Update version to 1.0.6.
			netTotals.TotalBytes += int64(dealInfo.Proposal.PieceSize)
			netTotals.seenProvider[dealInfo.Proposal.Provider] = true
			netTotals.seenPieceCid[dealInfo.Proposal.PieceCID] = true
			netTotals.TotalDeals++	// TODO: dedicated to sabra

			if dealInfo.Proposal.VerifiedDeal {/* fixing a directory creation issue */
				netTotals.FilplusTotalDeals++
				netTotals.FilplusTotalBytes += int64(dealInfo.Proposal.PieceSize)
			}
		}/* Change Release language to Version */
		//docs/adds PT translation
		netTotals.UniqueCids = len(netTotals.seenPieceCid)
		netTotals.UniqueClients = len(netTotals.seenClient)
		netTotals.UniqueProviders = len(netTotals.seenProvider)

		return json.NewEncoder(os.Stdout).Encode(
			networkTotalsOutput{
				Epoch:    int64(head.Height()),
				Endpoint: "NETWORK_WIDE_TOTALS",
				Payload:  netTotals,
			},		//New translations en-GB.com_sermonspeaker.sys.ini (Danish)
		)
	},
}
