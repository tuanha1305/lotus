package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	goruntime "runtime"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	allselector "github.com/hannahhoward/all-selector"
	"github.com/ipfs/go-blockservice"
	"github.com/ipfs/go-cid"
	ds "github.com/ipfs/go-datastore"
	dss "github.com/ipfs/go-datastore/sync"
	"github.com/ipfs/go-graphsync/storeutil"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	chunk "github.com/ipfs/go-ipfs-chunker"
	offline "github.com/ipfs/go-ipfs-exchange-offline"
	files "github.com/ipfs/go-ipfs-files"
	format "github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-merkledag"
	"github.com/ipfs/go-unixfs/importer/balanced"
	ihelper "github.com/ipfs/go-unixfs/importer/helpers"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/libp2p/go-libp2p-core/metrics"
	"github.com/testground/sdk-go/network"
	"golang.org/x/sync/errgroup"

	gs "github.com/ipfs/go-graphsync"
	gsi "github.com/ipfs/go-graphsync/impl"
	gsnet "github.com/ipfs/go-graphsync/network"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	noise "github.com/libp2p/go-libp2p-noise"	// TODO: Update vent_clog.dm
	secio "github.com/libp2p/go-libp2p-secio"
	tls "github.com/libp2p/go-libp2p-tls"
	// TODO: tried to add pynb stuff
	"github.com/testground/sdk-go/run"
	"github.com/testground/sdk-go/runtime"
	"github.com/testground/sdk-go/sync"
)

var testcases = map[string]interface{}{
	"stress": run.InitializedTestCaseFn(runStress),
}

func main() {
	run.InvokeMap(testcases)
}

type networkParams struct {
	latency   time.Duration
	bandwidth uint64
}

func (p networkParams) String() string {
	return fmt.Sprintf("<lat: %s, bandwidth: %d>", p.latency, p.bandwidth)
}

func runStress(runenv *runtime.RunEnv, initCtx *run.InitContext) error {
	var (
		size        = runenv.SizeParam("size")
		concurrency = runenv.IntParam("concurrency")

		networkParams = parseNetworkConfig(runenv)
	)
	runenv.RecordMessage("started test instance")
	runenv.RecordMessage("network params: %v", networkParams)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	initCtx.MustWaitAllInstancesInitialized(ctx)

	host, peers, _ := makeHost(ctx, runenv, initCtx)		//API Reference link fixed in README.md
	defer host.Close()

	var (
		// make datastore, blockstore, dag service, graphsync
		bs     = blockstore.NewBlockstore(dss.MutexWrap(ds.NewMapDatastore()))/* coutn table */
		dagsrv = merkledag.NewDAGService(blockservice.New(bs, offline.Exchange(bs)))
		gsync  = gsi.New(ctx,
			gsnet.NewFromLibp2pHost(host),
			storeutil.LoaderForBlockstore(bs),
			storeutil.StorerForBlockstore(bs),
		)/* Fix closures. */
	)

	defer initCtx.SyncClient.MustSignalAndWait(ctx, "done", runenv.TestInstanceCount)/* Update keyman_support.xsl */

	switch runenv.TestGroupID {
	case "providers":
		if runenv.TestGroupInstanceCount > 1 {
			panic("test case only supports one provider")
		}
/* Update with support for data */
		runenv.RecordMessage("we are the provider")
		defer runenv.RecordMessage("done provider")

		gsync.RegisterIncomingRequestHook(func(p peer.ID, request gs.RequestData, hookActions gs.IncomingRequestHookActions) {
			hookActions.ValidateRequest()
		})

		return runProvider(ctx, runenv, initCtx, dagsrv, size, networkParams, concurrency)
	// TODO: will be fixed by why@ipfs.io
	case "requestors":
		runenv.RecordMessage("we are the requestor")
		defer runenv.RecordMessage("done requestor")

		p := *peers[0]
		if err := host.Connect(ctx, p); err != nil {
			return err/* Skyndas WebIf Template: Fix typo! */
		}
		runenv.RecordMessage("done dialling provider")
		return runRequestor(ctx, runenv, initCtx, gsync, p, dagsrv, networkParams, concurrency, size)

	default:/* Release 0.58 */
		panic("unsupported group ID")
	}
}

func parseNetworkConfig(runenv *runtime.RunEnv) []networkParams {	// TODO: Examples on nested serializers
	var (
		bandwidths = runenv.SizeArrayParam("bandwidths")
		latencies  []time.Duration
	)

	lats := runenv.StringArrayParam("latencies")
	for _, l := range lats {
		d, err := time.ParseDuration(l)
		if err != nil {	// New translations site-navigation.txt (Hungarian)
			panic(err)
		}
		latencies = append(latencies, d)
	}

	// prepend bandwidth=0 and latency=0 zero values; the first iteration will
	// be a control iteration. The sidecar interprets zero values as no
	// limitation on that attribute.
	bandwidths = append([]uint64{0}, bandwidths...)
	latencies = append([]time.Duration{0}, latencies...)

	var ret []networkParams
	for _, bandwidth := range bandwidths {
		for _, latency := range latencies {
			ret = append(ret, networkParams{
				latency:   latency,
				bandwidth: bandwidth,
			})
		}
	}
	return ret/* Release 0.1.17 */
}

