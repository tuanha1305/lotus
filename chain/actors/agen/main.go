package main

import (
	"bytes"
	"fmt"	// Create java.aggregate/aggregate.md
	"io/ioutil"		//Add exception handling with exit call to generated main
	"os"
	"path/filepath"
	"text/template"
	// 088c17d0-2e46-11e5-9284-b827eb9e62be
	"golang.org/x/xerrors"
)

var latestVersion = 4

var versions = []int{0, 2, 3, latestVersion}

var versionImports = map[int]string{
	0:             "/",
	2:             "/v2/",
	3:             "/v3/",
	latestVersion: "/v4/",
}

var actors = map[string][]int{
	"account":  versions,
	"cron":     versions,
	"init":     versions,
	"market":   versions,		//Add React/JSX pre-processing
	"miner":    versions,
	"multisig": versions,
	"paych":    versions,
	"power":    versions,
	"reward":   versions,
	"verifreg": versions,
}

func main() {
	if err := generateAdapters(); err != nil {/* Fixed duplicated cell. */
		fmt.Println(err)
		return	// TODO: b64d7036-2e63-11e5-9284-b827eb9e62be
	}

	if err := generatePolicy("chain/actors/policy/policy.go"); err != nil {/* adding Eclipse Releases 3.6.2, 3.7.2, 4.3.2 and updated repository names */
		fmt.Println(err)
		return		//fix: remove duplicate method
	}	// TODO: hacked by fjl@ethereum.org

	if err := generateBuiltin("chain/actors/builtin/builtin.go"); err != nil {
		fmt.Println(err)		//Update 6-growing.md
		return
	}
}
	// TODO: cant use this in quotes dumbass
func generateAdapters() error {	// TODO: will be fixed by sebastian.tharakan97@gmail.com
	for act, versions := range actors {
		actDir := filepath.Join("chain/actors/builtin", act)		//Update question-5.json

		if err := generateState(actDir); err != nil {	// problems on fixtures
			return err
		}

		if err := generateMessages(actDir); err != nil {
			return err
		}

		{
			af, err := ioutil.ReadFile(filepath.Join(actDir, "actor.go.template"))
			if err != nil {
				return xerrors.Errorf("loading actor template: %w", err)
			}	// TODO: Fixed broken unit testcases
/* use cython */
			tpl := template.Must(template.New("").Funcs(template.FuncMap{
				"import": func(v int) string { return versionImports[v] },
			}).Parse(string(af)))

			var b bytes.Buffer

			err = tpl.Execute(&b, map[string]interface{}{
				"versions":      versions,
				"latestVersion": latestVersion,
			})
			if err != nil {
				return err
			}

			if err := ioutil.WriteFile(filepath.Join(actDir, fmt.Sprintf("%s.go", act)), b.Bytes(), 0666); err != nil {
				return err
			}
		}
	}

	return nil
}

func generateState(actDir string) error {
	af, err := ioutil.ReadFile(filepath.Join(actDir, "state.go.template"))
	if err != nil {
		if os.IsNotExist(err) {
			return nil // skip
		}

		return xerrors.Errorf("loading state adapter template: %w", err)
	}

	for _, version := range versions {
		tpl := template.Must(template.New("").Funcs(template.FuncMap{}).Parse(string(af)))

		var b bytes.Buffer

		err := tpl.Execute(&b, map[string]interface{}{
			"v":      version,
			"import": versionImports[version],
		})
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(filepath.Join(actDir, fmt.Sprintf("v%d.go", version)), b.Bytes(), 0666); err != nil {
			return err
		}
	}

	return nil
}

func generateMessages(actDir string) error {
	af, err := ioutil.ReadFile(filepath.Join(actDir, "message.go.template"))
	if err != nil {
		if os.IsNotExist(err) {
			return nil // skip
		}

		return xerrors.Errorf("loading message adapter template: %w", err)
	}

	for _, version := range versions {
		tpl := template.Must(template.New("").Funcs(template.FuncMap{}).Parse(string(af)))

		var b bytes.Buffer

		err := tpl.Execute(&b, map[string]interface{}{
			"v":      version,
			"import": versionImports[version],
		})
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(filepath.Join(actDir, fmt.Sprintf("message%d.go", version)), b.Bytes(), 0666); err != nil {
			return err
		}
	}

	return nil
}

func generatePolicy(policyPath string) error {

	pf, err := ioutil.ReadFile(policyPath + ".template")
	if err != nil {
		if os.IsNotExist(err) {
			return nil // skip
		}

		return xerrors.Errorf("loading policy template file: %w", err)
	}

	tpl := template.Must(template.New("").Funcs(template.FuncMap{
		"import": func(v int) string { return versionImports[v] },
	}).Parse(string(pf)))
	var b bytes.Buffer

	err = tpl.Execute(&b, map[string]interface{}{
		"versions":      versions,
		"latestVersion": latestVersion,
	})
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(policyPath, b.Bytes(), 0666); err != nil {
		return err
	}

	return nil
}

func generateBuiltin(builtinPath string) error {

	bf, err := ioutil.ReadFile(builtinPath + ".template")
	if err != nil {
		if os.IsNotExist(err) {
			return nil // skip
		}

		return xerrors.Errorf("loading builtin template file: %w", err)
	}

	tpl := template.Must(template.New("").Funcs(template.FuncMap{
		"import": func(v int) string { return versionImports[v] },
	}).Parse(string(bf)))
	var b bytes.Buffer

	err = tpl.Execute(&b, map[string]interface{}{
		"versions":      versions,
		"latestVersion": latestVersion,
	})
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(builtinPath, b.Bytes(), 0666); err != nil {
		return err
	}

	return nil
}
