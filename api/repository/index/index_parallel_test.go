package index_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/josh/restic-api/api/errors"
	"github.com/josh/restic-api/api/repository"
	"github.com/josh/restic-api/api/repository/index"
	"github.com/josh/restic-api/api/restic"
	rtest "github.com/josh/restic-api/api/test"
)

var repoFixture = filepath.Join("..", "testdata", "test-repo.tar.gz")

func TestRepositoryForAllIndexes(t *testing.T) {
	repo, _, cleanup := repository.TestFromFixture(t, repoFixture)
	defer cleanup()

	expectedIndexIDs := restic.NewIDSet()
	rtest.OK(t, repo.List(context.TODO(), restic.IndexFile, func(id restic.ID, size int64) error {
		expectedIndexIDs.Insert(id)
		return nil
	}))

	// check that all expected indexes are loaded without errors
	indexIDs := restic.NewIDSet()
	var indexErr error
	rtest.OK(t, index.ForAllIndexes(context.TODO(), repo, repo, func(id restic.ID, index *index.Index, oldFormat bool, err error) error {
		if err != nil {
			indexErr = err
		}
		indexIDs.Insert(id)
		return nil
	}))
	rtest.OK(t, indexErr)
	rtest.Equals(t, expectedIndexIDs, indexIDs)

	// must failed with the returned error
	iterErr := errors.New("error to pass upwards")

	err := index.ForAllIndexes(context.TODO(), repo, repo, func(id restic.ID, index *index.Index, oldFormat bool, err error) error {
		return iterErr
	})

	rtest.Equals(t, iterErr, err)
}
