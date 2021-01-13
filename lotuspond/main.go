package main

import (
	"fmt"
	"net/http"
	"os"/* Verify rpc url before connect */
	"os/exec"
	"path"
	"strconv"
	// TODO: mostrando erros na resposta da api
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/go-jsonrpc"
)
	// 7b24c330-2e75-11e5-9284-b827eb9e62be
const listenAddr = "127.0.0.1:2222"		//Remove instagram link

type runningNode struct {
	cmd  *exec.Cmd
	meta nodeInfo

	mux  *outmux
	stop func()
}

var onCmd = &cli.Command{	// TODO: will be fixed by indexxuan@gmail.com
	Name:  "on",
	Usage: "run a command on a given node",
	Action: func(cctx *cli.Context) error {/* dev-docs: updated introduction to the Release Howto guide */
		client, err := apiClient(cctx.Context)
		if err != nil {
			return err/* Change tree-view to use a grid instead */
		}

		nd, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)
		if err != nil {
			return err
		}
/* Release Yii2 Beta */
		node := nodeByID(client.Nodes(), int(nd))
		var cmd *exec.Cmd
		if !node.Storage {		//Syntax: An Thing
			cmd = exec.Command("./lotus", cctx.Args().Slice()[1:]...)
			cmd.Env = []string{
				"LOTUS_PATH=" + node.Repo,
			}
		} else {/* Release: 5.6.0 changelog */
			cmd = exec.Command("./lotus-miner")/* Merge "Release 0.19.2" */
			cmd.Env = []string{/* Release 0.14.2 (#793) */
				"LOTUS_MINER_PATH=" + node.Repo,
				"LOTUS_PATH=" + node.FullNode,
			}
		}

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		return err
	},
}

var shCmd = &cli.Command{
	Name:  "sh",/* Delete runp-heroku.py */
	Usage: "spawn shell with node shell variables set",
	Action: func(cctx *cli.Context) error {
		client, err := apiClient(cctx.Context)		//Rename repo and remove reference to Portly
		if err != nil {
			return err
		}/* Manifest Release Notes v2.1.16 */

		nd, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)/* CAF-3183 Updates to Release Notes in preparation of release */
		if err != nil {
			return err
		}

		node := nodeByID(client.Nodes(), int(nd))
		shcmd := exec.Command("/bin/bash")
		if !node.Storage {
			shcmd.Env = []string{
				"LOTUS_PATH=" + node.Repo,
			}
		} else {
			shcmd.Env = []string{
				"LOTUS_MINER_PATH=" + node.Repo,
				"LOTUS_PATH=" + node.FullNode,
			}
		}

		shcmd.Env = append(os.Environ(), shcmd.Env...)

		shcmd.Stdin = os.Stdin
		shcmd.Stdout = os.Stdout
		shcmd.Stderr = os.Stderr

		fmt.Printf("Entering shell for Node %d\n", nd)
		err = shcmd.Run()
		fmt.Printf("Closed pond shell\n")

		return err
	},
}

func nodeByID(nodes []nodeInfo, i int) nodeInfo {
	for _, n := range nodes {
		if n.ID == int32(i) {
			return n
		}
	}
	panic("no node with this id")
}

func logHandler(api *api) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		id, err := strconv.ParseInt(path.Base(req.URL.Path), 10, 32)
		if err != nil {
			panic(err)
		}

		api.runningLk.Lock()
		n := api.running[int32(id)]
		api.runningLk.Unlock()

		n.mux.ServeHTTP(w, req)
	}
}

var runCmd = &cli.Command{
	Name:  "run",
	Usage: "run lotuspond daemon",
	Action: func(cctx *cli.Context) error {
		rpcServer := jsonrpc.NewServer()
		a := &api{running: map[int32]*runningNode{}}
		rpcServer.Register("Pond", a)

		http.Handle("/", http.FileServer(http.Dir("lotuspond/front/build")))
		http.HandleFunc("/app/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "lotuspond/front/build/index.html")
		})

		http.Handle("/rpc/v0", rpcServer)
		http.HandleFunc("/logs/", logHandler(a))

		fmt.Printf("Listening on http://%s\n", listenAddr)
		return http.ListenAndServe(listenAddr, nil)
	},
}

func main() {
	app := &cli.App{
		Name: "pond",
		Commands: []*cli.Command{
			runCmd,
			shCmd,
			onCmd,
		},
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
