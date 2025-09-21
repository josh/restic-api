package main

import (
	"github.com/josh/restic-api/api/archiver"
	"github.com/josh/restic-api/api/debug"
	"github.com/josh/restic-api/api/errors"
	"github.com/josh/restic-api/api/fs"
	"github.com/josh/restic-api/api/repository"
)

// rejectResticCache returns a RejectByNameFunc that rejects the restic cache
// directory (if set).
func rejectResticCache(repo *repository.Repository) (archiver.RejectByNameFunc, error) {
	if repo.Cache() == nil {
		return func(string) bool {
			return false
		}, nil
	}
	cacheBase := repo.Cache().BaseDir()

	if cacheBase == "" {
		return nil, errors.New("cacheBase is empty string")
	}

	return func(item string) bool {
		if fs.HasPathPrefix(cacheBase, item) {
			debug.Log("rejecting restic cache directory %v", item)
			return true
		}

		return false
	}, nil
}
