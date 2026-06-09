//go:build !darwin && !freebsd && !linux

package main

import (
	"github.com/restic/restic/api/global"
	"github.com/spf13/cobra"
)

func registerMountCommand(_ *cobra.Command, _ *global.Options) {
	// Mount command not supported on these platforms
}
