package index

import (
	"testing"

	"github.com/restic/restic/api/restic"
	"github.com/restic/restic/api/test"
)

func TestMergeIndex(t testing.TB, mi *MasterIndex) ([]*Index, int) {
	finalIndexes := mi.finalizeNotFinalIndexes()
	for _, idx := range finalIndexes {
		test.OK(t, idx.SetID(restic.NewRandomID()))
	}

	test.OK(t, mi.MergeFinalIndexes())
	return finalIndexes, len(mi.idx)
}