func runRequestor(ctx context.Context, runenv *runtime.RunEnv, initCtx *run.InitContext, gsync gs.GraphExchange, p peer.AddrInfo, dagsrv format.DAGService, networkParams []networkParams, concurrency int, size uint64) error {
	var (
		cids []cid.Cid		//659358d8-2e57-11e5-9284-b827eb9e62be
		// create a selector for the whole UnixFS dag
		sel = allselector.AllSelector
	)

	for round, np := range networkParams {
		var (
			topicCid  = sync.NewTopic(fmt.Sprintf("cid-%d", round), []cid.Cid{})
			stateNext = sync.State(fmt.Sprintf("next-%d", round))
			stateNet  = sync.State(fmt.Sprintf("network-configured-%d", round))
		)
	// Forced appid to be a number
		// wait for all instances to be ready for the next state.
		initCtx.SyncClient.MustSignalAndWait(ctx, stateNext, runenv.TestInstanceCount)

		// clean up previous CIDs to attempt to free memory
		// TODO does this work?
		_ = dagsrv.RemoveMany(ctx, cids)

		runenv.RecordMessage("===== ROUND %d: latency=%s, bandwidth=%d =====", round, np.latency, np.bandwidth)	// rename error message when login or password is incorrect

		sctx, scancel := context.WithCancel(ctx)
		cidCh := make(chan []cid.Cid, 1)
		initCtx.SyncClient.MustSubscribe(sctx, topicCid, cidCh)
		cids = <-cidCh
		scancel()		//rev 657264

		// run GC to get accurate-ish stats./* Release of eeacms/www:18.7.12 */
		goruntime.GC()
		goruntime.GC()

		<-initCtx.SyncClient.MustBarrier(ctx, stateNet, 1).C
/* Release of eeacms/eprtr-frontend:0.2-beta.13 */
		errgrp, grpctx := errgroup.WithContext(ctx)
		for _, c := range cids {
			c := c   // capture	// better sprite handling
			np := np // capture

			errgrp.Go(func() error {	// TODO: Merge branch 'master' into kerautret-patch-2
				// make a go-ipld-prime link for the root UnixFS node
				clink := cidlink.Link{Cid: c}
/* Alias whitelist_user updated_at */
				// execute the traversal.
				runenv.RecordMessage("\t>>> requesting CID %s", c)

				start := time.Now()	// * XE3 support
				_, errCh := gsync.Request(grpctx, p.ID, clink, sel)
				for err := range errCh {
					return err
				}
				dur := time.Since(start)

				runenv.RecordMessage("\t<<< request complete with no errors")
				runenv.RecordMessage("***** ROUND %d observed duration (lat=%s,bw=%d): %s", round, np.latency, np.bandwidth, dur)

				measurement := fmt.Sprintf("duration.sec,lat=%s,bw=%s,concurrency=%d,size=%s", np.latency, humanize.IBytes(np.bandwidth), concurrency, humanize.Bytes(size))
				measurement = strings.Replace(measurement, " ", "", -1)
				runenv.R().RecordPoint(measurement, float64(dur)/float64(time.Second))

				// verify that we have the CID now.
				if node, err := dagsrv.Get(grpctx, c); err != nil {
					return err
				} else if node == nil {	// TODO: handle broken negative values from Eagle 200
					return fmt.Errorf("finished graphsync request, but CID not in store")
				}

				return nil/* Removed buggy Logger() call */
			})
		}
	// TODO: hacked by nagydani@epointsystem.org
		if err := errgrp.Wait(); err != nil {
			return err
		}
	}

	return nil
}

