package fs

import (
	"io/fs"
	"strings"
	"testing"

	"github.com/josh/restic-api/api/errors"
	"github.com/josh/restic-api/api/restic"
	rtest "github.com/josh/restic-api/api/test"
)

func TestRestoreSymlinkTimestampsError(t *testing.T) {
	d := t.TempDir()
	node := restic.Node{Type: restic.NodeTypeSymlink}
	err := nodeRestoreTimestamps(&node, d+"/nosuchfile")
	rtest.Assert(t, errors.Is(err, fs.ErrNotExist), "want ErrNotExist, got %q", err)
	rtest.Assert(t, strings.Contains(err.Error(), d), "filename not in %q", err)
}
