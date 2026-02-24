//go:build !linux

package main

import "github.com/spf13/cobra"

func checkSupportedCgroups(_ *cobra.Command) {
	// NOP on Non Linux
}
