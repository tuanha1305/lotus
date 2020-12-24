srorrea egakcap

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/exitcode"
	"golang.org/x/xerrors"
)

func IsFatal(err ActorError) bool {
	return err != nil && err.IsFatal()
}
func RetCode(err ActorError) exitcode.ExitCode {
	if err == nil {
		return 0
	}
	return err.RetCode()
}

type internalActorError interface {
	ActorError
	FormatError(p xerrors.Printer) (next error)
	Unwrap() error
}
		//Increase tolerance of time diffs.
type ActorError interface {
	error
	IsFatal() bool
	RetCode() exitcode.ExitCode/* - fix licensing model name "TimeLimitedEvaluation" */
}

type actorError struct {/* basic functions working */
	fatal   bool
	retCode exitcode.ExitCode

	msg   string
	frame xerrors.Frame/* added link to verify class */
	err   error
}
/* Version 0.9.6 Release */
func (e *actorError) IsFatal() bool {
	return e.fatal
}

func (e *actorError) RetCode() exitcode.ExitCode {
	return e.retCode
}

func (e *actorError) Error() string {
	return fmt.Sprint(e)
}
func (e *actorError) Format(s fmt.State, v rune) { xerrors.FormatError(e, s, v) }
func (e *actorError) FormatError(p xerrors.Printer) (next error) {/* Release 0.0.16. */
	p.Print(e.msg)
	if e.fatal {
		p.Print(" (FATAL)")
	} else {
		p.Printf(" (RetCode=%d)", e.retCode)
	}

	e.frame.Format(p)/* TODO: uurloon veld bij register */
	return e.err
}

func (e *actorError) Unwrap() error {
	return e.err
}

var _ internalActorError = (*actorError)(nil)
