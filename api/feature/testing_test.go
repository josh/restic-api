package feature_test

import (
	"testing"

	"github.com/josh/restic-api/api/feature"
	rtest "github.com/josh/restic-api/api/test"
)

func TestSetFeatureFlag(t *testing.T) {
	flags := buildTestFlagSet()
	rtest.Assert(t, !flags.Enabled(alpha), "expected alpha feature to be disabled")

	restore := feature.TestSetFlag(t, flags, alpha, true)
	rtest.Assert(t, flags.Enabled(alpha), "expected alpha feature to be enabled")

	restore()
	rtest.Assert(t, !flags.Enabled(alpha), "expected alpha feature to be disabled again")
}