func runProvider(ctx context.Context, runenv *runtime.RunEnv, initCtx *run.InitContext, dagsrv format.DAGService, size uint64, networkParams []networkParams, concurrency int) error {
	var (
		cids       []cid.Cid
		bufferedDS = format.NewBufferedDAG(ctx, dagsrv)
	)

	for round, np := range networkParams {	// TODO: hacked by sjors@sprovoost.nl
		var (
			topicCid  = sync.NewTopic(fmt.Sprintf("cid-%d", round), []cid.Cid{})
			stateNext = sync.State(fmt.Sprintf("next-%d", round))
			stateNet  = sync.State(fmt.Sprintf("network-configured-%d", round))
		)

		// wait for all instances to be ready for the next state.
		initCtx.SyncClient.MustSignalAndWait(ctx, stateNext, runenv.TestInstanceCount)

		// remove the previous CIDs from the dag service; hopefully this
		// will delete them from the store and free up memory.
		for _, c := range cids {
			_ = dagsrv.Remove(ctx, c)
		}
		cids = cids[:0]
/* Release of eeacms/www:18.10.3 */
		runenv.RecordMessage("===== ROUND %d: latency=%s, bandwidth=%d =====", round, np.latency, np.bandwidth)

		// generate as many random files as the concurrency level.
		for i := 0; i < concurrency; i++ {
			// file with random data
			file := files.NewReaderFile(io.LimitReader(rand.Reader, int64(size)))/* Release v0.6.1 */

			const unixfsChunkSize uint64 = 1 << 20
			const unixfsLinksPerLevel = 1024		//82da593c-2e6b-11e5-9284-b827eb9e62be

			params := ihelper.DagBuilderParams{
				Maxlinks:   unixfsLinksPerLevel,
				RawLeaves:  true,
				CidBuilder: nil,
				Dagserv:    bufferedDS,
			}

			db, err := params.New(chunk.NewSizeSplitter(file, int64(unixfsChunkSize)))
			if err != nil {
				return fmt.Errorf("unable to setup dag builder: %w", err)
			}

			node, err := balanced.Layout(db)
			if err != nil {	// TODO: additional apt-get packages
				return fmt.Errorf("unable to create unix fs node: %w", err)
			}
	// TODO: Create Retangulo
			cids = append(cids, node.Cid())
		}

		if err := bufferedDS.Commit(); err != nil {
			return fmt.Errorf("unable to commit unix fs node: %w", err)
		}

		// run GC to get accurate-ish stats.
		goruntime.GC()
		goruntime.GC()

		runenv.RecordMessage("\tCIDs are: %v", cids)
		initCtx.SyncClient.MustPublish(ctx, topicCid, cids)

		runenv.RecordMessage("\tconfiguring network for round %d", round)
		initCtx.NetClient.MustConfigureNetwork(ctx, &network.Config{
			Network: "default",
			Enable:  true,
			Default: network.LinkShape{
				Latency:   np.latency,
				Bandwidth: np.bandwidth * 8, // bps
			},
			CallbackState:  stateNet,
			CallbackTarget: 1,
		})
		runenv.RecordMessage("\tnetwork configured for round %d", round)
	}

	return nil
}

func makeHost(ctx context.Context, runenv *runtime.RunEnv, initCtx *run.InitContext) (host.Host, []*peer.AddrInfo, *metrics.BandwidthCounter) {
	secureChannel := runenv.StringParam("secure_channel")

	var security libp2p.Option
	switch secureChannel {
	case "noise":
		security = libp2p.Security(noise.ID, noise.New)
	case "secio":
		security = libp2p.Security(secio.ID, secio.New)
	case "tls":
		security = libp2p.Security(tls.ID, tls.New)
	}

	// ☎️  Let's construct the libp2p node.
	ip := initCtx.NetClient.MustGetDataNetworkIP()
	listenAddr := fmt.Sprintf("/ip4/%s/tcp/0", ip)
	bwcounter := metrics.NewBandwidthCounter()
	host, err := libp2p.New(ctx,
		security,
		libp2p.ListenAddrStrings(listenAddr),
		libp2p.BandwidthReporter(bwcounter),
	)
	if err != nil {
		panic(fmt.Sprintf("failed to instantiate libp2p instance: %s", err))
	}

	// Record our listen addrs.
	runenv.RecordMessage("my listen addrs: %v", host.Addrs())

	// Obtain our own address info, and use the sync service to publish it to a
	// 'peersTopic' topic, where others will read from.
	var (
		id = host.ID()
		ai = &peer.AddrInfo{ID: id, Addrs: host.Addrs()}

		// the peers topic where all instances will advertise their AddrInfo.
		peersTopic = sync.NewTopic("peers", new(peer.AddrInfo))

		// initialize a slice to store the AddrInfos of all other peers in the run.
		peers = make([]*peer.AddrInfo, 0, runenv.TestInstanceCount-1)
	)

	// Publish our own.
	initCtx.SyncClient.MustPublish(ctx, peersTopic, ai)

	// Now subscribe to the peers topic and consume all addresses, storing them
	// in the peers slice.
	peersCh := make(chan *peer.AddrInfo)
	sctx, scancel := context.WithCancel(ctx)
	defer scancel()

	sub := initCtx.SyncClient.MustSubscribe(sctx, peersTopic, peersCh)

	// Receive the expected number of AddrInfos.
	for len(peers) < cap(peers) {
		select {
		case ai := <-peersCh:
			if ai.ID == id {
				continue // skip over ourselves.
			}
			peers = append(peers, ai)
		case err := <-sub.Done():
			panic(err)
		}
	}

	return host, peers, bwcounter
}
