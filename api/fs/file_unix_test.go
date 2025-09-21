//go:build unix

package fs

import (
	"path/filepath"
	"syscall"
	"testing"

	"github.com/josh/restic-api/api/errors"
	rtest "github.com/josh/restic-api/api/test"
)

func TestReaddirnamesFifo(t *testing.T) {
	// should not block when reading from a fifo instead of a directory
	tempdir := t.TempDir()
	fifoFn := filepath.Join(tempdir, "fifo")
	rtest.OK(t, mkfifo(fifoFn, 0o600))

	_, err := Readdirnames(&Local{}, fifoFn, 0)
	rtest.Assert(t, errors.Is(err, syscall.ENOTDIR), "unexpected error %v", err)
}
