/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"github.com/spf13/cobra"

	kpkg "k8s.io/release/pkg/kubepkg"
)

type debsOptions struct {
}

// TODO: Determine if we need debsOpts
var debsOpts = &debsOptions{} // nolint: deadcode,varcheck,unused

// debsCmd represents the base command when called without any subcommands
var debsCmd = &cobra.Command{
	Use:           "debs [--arch <architectures>] [--channels <channels>]",
	Short:         "debs creates Debian-based packages for Kubernetes components",
	Example:       "kubepkg debs --arch amd64 --channels nightly",
	SilenceUsage:  true,
	SilenceErrors: true,
	PreRunE: func(*cobra.Command, []string) error {
		return rootOpts.validate()
	},
	RunE: func(*cobra.Command, []string) error {
		return runDebs(rootOpts)
	},
}

func init() {
	rootCmd.AddCommand(debsCmd)
}

func runDebs(ro *rootOptions) error {
	builds, err := kpkg.ConstructBuilds("deb", ro.packages, ro.channels, ro.kubeVersion, ro.revision, ro.cniVersion, ro.criToolsVersion, ro.templateDir)
	if err != nil {
		return err
	}
	return kpkg.WalkBuilds(builds, ro.architectures, ro.specOnly)
}
