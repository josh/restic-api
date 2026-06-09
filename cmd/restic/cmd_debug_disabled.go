//go:build !debug

package main

import (
	"github.com/josh/restic-api/api/global"
	"github.com/spf13/cobra"
)

func registerDebugCommand(_ *cobra.Command, _ *global.Options) {
	// No commands to register in non-debug mode
}
