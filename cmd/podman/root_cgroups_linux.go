//go:build linux

package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.podman.io/common/pkg/cgroups"
	"go.podman.io/common/pkg/config"
)

func checkSupportedCgroups(cmd *cobra.Command) {
	// If the user explicitly disabled cgroup management, skip the check entirely.
	if f := cmd.Root().Flag("cgroup-manager"); f != nil && f.Value.String() == config.DisabledCgroupsManager {
		return
	}
	unified, err := cgroups.IsCgroup2UnifiedMode()
	if err != nil {
		logrus.Fatalf("Error determining cgroups mode")
	}
	if !unified {
		logrus.Fatalf("Cgroups v1 not supported")
	}
}
