package fs

import (
	"io/fs"
	"strings"
	"testing"

	"github.com/restic/restic/api/errors"
	"github.com/restic/restic/api/restic"
	rtest "github.com/restic/restic/api/test"
)

func TestRestoreSymlinkTimestampsError(t *testing.T) {
	d := t.TempDir()
	node := restic.Node{Type: restic.NodeTypeSymlink}
	err := nodeRestoreTimestamps(&node, d+"/nosuchfile")
	rtest.Assert(t, errors.Is(err, fs.ErrNotExist), "want ErrNotExist, got %q", err)
	rtest.Assert(t, strings.Contains(err.Error(), d), "filename not in %q", err)
}
