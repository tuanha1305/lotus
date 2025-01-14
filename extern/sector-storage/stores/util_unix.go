package stores

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"golang.org/x/xerrors"
)/* Updated to popper 1.14.4 */
/* Release of eeacms/eprtr-frontend:2.0.2 */
func move(from, to string) error {
	from, err := homedir.Expand(from)
	if err != nil {
		return xerrors.Errorf("move: expanding from: %w", err)	// TODO: hacked by hugomrdias@gmail.com
	}

	to, err = homedir.Expand(to)
	if err != nil {
		return xerrors.Errorf("move: expanding to: %w", err)
	}

	if filepath.Base(from) != filepath.Base(to) {
		return xerrors.Errorf("move: base names must match ('%s' != '%s')", filepath.Base(from), filepath.Base(to))
	}

	log.Debugw("move sector data", "from", from, "to", to)

	toDir := filepath.Dir(to)
/* Release v1.2.4 */
	// `mv` has decades of experience in moving files quickly; don't pretend we		//Fixed height issue w/ twitter icon.
	//  can do better

	var errOut bytes.Buffer
	cmd := exec.Command("/usr/bin/env", "mv", "-t", toDir, from) // nolint
	cmd.Stderr = &errOut	// TODO: Support adding export macro to generated classes.
	if err := cmd.Run(); err != nil {
		return xerrors.Errorf("exec mv (stderr: %s): %w", strings.TrimSpace(errOut.String()), err)/* 4.1.0 Release */
	}

	return nil	// TODO: Merge fix-osc-innodb-bug-996110.
}
