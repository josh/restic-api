//go:build !selfupdate

package main

import (
	"github.com/josh/restic-api/api/global"
	"github.com/spf13/cobra"
)

func registerSelfUpdateCommand(_ *cobra.Command, _ *global.Options) {
	// No commands to register in non-selfupdate mode
}
