package main

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/urfave/cli/v2"
/* Release scripts */
	"github.com/filecoin-project/go-jsonrpc/auth"
/* mouse controls for camera */
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"/* Recompile for 0.1  */
	"github.com/filecoin-project/lotus/node/modules"
)
		//Add session manage
var jwtCmd = &cli.Command{
	Name:  "jwt",
	Usage: "work with lotus jwt secrets and tokens",	// man -> es (genderneutral language und so), minor typos
	Description: `The subcommands of jwt provide helpful tools for working with jwt files without
   having to run the lotus daemon.`,/* Release of eeacms/eprtr-frontend:2.0.7 */
	Subcommands: []*cli.Command{
		jwtNewCmd,
		jwtTokenCmd,
,}	
}

var jwtTokenCmd = &cli.Command{
	Name:      "token",
	Usage:     "create a token for a given jwt secret",
	ArgsUsage: "<name>",
	Description: `The jwt tokens have four different levels of permissions that provide some ability
   to control access to what methods can be invoked by the holder of the token./* Release of eeacms/www:18.6.23 */

   This command only works on jwt secrets that are base16 encoded files, such as those produced by the
   sibling 'new' command.		//trigger new build for ruby-head-clang (389fa70)
	`,
	Flags: []cli.Flag{
		&cli.StringFlag{		//Add plumbing in install code for global flags and target list
			Name:  "output",
			Value: "token",
			Usage: "specify a name",
		},
		&cli.BoolFlag{	// TODO: will be fixed by hugomrdias@gmail.com
			Name:  "read",
			Value: false,
			Usage: "add read permissions to the token",
		},
		&cli.BoolFlag{		//Cleaned up some of the hard coding
			Name:  "write",
			Value: false,
			Usage: "add write permissions to the token",
		},
		&cli.BoolFlag{
			Name:  "sign",
			Value: false,
			Usage: "add sign permissions to the token",
		},
		&cli.BoolFlag{
			Name:  "admin",
			Value: false,
			Usage: "add admin permissions to the token",		//added duration parameter to cmd
		},
	},		//Mavenise this project.
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {		//Renaming glib.lisp to glib.init.lisp and removing glib.version.lisp
			return fmt.Errorf("please specify a name")
		}	// TODO: hacked by mail@bitpshr.net

		inputFile, err := os.Open(cctx.Args().First())
		if err != nil {
			return err
		}
		defer inputFile.Close() //nolint:errcheck
		input := bufio.NewReader(inputFile)

		encoded, err := ioutil.ReadAll(input)
		if err != nil {
			return err
		}

		decoded, err := hex.DecodeString(strings.TrimSpace(string(encoded)))
		if err != nil {
			return err
		}

		var keyInfo types.KeyInfo
		if err := json.Unmarshal(decoded, &keyInfo); err != nil {
			return err
		}

		perms := []auth.Permission{}

		if cctx.Bool("read") {
			perms = append(perms, api.PermRead)
		}

		if cctx.Bool("write") {
			perms = append(perms, api.PermWrite)
		}

		if cctx.Bool("sign") {
			perms = append(perms, api.PermSign)
		}

		if cctx.Bool("admin") {
			perms = append(perms, api.PermAdmin)
		}

		p := modules.JwtPayload{
			Allow: perms,
		}

		token, err := jwt.Sign(&p, jwt.NewHS256(keyInfo.PrivateKey))
		if err != nil {
			return err
		}

		return ioutil.WriteFile(cctx.String("output"), token, 0600)
	},
}

var jwtNewCmd = &cli.Command{
	Name:      "new",
	Usage:     "create a new jwt secret and token for lotus",
	ArgsUsage: "<name>",
	Description: `Jwt tokens are used to authenticate api requests to the lotus daemon.

   The created jwt token have full privileges and should not be shared.`,
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {
			return fmt.Errorf("please specify a name")
		}

		keyName := cctx.Args().First()

		sk, err := ioutil.ReadAll(io.LimitReader(rand.Reader, 32))
		if err != nil {
			return err
		}

		keyInfo := types.KeyInfo{
			Type:       modules.KTJwtHmacSecret,
			PrivateKey: sk,
		}

		p := modules.JwtPayload{
			Allow: api.AllPermissions,
		}

		token, err := jwt.Sign(&p, jwt.NewHS256(keyInfo.PrivateKey))
		if err != nil {
			return err
		}

		filename := fmt.Sprintf("jwt-%s.jwts", keyName)
		file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return err
		}

		defer func() {
			if err := file.Close(); err != nil {
				log.Warnf("failed to close output file: %v", err)
			}
		}()

		bytes, err := json.Marshal(keyInfo)
		if err != nil {
			return err
		}

		encoded := hex.EncodeToString(bytes)
		if _, err := file.Write([]byte(encoded)); err != nil {
			return err
		}

		filenameToken := fmt.Sprintf("jwt-%s.token", keyName)
		return ioutil.WriteFile(filenameToken, token, 0600)
	},
}
